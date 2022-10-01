package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/hoyho/wg-wormhole/proto"
	"github.com/spf13/cobra"
	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"google.golang.org/grpc"
)

var (
	subCmdNameNode = "node"
	nodeOpt        nodeOption
)

var NodeCommand = &cobra.Command{
	Use:   subCmdNameNode,
	Short: "start a node service",
	Long:  fmt.Sprintf("The `%s` command starts a node service.", subCmdNameNode),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		err := nodeOpt.validate()
		if err != nil {
			fmt.Println("invalid parameter. you may need add -h for help")
			fmt.Println()
		}
		return err
	},
	Hidden:            false,
	DisableAutoGenTag: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runNode()
	},
}

func init() {
	NodeCommand.PersistentFlags().StringVarP(&nodeOpt.iface, "iface", "i", "", "the wireguard interface used for this node [optional] ")
	NodeCommand.PersistentFlags().StringVarP(&nodeOpt.token, "token", "t", "", "token for a simple verification")
	NodeCommand.PersistentFlags().StringVarP(&nodeOpt.regAddr, "address", "a", "", "remote rpc address of registry")
	NodeCommand.PersistentFlags().DurationVarP(&nodeOpt.refreshInterval, "refresh", "r", 15*time.Second, "refresh interval for the peers")
}

func (o nodeOption) validate() error {
	if o.regAddr == "" {
		return errors.New("`address` is required")
	}

	if o.token == "" {
		return errors.New("`token`is required")
	}

	return nil
}

//nodeOption defines the option for node
type nodeOption struct {
	iface           string
	regAddr         string
	token           string
	refreshInterval time.Duration
}

func runNode() error {
	wgClient, err := wgctrl.New()
	if err != nil {
		log.Fatalf("error constructing Wireguard control client: %v", err)
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure(), grpc.WithBlock())

	conn, err := grpc.Dial(nodeOpt.regAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial:%v, err: %v", nodeOpt.regAddr, err)
	}
	defer conn.Close()
	client := pb.NewRegistryClient(conn)

	for {

		devs, err := wgClient.Devices()
		if err != nil {
			log.Println("fetch devices error", err)
		}

		for _, dev := range devs {
			//if specific any filter for dev. todo(Hoyho): use regex
			if len(nodeOpt.iface) > 0 && dev.Name != nodeOpt.iface {
				continue
			}
			if err := syncDevice(dev, client); err != nil {
				log.Println("sync device error:", err)
			}
		}

		time.Sleep(nodeOpt.refreshInterval)
		fmt.Println("   ")
		fmt.Println("   ")

	}
}

var errDevUndefine = errors.New("device is nil")
var errRegUndefine = errors.New("registryInfo is nil")

func syncDevice(dev *wgtypes.Device, regClient pb.RegistryClient) error {
	log.Println("-------device handler-----")
	if dev == nil {
		return errDevUndefine
	}

	regInfo, callErr := regClient.GetEndpoint(context.Background(), &pb.GetEndpointRequest{Token: nodeOpt.token})
	if callErr != nil {
		log.Println("fetch devices error", callErr)
		return callErr
	}

	log.Println("Device name:", dev.Name)
	log.Println("PublicKey:", dev.PublicKey)
	log.Println("Type:", dev.Type)
	log.Println("PeerCount:", len(dev.Peers))

	for _, peer := range dev.Peers {
		syncPeer(dev, peer, regInfo)
	}
	return nil

}

func syncPeer(dev *wgtypes.Device, localPeer wgtypes.Peer, registryInfo *pb.GetEndpointReply) error {

	if registryInfo == nil || len(registryInfo.Peers) == 0 {
		return errRegUndefine
	}
	log.Println("=======peer handler======")
	log.Printf("    endpoint:%v,publicKey:%v LastHandshakeTime:%v\n", localPeer.Endpoint, localPeer.PublicKey, localPeer.LastHandshakeTime)

	newConfig := wgtypes.Config{
		PrivateKey:   &dev.PrivateKey,
		ListenPort:   &dev.ListenPort,
		ReplacePeers: false,
	}
	if dev.FirewallMark > 0 {
		newConfig.FirewallMark = &dev.FirewallMark
	}
	updated := false

	for _, rPeer := range registryInfo.Peers {
		lPubKey := localPeer.PublicKey.String()
		if rPeer.PubKey != lPubKey {
			continue
		}
		//found a match peer from local to registry
		if isNewEndpoint(localPeer.Endpoint, rPeer.EndpointIP, int(rPeer.EndpointPort)) {
			rPeerEndpoint := fmt.Sprintf("%s:%d", rPeer.EndpointIP, rPeer.EndpointPort)
			log.Printf("!!! peer's endpoint updated. from:%v ----> to: %v \n", localPeer.Endpoint.String(), rPeerEndpoint)
			updated = updated || true

			newPeer := wgtypes.PeerConfig{
				PublicKey:                   localPeer.PublicKey,
				PresharedKey:                &localPeer.PresharedKey,
				PersistentKeepaliveInterval: &localPeer.PersistentKeepaliveInterval,
				AllowedIPs:                  localPeer.AllowedIPs,
				UpdateOnly:                  true,
				Endpoint: &net.UDPAddr{
					IP:   net.ParseIP(rPeer.EndpointIP),
					Port: int(rPeer.EndpointPort),
				},
			}
			newConfig.Peers = append(newConfig.Peers, newPeer)
		}

	}

	if updated {
		log.Println("updating device:", dev.Name)
		wgClient, err := wgctrl.New()
		if err != nil {
			log.Fatalf("error constructing Wireguard control client: %v", err)
		}
		err = wgClient.ConfigureDevice(dev.Name, newConfig)
		if err != nil {
			log.Fatalf("error ConfigureDevice Wireguard control client: %v", err)
		}
	}

	return nil
}

func isNewEndpoint(localEndpoint *net.UDPAddr, remoteEndpointIP string, remoteEndpointPort int) bool {
	if localEndpoint != nil {
		return localEndpoint.IP.String() != remoteEndpointIP ||
			localEndpoint.Port != remoteEndpointPort
	} else {
		//local endpoint is nil, but remote is valid, then it is updated event
		return len(remoteEndpointIP) > 0 && remoteEndpointPort > 0
	}

}
