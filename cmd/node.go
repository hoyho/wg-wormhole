package main

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var subCmdNameNode = "node"

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

var (
	nodeOpt nodeOption
)

func init() {
	NodeCommand.PersistentFlags().StringVarP(&nodeOpt.iface, "iface", "i", "", "the wireguard interface used for node")
	NodeCommand.PersistentFlags().StringVarP(&nodeOpt.token, "token", "t", "", "token for a simple verification")
	NodeCommand.PersistentFlags().StringVarP(&nodeOpt.regAddr, "address", "a", "", "remote rpc address of registry")
}

func (o nodeOption) validate() error {
	if o.iface == "" {
		return errors.New("`iface` is required")
	}

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
	iface   string
	regAddr string
	token   string
}

func runNode() error {
	return nil
}
