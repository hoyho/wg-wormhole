package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"golang.zx2c4.com/wireguard/wgctrl"
)

//NewRegistryCommand return a registry cmd definition
func NewRegistryCommand() *cobra.Command {
	var subCmdName = "registry"
	c := &cobra.Command{
		Use:   subCmdName,
		Short: "start a registry service",
		Long:  fmt.Sprintf("The `%s` command starts a %s service.", subCmdName, subCmdName),
		PreRun: func(cmd *cobra.Command, args []string) {
			//fmt.Println("cmdList PreRun")
		},
		Hidden:            false,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runRegistry()
		},
	}

	return c
}

func runRegistry() error {
	return nil
}

func ListPeers() {
	wgClient, err := wgctrl.New()
	if err != nil {
		log.Fatalf("error constructing Wireguard control client: %v",
			err)
	}
	wgDevice, err := wgClient.Device("utun7")
	if err != nil {
		log.Fatalf(
			"error retrieving Wireguard device '%s' info: %v",
			"wg0", err)
	}
	if len(wgDevice.Peers) < 1 {
		log.Println("no peers found")
		os.Exit(0)
	}

	buf, err := json.Marshal(wgDevice)
	if err != nil {
		log.Fatalf(
			"error marshal Wireguard device '%s' info: %v",
			"wg0", err)
	}
	fmt.Printf("%v", string(buf))

	// fmt.Printf("wgDevice: name:%s, priKey=%v,pubKey=%v, type=%v \n", wgDevice.Name,
	// 	wgDevice.PrivateKey.String(), wgDevice.PublicKey.String(), wgDevice.Type.String())

	// for _, p := range wgDevice.Peers {

}
