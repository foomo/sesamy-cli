package cmd

import (
	"os"
	"path"
	"path/filepath"

	"github.com/foomo/gocontemplate/pkg/contemplate"
	"github.com/foomo/sesamy-cli/pkg/typescript/generator"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
)

// typescriptCmd represents the typescript command
var typescriptCmd = &cobra.Command{
	Use:               "typescript",
	Short:             "Generate typescript events",
	PersistentPreRunE: preRunReadConfig,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctpl, err := contemplate.Load(&cfg.Typescript.Config)
		if err != nil {
			return err
		}

		files, err := generator.Generate(logger, ctpl)
		if err != nil {
			return errors.Wrap(err, "failed to get build typescript")
		}

		outPath, err := filepath.Abs(cfg.Typescript.OutputPath)
		if err != nil {
			return errors.Wrap(err, "failed to get output path")
		}

		if err = os.MkdirAll(outPath, os.ModePerm); err != nil {
			return errors.Wrap(err, "failed to create typescript output directory")
		}

		logger.InfoContext(cmd.Context(), "generated typescript code", "dir", outPath, "files", maps.Keys(files))
		for filename, file := range files {
			if err = os.WriteFile(path.Join(outPath, filename), []byte(file.String()), 0600); err != nil {
				return errors.Wrap(err, "failed to write typescript code")
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(typescriptCmd)
}
