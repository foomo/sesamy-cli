package tagmanager

import (
	pkgcmd "github.com/foomo/sesamy-cli/pkg/cmd"
	cookiebotprovider "github.com/foomo/sesamy-cli/pkg/provider/cookiebot"
	emarsysprovider "github.com/foomo/sesamy-cli/pkg/provider/emarsys"
	googleanaylticsprovider "github.com/foomo/sesamy-cli/pkg/provider/googleanalytics"
	googletagprovider "github.com/foomo/sesamy-cli/pkg/provider/googletag"
	hotjarprovider "github.com/foomo/sesamy-cli/pkg/provider/hotjar"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewWeb represents the web command
func NewWeb(root *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "web",
		Short: "Provision Google Tag Manager Web Container",
		RunE: func(cmd *cobra.Command, args []string) error {
			l := pkgcmd.Logger()

			tags, err := cmd.Flags().GetStringSlice("tags")
			if err != nil {
				return errors.Wrap(err, "error reading tags flag")
			}

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

			if pkgcmd.Tag(googletagprovider.Tag, tags) {
				if err := googletagprovider.Web(tm, cfg.GoogleTag); err != nil {
					return errors.Wrap(err, "failed to provision google provider")
				}
			}

			if cfg.GoogleAnalytics.Enabled && pkgcmd.Tag(googleanaylticsprovider.Tag, tags) {
				if err := googleanaylticsprovider.Web(tm, cfg.GoogleAnalytics); err != nil {
					return errors.Wrap(err, "failed to provision google analytics provider")
				}
			}

			if cfg.Emarsys.Enabled && pkgcmd.Tag(emarsysprovider.Tag, tags) {
				if err := emarsysprovider.Web(tm, cfg.Emarsys); err != nil {
					return errors.Wrap(err, "failed to provision emarsys provider")
				}
			}

			if cfg.Hotjar.Enabled && pkgcmd.Tag(hotjarprovider.Tag, tags) {
				if err := hotjarprovider.Web(tm, cfg.Hotjar); err != nil {
					return errors.Wrap(err, "failed to provision hotjar provider")
				}
			}

			if cfg.Cookiebot.Enabled && pkgcmd.Tag(cookiebotprovider.Tag, tags) {
				if err := cookiebotprovider.Web(tm, cfg.Cookiebot); err != nil {
					return errors.Wrap(err, "failed to provision cookiebot provider")
				}
			}

			return nil
		},
	}

	cmd.Flags().StringSlice("tags", nil, "list of tags to provision")
	_ = viper.BindPFlag("tags", cmd.Flags().Lookup("tags"))

	root.AddCommand(cmd)
}
