package cmd

import (
	"github.com/spf13/cobra"
)

// tagmanagerCmd represents the tagmanager command
var tagmanagerCmd = &cobra.Command{
	Use:   "tagmanager",
	Short: "Provision Google Tag Manager containers",
}

func init() {
	rootCmd.AddCommand(tagmanagerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tagmanagerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tagmanagerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
