package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "latest"

func NewVersionCmd(root *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "show version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	}

	root.AddCommand(cmd)
}

func init() {
	NewVersionCmd(rootCmd)
}
