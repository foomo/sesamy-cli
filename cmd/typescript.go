package cmd

import (
	"os"
	"path/filepath"

	"github.com/foomo/sesamy-cli/internal"
	"github.com/foomo/sesamy-cli/pkg/typescript"
	"github.com/pkg/errors"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// typescriptCmd represents the typescript command
var typescriptCmd = &cobra.Command{
	Use:               "typescript",
	Short:             "Generate typescript events",
	PersistentPreRunE: preRunReadConfig,
	RunE: func(cmd *cobra.Command, args []string) error {

		eventTypes, err := internal.GetStructTypes(cmd.Context(), cfg.Typescript.Packages)
		if err != nil {
			return err
		}

		code, err := typescript.Generate(eventTypes)
		if err != nil {
			return err
		}

		outPath, err := filepath.Abs(cfg.Typescript.OutputPath)
		if err != nil {
			return errors.Wrap(err, "failed to get output path")
		}
		pterm.Info.Printfln("Generated typescript code to: %s", outPath)

		if err = os.MkdirAll(filepath.Dir(outPath), os.ModePerm); err != nil {
			return errors.Wrap(err, "failed to create typescript output directory")
		}

		if err = os.WriteFile(outPath, []byte(code), 0600); err != nil {
			return errors.Wrap(err, "failed to write typescript code")
		}

		return nil
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
