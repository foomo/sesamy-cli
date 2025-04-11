package provision

import (
	"log/slog"

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
	ptermx "github.com/foomo/sesamy-cli/pkg/pterm"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"github.com/pkg/errors"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewServer represents the server command
func NewServer(l *slog.Logger) *cobra.Command {
	c := viper.New()

	cmd := &cobra.Command{
		Use:   "server",
		Short: "Provision Google Tag Manager Server Container",
		RunE: func(cmd *cobra.Command, args []string) error {
			l := pkgcmd.NewLogger()
			l.Info("☕ Provisioning Server Container")

			tags, err := cmd.Flags().GetStringSlice("tags")
			if err != nil {
				return errors.Wrap(err, "error reading tags flag")
			}

			cfg, err := pkgcmd.ReadConfig(l, c, cmd)
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

			if err := tm.EnsureWorkspaceID(cmd.Context()); err != nil {
				return err
			}

			if utils.Tag(googletagprovider.Tag, tags) {
				if err := googletagprovider.Server(cmd.Context(), tm, cfg.GoogleTag); err != nil {
					return errors.Wrap(err, "failed to provision google tag provider")
				}
			}

			if utils.Tag(googletagmanagerprovider.Tag, tags) {
				if err := googletagmanagerprovider.Server(cmd.Context(), tm, cfg.GoogleTagManager, cfg.EnableGeoResolution); err != nil {
					return errors.Wrap(err, "failed to provision google tag manager")
				}
			}

			if cfg.GoogleAnalytics.Enabled && utils.Tag(googleanalyticsprovider.Tag, tags) {
				l.Info("🅿️ Running provider", "name", googleanalyticsprovider.Name, "tag", googleanalyticsprovider.Tag)
				if err := googleanalyticsprovider.Server(cmd.Context(), tm, cfg.GoogleAnalytics, cfg.RedactVisitorIP, cfg.EnableGeoResolution); err != nil {
					return errors.Wrap(err, "failed to provision google analytics")
				}
			}

			if cfg.ConversionLinker.Enabled && utils.Tag(conversionlinkerprovider.Tag, tags) {
				l.Info("🅿️ Running provider", "name", conversionlinkerprovider.Name, "tag", conversionlinkerprovider.Tag)
				if err := conversionlinkerprovider.Server(cmd.Context(), tm, cfg.ConversionLinker); err != nil {
					return errors.Wrap(err, "failed to provision conversion linker")
				}
			}

			if cfg.Umami.Enabled && utils.Tag(umamiprovider.Tag, tags) {
				l.Info("🅿️ Running provider", "name", umamiprovider.Name, "tag", umamiprovider.Tag)
				if err := umamiprovider.Server(cmd.Context(), tm, cfg.Umami); err != nil {
					return errors.Wrap(err, "failed to provision umami")
				}
			}

			if cfg.Facebook.Enabled && utils.Tag(facebookprovider.Tag, tags) {
				l.Info("🅿️ Running provider", "name", facebookprovider.Name, "tag", facebookprovider.Tag)
				if err := facebookprovider.Server(cmd.Context(), l, tm, cfg.Facebook); err != nil {
					return errors.Wrap(err, "failed to provision facebook")
				}
			}

			if cfg.GoogleAds.Enabled && utils.Tag(googleadsprovider.Tag, tags) {
				l.Info("🅿️ Running provider", "name", googleadsprovider.Name, "tag", googleadsprovider.Tag)
				if err := googleadsprovider.Server(cmd.Context(), l, tm, cfg.GoogleAds); err != nil {
					return errors.Wrap(err, "failed to provision google ads")
				}
			}

			if cfg.Emarsys.Enabled && utils.Tag(emarsysprovider.Tag, tags) {
				l.Info("🅿️ Running provider", "name", emarsysprovider.Name, "tag", emarsysprovider.Tag)
				if err := emarsysprovider.Server(cmd.Context(), l, tm, cfg.Emarsys); err != nil {
					return errors.Wrap(err, "failed to provision emarsys")
				}
			}

			if cfg.Tracify.Enabled && utils.Tag(tracifyprovider.Tag, tags) {
				l.Info("🅿️ Running provider", "name", tracifyprovider.Name, "tag", tracifyprovider.Tag)
				if err := tracifyprovider.Server(cmd.Context(), l, tm, cfg.Tracify); err != nil {
					return errors.Wrap(err, "failed to provision tracify")
				}
			}

			if cfg.Criteo.Enabled && utils.Tag(criteoprovider.Tag, tags) {
				l.Info("🅿️ Running provider", "name", criteoprovider.Name, "tag", criteoprovider.Tag)
				if err := criteoprovider.Server(cmd.Context(), l, tm, cfg.Criteo); err != nil {
					return errors.Wrap(err, "failed to provision criteo")
				}
			}

			if cfg.MicrosoftAds.Enabled && utils.Tag(microsoftadsprovider.Tag, tags) {
				l.Info("🅿️ Running provider", "name", microsoftadsprovider.Name, "tag", microsoftadsprovider.Tag)
				if err := microsoftadsprovider.Server(cmd.Context(), l, tm, cfg.MicrosoftAds); err != nil {
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

				if err := pterm.DefaultTree.WithRoot(tree).WithWriter(ptermx.NewWriter(pterm.Warning)).Render(); err != nil {
					l.Warn("failed to render missed resources", "error", err)
				}
			}

			return nil
		},
	}

	flags := cmd.Flags()

	flags.StringSliceP("config", "c", []string{"sesamy.yaml"}, "config files (default is sesamy.yaml)")
	_ = c.BindPFlag("config", flags.Lookup("config"))

	flags.StringSlice("tags", nil, "list of tags to provision")
	_ = c.BindPFlag("tags", flags.Lookup("tags"))

	return cmd
}
