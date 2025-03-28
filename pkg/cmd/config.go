package cmd

import (
	"io"
	"log/slog"

	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/knadh/koanf/v2"
	"github.com/pkg/errors"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func ReadConfig(l *slog.Logger, c *viper.Viper, cmd *cobra.Command) (*config.Config, error) {
	k := koanf.NewWithConf(koanf.Conf{
		Delim: "/",
	})

	for _, filename := range c.GetStringSlice("config") {
		var p koanf.Provider
		if filename == "-" {
			pterm.Debug.Println("reading config from stdin")
			if b, err := io.ReadAll(cmd.InOrStdin()); err != nil {
				return nil, err
			} else {
				p = rawbytes.Provider(b)
			}
		} else {
			pterm.Debug.Println("reading config from filename: " + filename)
			p = file.Provider(filename)
		}
		if err := k.Load(p, yaml.Parser()); err != nil {
			return nil, errors.Wrap(err, "error loading config file: "+filename)
		}
	}

	var cfg *config.Config
	pterm.Debug.Println("unmarshalling config")
	if err := k.UnmarshalWithConf("", &cfg, koanf.UnmarshalConf{
		Tag: "yaml",
	}); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal config")
	}

	if cfg.Version != config.Version {
		return nil, errors.New("missing or invalid config version: " + cfg.Version + " != '" + config.Version + "'")
	}

	return cfg, nil
}
