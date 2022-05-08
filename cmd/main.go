package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func main() {
	baseName := filepath.Base(os.Args[0])
	err := NewCommand(baseName).Execute()

	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
}

//NewCommand starts the main program
func NewCommand(name string) *cobra.Command {

	var c = &cobra.Command{
		Use:               name,
		Short:             "UDP hole punching base on wireguard",
		SilenceErrors:     true,
		SilenceUsage:      true,
		DisableAutoGenTag: true,

		PersistentPreRunE: func(c *cobra.Command, args []string) error {
			return nil
		},
	}

	c.AddCommand(
		NewRegistryCommand(),
		NewNodeCommand(),
	)

	return c
}
