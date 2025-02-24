package provision

import (
	pkgcmd "github.com/foomo/sesamy-cli/pkg/cmd"
	conversionlinkerprovider "github.com/foomo/sesamy-cli/pkg/provider/conversionlinker"
	criteoprovider "github.com/foomo/sesamy-cli/pkg/provider/criteo"
	emarsysprovider "github.com/foomo/sesamy-cli/pkg/provider/emarsys"
	facebookprovider "github.com/foomo/sesamy-cli/pkg/provider/facebook"
	googleadsprovider "github.com/foomo/sesamy-cli/pkg/provider/googleads"
	googleanalyticsprovider "github.com/foomo/sesamy-cli/pkg/provider/googleanalytics"
	googletagprovider "github.com/foomo/sesamy-cli/pkg/provider/googletag"
	googletagmanagerprovider "github.com/foomo/sesamy-cli/pkg/provider/googletagmanager"
	microsoftadsprovider "github.com/foomo/sesamy-cli/pkg/provider/microsoftads"
	tracifyprovider "github.com/foomo/sesamy-cli/pkg/provider/tracify"
	umamiprovider "github.com/foomo/sesamy-cli/pkg/provider/umami"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	"github.com/pkg/errors"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewServer represents the server command
func NewServer(root *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Provision Google Tag Manager Server Container",
		RunE: func(cmd *cobra.Command, args []string) error {
			l := pkgcmd.Logger()
			l.Info("☕ Provisioning Server Container")

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
				cfg.GoogleTagManager.ServerContainer,
				tagmanager.WithRequestQuota(cfg.GoogleAPI.RequestQuota),
				tagmanager.WithClientOptions(cfg.GoogleAPI.GetClientOption()),
			)
			if err != nil {
				return err
			}

			if pkgcmd.Tag(googletagprovider.Tag, tags) {
				if err := googletagprovider.Server(tm, cfg.GoogleTag); err != nil {
					return errors.Wrap(err, "failed to provision google tag provider")
				}
			}

			if pkgcmd.Tag(googletagmanagerprovider.Tag, tags) {
				if err := googletagmanagerprovider.Server(tm, cfg.GoogleTagManager, cfg.EnableGeoResolution); err != nil {
					return errors.Wrap(err, "failed to provision google tag manager")
				}
			}

			if cfg.GoogleAnalytics.Enabled && pkgcmd.Tag(googleanalyticsprovider.Tag, tags) {
				l.Info("🅿️ Running provider", "name", googleanalyticsprovider.Name, "tag", googleanalyticsprovider.Tag)
				if err := googleanalyticsprovider.Server(tm, cfg.GoogleAnalytics, cfg.RedactVisitorIP, cfg.EnableGeoResolution); err != nil {
					return errors.Wrap(err, "failed to provision google analytics")
				}
			}

			if cfg.ConversionLinker.Enabled && pkgcmd.Tag(conversionlinkerprovider.Tag, tags) {
				l.Info("🅿️ Running provider", "name", conversionlinkerprovider.Name, "tag", conversionlinkerprovider.Tag)
				if err := conversionlinkerprovider.Server(tm, cfg.ConversionLinker); err != nil {
					return errors.Wrap(err, "failed to provision conversion linker")
				}
			}

			if cfg.Umami.Enabled && pkgcmd.Tag(umamiprovider.Tag, tags) {
				l.Info("🅿️ Running provider", "name", umamiprovider.Name, "tag", umamiprovider.Tag)
				if err := umamiprovider.Server(tm, cfg.Umami); err != nil {
					return errors.Wrap(err, "failed to provision umami")
				}
			}

			if cfg.Facebook.Enabled && pkgcmd.Tag(facebookprovider.Tag, tags) {
				l.Info("🅿️ Running provider", "name", facebookprovider.Name, "tag", facebookprovider.Tag)
				if err := facebookprovider.Server(l, tm, cfg.Facebook); err != nil {
					return errors.Wrap(err, "failed to provision facebook")
				}
			}

			if cfg.GoogleAds.Enabled && pkgcmd.Tag(googleadsprovider.Tag, tags) {
				l.Info("🅿️ Running provider", "name", googleadsprovider.Name, "tag", googleadsprovider.Tag)
				if err := googleadsprovider.Server(l, tm, cfg.GoogleAds); err != nil {
					return errors.Wrap(err, "failed to provision google ads")
				}
			}

			if cfg.Emarsys.Enabled && pkgcmd.Tag(emarsysprovider.Tag, tags) {
				l.Info("🅿️ Running provider", "name", emarsysprovider.Name, "tag", emarsysprovider.Tag)
				if err := emarsysprovider.Server(l, tm, cfg.Emarsys); err != nil {
					return errors.Wrap(err, "failed to provision emarsys")
				}
			}

			if cfg.Tracify.Enabled && pkgcmd.Tag(tracifyprovider.Tag, tags) {
				l.Info("🅿️ Running provider", "name", tracifyprovider.Name, "tag", tracifyprovider.Tag)
				if err := tracifyprovider.Server(l, tm, cfg.Tracify); err != nil {
					return errors.Wrap(err, "failed to provision tracify")
				}
			}

			if cfg.Criteo.Enabled && pkgcmd.Tag(criteoprovider.Tag, tags) {
				l.Info("🅿️ Running provider", "name", criteoprovider.Name, "tag", criteoprovider.Tag)
				if err := criteoprovider.Server(l, tm, cfg.Criteo); err != nil {
					return errors.Wrap(err, "failed to provision criteo")
				}
			}

			if cfg.MicrosoftAds.Enabled && pkgcmd.Tag(microsoftadsprovider.Tag, tags) {
				l.Info("🅿️ Running provider", "name", microsoftadsprovider.Name, "tag", microsoftadsprovider.Tag)
				if err := microsoftadsprovider.Server(l, tm, cfg.MicrosoftAds); err != nil {
					return errors.Wrap(err, "failed to provision microsoftads")
				}
			}

			if missed := tm.Missed(); len(tags) == 0 && len(missed) > 0 {
				tree := pterm.TreeNode{
					Text: "♻️ Missed resources (potentially garbage)",
				}
				for k, i := range missed {
					child := pterm.TreeNode{
						Text: k,
					}
					for _, s := range i {
						child.Children = append(child.Children, pterm.TreeNode{Text: s})
					}
					tree.Children = append(tree.Children, child)
				}
				if err := pterm.DefaultTree.WithRoot(tree).Render(); err != nil {
					l.Warn("failed to render missed resources", "error", err)
				}
			}

			return nil
		},
	}

	cmd.Flags().StringSlice("tags", nil, "list of tags to provision")
	_ = viper.BindPFlag("tags", cmd.Flags().Lookup("tags"))

	root.AddCommand(cmd)
}
