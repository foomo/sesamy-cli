package provision

import (
	"log/slog"

	pkgcmd "github.com/foomo/sesamy-cli/pkg/cmd"
	cookiebotprovider "github.com/foomo/sesamy-cli/pkg/provider/cookiebot"
	criteoprovider "github.com/foomo/sesamy-cli/pkg/provider/criteo"
	emarsysprovider "github.com/foomo/sesamy-cli/pkg/provider/emarsys"
	googleanaylticsprovider "github.com/foomo/sesamy-cli/pkg/provider/googleanalytics"
	googletagprovider "github.com/foomo/sesamy-cli/pkg/provider/googletag"
	googletagmanagerprovider "github.com/foomo/sesamy-cli/pkg/provider/googletagmanager"
	hotjarprovider "github.com/foomo/sesamy-cli/pkg/provider/hotjar"
	ptermx "github.com/foomo/sesamy-cli/pkg/pterm"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"github.com/pkg/errors"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewWeb represents the web command
func NewWeb(l *slog.Logger) *cobra.Command {
	c := viper.New()

	cmd := &cobra.Command{
		Use:   "web",
		Short: "Provision Google Tag Manager Web Container",
		RunE: func(cmd *cobra.Command, args []string) error {
			l := pkgcmd.NewLogger()
			l.Info("☕ Provisioning Web Container")

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
				cfg.GoogleTagManager.WebContainer,
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
				l.Info("🅿️ Running provider", "name", googletagprovider.Name, "tag", googletagprovider.Tag)
				if err := googletagprovider.Web(cmd.Context(), tm, cfg.GoogleTag); err != nil {
					return errors.Wrap(err, "failed to provision google tag provider")
				}
			}

			if utils.Tag(googletagmanagerprovider.Tag, tags) {
				if err := googletagmanagerprovider.Web(cmd.Context(), tm, cfg.GoogleTagManager); err != nil {
					return errors.Wrap(err, "failed to provision google tag manager")
				}
			}

			if cfg.GoogleAnalytics.Enabled && utils.Tag(googleanaylticsprovider.Tag, tags) {
				l.Info("🅿️ Running provider", "name", googleanaylticsprovider.Name, "tag", googleanaylticsprovider.Tag)
				if err := googleanaylticsprovider.Web(cmd.Context(), tm, cfg.GoogleAnalytics); err != nil {
					return errors.Wrap(err, "failed to provision google analytics provider")
				}
			}

			if cfg.Emarsys.Enabled && utils.Tag(emarsysprovider.Tag, tags) {
				l.Info("🅿️ Running provider", "name", emarsysprovider.Name, "tag", emarsysprovider.Tag)
				if err := emarsysprovider.Web(cmd.Context(), tm, cfg.Emarsys); err != nil {
					return errors.Wrap(err, "failed to provision emarsys provider")
				}
			}

			if cfg.Hotjar.Enabled && utils.Tag(hotjarprovider.Tag, tags) {
				l.Info("🅿️ Running provider", "name", hotjarprovider.Name, "tag", hotjarprovider.Tag)
				if err := hotjarprovider.Web(cmd.Context(), tm, cfg.Hotjar); err != nil {
					return errors.Wrap(err, "failed to provision hotjar provider")
				}
			}

			if cfg.Criteo.Enabled && utils.Tag(criteoprovider.Tag, tags) {
				l.Info("🅿️ Running provider", "name", criteoprovider.Name, "tag", criteoprovider.Tag)
				if err := criteoprovider.Web(cmd.Context(), l, tm, cfg.Criteo); err != nil {
					return errors.Wrap(err, "failed to provision criteo provider")
				}
			}

			if cfg.Cookiebot.Enabled && utils.Tag(cookiebotprovider.Tag, tags) {
				l.Info("🅿️ Running provider", "name", cookiebotprovider.Name, "tag", cookiebotprovider.Tag)
				if err := cookiebotprovider.Web(cmd.Context(), tm, cfg.Cookiebot); err != nil {
					return errors.Wrap(err, "failed to provision cookiebot provider")
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
