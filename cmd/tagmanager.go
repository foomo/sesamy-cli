package cmd

import (
	"github.com/spf13/cobra"
)

// NewTagmanagerCmd represents the tagmanager command
func NewTagmanagerCmd(root *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tagmanager",
		Short: "Provision Google Tag Manager containers",
	}

	root.AddCommand(cmd)

	return cmd
}

func init() {
	cmd := NewTagmanagerCmd(rootCmd)
	NewTagManagerServerCmd(cmd)
	NewTagmanagerWebCmd(cmd)
}
