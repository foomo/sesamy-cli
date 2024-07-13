package tagmanager

import (
	pkgcmd "github.com/foomo/sesamy-cli/pkg/cmd"
	googleanaylticsprovider "github.com/foomo/sesamy-cli/pkg/provider/googleanalytics"
	googletagprovider "github.com/foomo/sesamy-cli/pkg/provider/googletag"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// NewWeb represents the web command
func NewWeb(root *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "web",
		Short: "Provision Google Tag Manager Web Container",
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
				cfg.GoogleTagManager.WebContainer,
				tagmanager.WithRequestQuota(cfg.GoogleAPI.RequestQuota),
				tagmanager.WithClientOptions(cfg.GoogleAPI.GetClientOption()),
			)
			if err != nil {
				return err
			}

			if pkgcmd.Tag(googletagprovider.Tag) {
				if err := googletagprovider.Web(tm, cfg.GoogleTag); err != nil {
					return errors.Wrap(err, "failed to provision google tag")
				}
			}

			if cfg.GoogleAnalytics.Enabled && pkgcmd.Tag(googleanaylticsprovider.Tag) {
				if err := googleanaylticsprovider.Web(tm, cfg.GoogleAnalytics); err != nil {
					return errors.Wrap(err, "failed to provision google analytics tag")
				}
			}

			return nil
		},
	}

	cmd.Flags().StringSlice("tags", nil, "list of tags to provision")

	root.AddCommand(cmd)
}
