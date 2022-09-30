package rpc

import (
	"context"
	"errors"
	"log"

	pb "github.com/hoyho/wg-wormhole/proto"
	"golang.zx2c4.com/wireguard/wgctrl"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	verifyCode        string
	registryInterface string
)

//serverImpl is an implement of services defined in protobuf
type serverImpl struct {
	pb.UnimplementedRegistryServer
	wgClient *wgctrl.Client
}

func newRegistryServer() serverImpl {
	wgClient, err := wgctrl.New()
	if err != nil {
		log.Fatalf("error constructing Wireguard control client: %v",
			err)
	}
	return serverImpl{
		wgClient: wgClient,
	}
}

func (si serverImpl) preflight() error {
	if registryInterface == "" {
		return errors.New("service interface is not set")
	}
	_, err := si.wgClient.Device(registryInterface)
	if err != nil {
		log.Fatalf("error retrieving Wireguard device '%s' info: %v", registryInterface, err)
	}
	return nil
}

//GetEndpoint implement fetching endpoint
func (si serverImpl) GetEndpoint(ctx context.Context, req *pb.GetEndpointRequest) (reply *pb.GetEndpointReply, err error) {
	dev, err := si.wgClient.Device(registryInterface)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	reply = &pb.GetEndpointReply{
		DeviceName: dev.Name,
		DeviceType: dev.Type.String(),
		PubKey:     dev.PublicKey.String(),
	}

	for _, p := range dev.Peers {
		peer := pb.Peer{
			PubKey:        p.PublicKey.String(),
			TransmitBytes: p.TransmitBytes,
			ReceiveBytes:  p.ReceiveBytes,
		}
		if p.Endpoint != nil {
			peer.EndpointIP = p.Endpoint.IP.String()
			peer.EndpointPort = int32(p.Endpoint.Port)
		}
		if !p.LastHandshakeTime.IsZero() {
			peer.LastHandshakeUnixTimestamp = p.LastHandshakeTime.Unix()
		}

		for _, ip := range p.AllowedIPs {
			peer.AllowedIPs = append(peer.AllowedIPs, ip.String())
		}
		reply.Peers = append(reply.Peers, &peer)
	}

	return reply, nil

}
