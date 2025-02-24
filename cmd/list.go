package cmd

import (
	"github.com/foomo/sesamy-cli/cmd/list"
	"github.com/spf13/cobra"
)

// NewList represents the list command
func NewList(root *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List Google Tag Manager containers",
	}

	list.NewServer(cmd)
	list.NewWeb(cmd)
	root.AddCommand(cmd)

	return cmd
}
