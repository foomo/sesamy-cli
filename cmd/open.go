package cmd

import (
	"fmt"
	"log/slog"

	pkgcmd "github.com/foomo/sesamy-cli/pkg/cmd"
	"github.com/pkg/browser"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewOpen represents the open command
func NewOpen(l *slog.Logger) *cobra.Command {
	c := viper.New()

	cmd := &cobra.Command{
		Use:   "open",
		Short: "Open links in the browser",
		Args:  cobra.OnlyValidArgs,
		ValidArgs: []cobra.Completion{
			"ga",
			"gtm-web",
			"gtm-server",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := pkgcmd.ReadConfig(l, c, cmd)
			if err != nil {
				return err
			}

			var url string
			switch args[0] {
			case "ga":
				if cfg.GoogleAnalytics.PropertyID == "" {
					return errors.New("missing Google Analytics Property ID")
				}
				url = fmt.Sprintf(
					"https://analytics.google.com/analytics/web/#/p%s/",
					cfg.GoogleAnalytics.PropertyID,
				)
			case "gtm-web":
				url = fmt.Sprintf(
					"https://tagmanager.google.com/#/container/accounts/%s/containers/%s/",
					cfg.GoogleTagManager.AccountID,
					cfg.GoogleTagManager.WebContainer.ContainerID,
				)
			case "gtm-server":
				url = fmt.Sprintf(
					"https://tagmanager.google.com/#/container/accounts/%s/containers/%s/",
					cfg.GoogleTagManager.AccountID,
					cfg.GoogleTagManager.ServerContainer.ContainerID,
				)
			default:
				return fmt.Errorf("invalid container type: %s", args[0])
			}

			l.Info("↗ Navigating to Google Tag Manager Container: " + url)

			return browser.OpenURL(url)
		},
	}

	flags := cmd.Flags()

	flags.StringSliceP("config", "c", []string{"sesamy.yaml"}, "config files (default is sesamy.yaml)")
	_ = c.BindPFlag("config", flags.Lookup("config"))

	return cmd
}
