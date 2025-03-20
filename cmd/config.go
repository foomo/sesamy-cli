package cmd

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/alecthomas/chroma/quick"
	pkgcmd "github.com/foomo/sesamy-cli/pkg/cmd"
	"github.com/itchyny/json2yaml"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewConfig(root *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Print config",
		RunE: func(cmd *cobra.Command, args []string) error {
			l := pkgcmd.Logger()

			cfg, err := pkgcmd.ReadConfig(l, cmd)
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

	root.AddCommand(cmd)
}
