package cmd

import (
	"maps"
	"os"
	"path"
	"path/filepath"
	"slices"

	"github.com/foomo/gocontemplate/pkg/contemplate"
	pkgcmd "github.com/foomo/sesamy-cli/pkg/cmd"
	"github.com/foomo/sesamy-cli/pkg/typescript/generator"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// NewTypeScript represents the typescript command
func NewTypeScript(root *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "typescript",
		Short: "Generate typescript events",
		RunE: func(cmd *cobra.Command, args []string) error {
			l := pkgcmd.Logger()

			cfg, err := pkgcmd.ReadConfig(l, cmd)
			if err != nil {
				return err
			}

			ctpl, err := contemplate.Load(cmd.Context(), &cfg.GoogleTag.TypeScript.Config)
			if err != nil {
				return err
			}

			files, err := generator.Generate(l, ctpl)
			if err != nil {
				return errors.Wrap(err, "failed to get build typescript")
			}

			outPath, err := filepath.Abs(cfg.GoogleTag.TypeScript.OutputPath)
			if err != nil {
				return errors.Wrap(err, "failed to get output path")
			}

			if err = os.MkdirAll(outPath, os.ModePerm); err != nil {
				return errors.Wrap(err, "failed to create typescript output directory")
			}

			l.InfoContext(cmd.Context(), "generated typescript code", "dir", outPath, "files", slices.AppendSeq(make([]string, 0, len(files)), maps.Keys(files)))
			for filename, file := range files {
				if err = os.WriteFile(path.Join(outPath, filename), []byte(file.String()), 0600); err != nil {
					return errors.Wrap(err, "failed to write typescript code")
				}
			}

			return nil
		},
	}

	root.AddCommand(cmd)

	return cmd
}
