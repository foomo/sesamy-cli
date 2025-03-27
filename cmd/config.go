package cmd

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"os"

	"github.com/alecthomas/chroma/quick"
	pkgcmd "github.com/foomo/sesamy-cli/pkg/cmd"
	"github.com/itchyny/json2yaml"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewConfig(l *slog.Logger) *cobra.Command {
	c := viper.New()

	cmd := &cobra.Command{
		Use:   "config",
		Short: "Print config",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := pkgcmd.ReadConfig(l, c, cmd)
			if err != nil {
				return err
			}

			out, err := json.MarshalIndent(cfg, "", "  ")
			if err != nil {
				return errors.Wrap(err, "failed to marshal config")
			}

			var buf bytes.Buffer
			if err := json2yaml.Convert(&buf, bytes.NewBuffer(out)); err != nil {
				return errors.Wrap(err, "failed to convert config")
			}

			return quick.Highlight(os.Stdout, buf.String(), "yaml", "terminal", "monokai")
		},
	}

	flags := cmd.Flags()

	flags.StringSliceP("config", "c", []string{"sesamy.yaml"}, "config files (default is sesamy.yaml)")
	_ = c.BindPFlag("config", flags.Lookup("config"))

	return cmd
}
