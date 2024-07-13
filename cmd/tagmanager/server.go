package tagmanager

import (
	pkgcmd "github.com/foomo/sesamy-cli/pkg/cmd"
	conversionlinkerprovider "github.com/foomo/sesamy-cli/pkg/provider/conversionlinker"
	googleanalyticsprovider "github.com/foomo/sesamy-cli/pkg/provider/googleanalytics"
	googletagprovider "github.com/foomo/sesamy-cli/pkg/provider/googletag"
	googletagmanagerprovider "github.com/foomo/sesamy-cli/pkg/provider/googletagmanager"
	umamiprovider "github.com/foomo/sesamy-cli/pkg/provider/umami"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// NewServer represents the server command
func NewServer(root *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Provision Google Tag Manager Server Container",
		RunE: func(cmd *cobra.Command, args []string) error {
			l := pkgcmd.Logger()

			cfg, err := pkgcmd.ReadConfig(l, cmd)
			if err != nil {
				return err
			}

			tm, err := tagmanager.New(
				cmd.Context(),
				l,
				cfg.GoogleTagManager.AccountID,
				cfg.GoogleTagManager.ServerContainer,
				tagmanager.WithRequestQuota(cfg.GoogleAPI.RequestQuota),
				tagmanager.WithClientOptions(cfg.GoogleAPI.GetClientOption()),
			)
			if err != nil {
				return err
			}

			if pkgcmd.Tag(googletagprovider.Tag) {
				if err := googletagprovider.Server(tm); err != nil {
					return errors.Wrap(err, "failed to provision google tag")
				}
			}

			if pkgcmd.Tag(googletagmanagerprovider.Tag) {
				if err := googletagmanagerprovider.Server(tm, cfg.GoogleTagManager); err != nil {
					return errors.Wrap(err, "failed to provision google tag manager")
				}
			}

			if cfg.GoogleAnalytics.Enabled && pkgcmd.Tag(googleanalyticsprovider.Tag) {
				if err := googleanalyticsprovider.Server(tm, cfg.GoogleAnalytics, cfg.RedactVisitorIP); err != nil {
					return errors.Wrap(err, "failed to provision google analytics")
				}
			}

			if cfg.ConversionLinker.Enabled && pkgcmd.Tag(conversionlinkerprovider.Tag) {
				if err := conversionlinkerprovider.Server(tm, cfg.ConversionLinker); err != nil {
					return errors.Wrap(err, "failed to provision conversion linker")
				}
			}

			if cfg.Umami.Enabled && pkgcmd.Tag(umamiprovider.Tag) {
				if err := umamiprovider.Server(tm, cfg.Umami); err != nil {
					return errors.Wrap(err, "failed to provision umammi")
				}
			}

			return nil
		},
	}

	cmd.Flags().StringSlice("tags", nil, "list of tags to provision")

	root.AddCommand(cmd)
}
