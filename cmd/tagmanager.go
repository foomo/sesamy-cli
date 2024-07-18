package cmd

import (
	"github.com/foomo/sesamy-cli/cmd/tagmanager"
	"github.com/spf13/cobra"
)

// NewTagmanager represents the tagmanager command
func NewTagmanager(root *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tagmanager",
		Short: "Provision Google Tag Manager containers",
	}

	tagmanager.NewServer(cmd)
	tagmanager.NewWeb(cmd)
	root.AddCommand(cmd)

	return cmd
}
