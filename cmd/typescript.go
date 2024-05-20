package cmd

import (
	"os"
	"path"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
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
		spew.Dump(cfg.Typescript)
		parser := internal.NewLoader(&cfg.Typescript.LoaderConfig)
		if err := parser.Load(cmd.Context()); err != nil {
			return err
		}

		generator := typescript.NewBuilder(parser)
		files, err := generator.Build(cmd.Context())
		if err != nil {
			return errors.Wrap(err, "failed to get build typescript")
		}

		outPath, err := filepath.Abs(cfg.Typescript.OutputPath)
		if err != nil {
			return errors.Wrap(err, "failed to get output path")
		}
		pterm.Info.Printfln("generated typescript code to: %s", outPath)

		if err = os.MkdirAll(outPath, os.ModePerm); err != nil {
			return errors.Wrap(err, "failed to create typescript output directory")
		}

		for filename, file := range files {
			pterm.Info.Printfln("...%s", filename)
			if err = os.WriteFile(path.Join(outPath, filename), []byte(file.String()), 0600); err != nil {
				return errors.Wrap(err, "failed to write typescript code")
			}
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
