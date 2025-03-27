package cmd

import (
	"log/slog"
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
	"github.com/spf13/viper"
)

// NewTypeScript represents the typescript command
func NewTypeScript(l *slog.Logger) *cobra.Command {
	c := viper.New()

	cmd := &cobra.Command{
		Use:   "typescript",
		Short: "Generate typescript events",
		RunE: func(cmd *cobra.Command, args []string) error {
			l := pkgcmd.NewLogger()

			cfg, err := pkgcmd.ReadConfig(l, c, cmd)
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

	flags := cmd.Flags()

	flags.StringSliceP("config", "c", []string{"sesamy.yaml"}, "config files (default is sesamy.yaml)")
	_ = c.BindPFlag("config", flags.Lookup("config"))

	return cmd
}
