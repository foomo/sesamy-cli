package cmd

import (
	"github.com/gzuidhof/tygo/tygo"
	"github.com/spf13/cobra"
)

// typescriptCmd represents the typescript command
var typescriptCmd = &cobra.Command{
	Use:               "typescript",
	Short:             "Generate typescript events",
	PersistentPreRunE: preRunReadConfig,
	RunE: func(cmd *cobra.Command, args []string) error {
		gen := tygo.New(&tygo.Config{
			Packages: cfg.Typescript.Packages,
		})
		for k, v := range cfg.Typescript.TypeMappings {
			gen.SetTypeMapping(k, v)
		}

		return gen.Generate()
	},
}

func init() {
	rootCmd.AddCommand(typescriptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// typescriptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// typescriptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
