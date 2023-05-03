package main

import (
	"errors"
	"fmt"

	rpcServer "github.com/hoyho/wg-wormhole/pkg/rpc"
	"github.com/hoyho/wg-wormhole/pkg/webserver"
	"github.com/spf13/cobra"
)

var subCmdNameReg = "registry"

var RegistryCommand = &cobra.Command{
	Use:   subCmdNameReg,
	Short: "start a registry service",
	Long:  fmt.Sprintf("The `%s` command starts a %s service.", subCmdNameReg, subCmdNameReg),
	// PreRun: func(cmd *cobra.Command, args []string) {
	// 	if err := validate(regOpt); err != nil {
	// 		fmt.Println("use -h for help")
	// 	}
	// },
	PreRunE: func(cmd *cobra.Command, args []string) error {
		err := regOpt.validate()
		if err != nil {
			fmt.Println("invalid parameter. you may need add -h for help")
			fmt.Println()
		}
		return err
	},
	Hidden:            false,
	DisableAutoGenTag: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runRegistry(regOpt)
	},
}

var (
	regOpt registryOption
)

func init() {
	RegistryCommand.PersistentFlags().StringVarP(&regOpt.iface, "ifcae", "i", "", "the wireguard interface used for registry")
	RegistryCommand.PersistentFlags().StringVarP(&regOpt.token, "token", "t", "", "token for a simple verification")
	RegistryCommand.PersistentFlags().StringVarP(&regOpt.rpcAddress, "address", "a", ":1638", "rpc service listening address")
	RegistryCommand.PersistentFlags().StringVarP(&regOpt.webAddress, "address-http", "A", ":1639", "http server listening address")
	RegistryCommand.PersistentFlags().StringVarP(&regOpt.watchPubKey, "watch-peer-key", "w", "nFz1T36SEfpGrQKUYAZoPTL4soGCesObVS+Groooooo=", "watch a peer endpoint in http server")
}

func (o registryOption) validate() error {
	if o.iface == "" {
		return errors.New("`iface` is required")
	}

	if o.rpcAddress == "" {
		return errors.New("`address` is required")
	}

	if o.token == "" {
		return errors.New("`token`is required")
	}

	return nil
}

// registryOption defines the option for registry
type registryOption struct {
	iface      string
	rpcAddress string
	token      string

	webAddress  string
	watchPubKey string
}

func runRegistry(o registryOption) error {

	go func() {
		if r := recover(); r != nil {
			fmt.Println("http server panic:", r)
		}
		webserver.Serve(o.webAddress, o.token, o.iface, o.watchPubKey)
	}()

	rpcServer.ServeRpcForever(o.rpcAddress, o.token, o.iface)
	return nil
}
