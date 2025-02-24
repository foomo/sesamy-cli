package cmd

import (
	"github.com/foomo/sesamy-cli/cmd/provision"
	"github.com/spf13/cobra"
)

// NewProvision represents the provision command
func NewProvision(root *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "provision",
		Short: "Provision Google Tag Manager containers",
	}

	provision.NewServer(cmd)
	provision.NewWeb(cmd)
	root.AddCommand(cmd)

	return cmd
}
