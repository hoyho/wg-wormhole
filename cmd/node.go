package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

//NewNodeCommand return a node cmd definition
func NewNodeCommand() *cobra.Command {
	var subCmdName = "node"
	c := &cobra.Command{
		Use:   subCmdName,
		Short: "start a node service",
		Long:  fmt.Sprintf("The `%s` command starts a node service.", subCmdName),
		PreRun: func(cmd *cobra.Command, args []string) {
		},
		Hidden:            false,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runNode()
		},
	}

	return c
}

func runNode() error {
	return nil
}
