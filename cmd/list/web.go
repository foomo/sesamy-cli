package list

import (
	"fmt"
	"log/slog"

	pkgcmd "github.com/foomo/sesamy-cli/pkg/cmd"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewWeb represents the web command
func NewWeb(l *slog.Logger) *cobra.Command {
	c := viper.New()

	cmd := &cobra.Command{
		Use:   "web",
		Short: "List Google Tag Manager Web Container",
		Args:  cobra.OnlyValidArgs,
		ValidArgs: []cobra.Completion{
			"built-in-variables",
			"environments",
			"folders",
			"gtag-config",
			"status",
			"tags",
			"templates",
			"templates-data",
			"transformations",
			"triggers",
			"variables",
			"workspaces",
			"zones",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			resource := args[0]
			l.Info("☕ Listing Web Container resources: " + resource)

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

			out, err := list(cmd.Context(), l, tm, resource)
			if err != nil {
				return err
			}

			if !c.GetBool("raw") {
				out = utils.Highlight(out)
			}
			_, err = fmt.Println(out)
			return err
		},
	}

	flags := cmd.Flags()

	flags.Bool("raw", false, "print raw output")
	_ = c.BindPFlag("raw", flags.Lookup("raw"))

	flags.StringSliceP("config", "c", []string{"sesamy.yaml"}, "config files (default is sesamy.yaml)")
	_ = c.BindPFlag("config", flags.Lookup("config"))

	return cmd
}
