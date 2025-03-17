package list

import (
	pkgcmd "github.com/foomo/sesamy-cli/pkg/cmd"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	"github.com/spf13/cobra"
)

// NewServer represents the server command
func NewServer(root *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "List Google Tag Manager Server Container",
		Args:  cobra.OnlyValidArgs,
		ValidArgs: []cobra.Completion{
			"built-in-variables",
			"clients",
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
			l := pkgcmd.Logger()
			l.Info("☕ Listing Server Container resources: " + resource)

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

			if err := tm.EnsureWorkspaceID(cmd.Context()); err != nil {
				return err
			}

			return list(cmd.Context(), l, tm, resource)
		},
	}

	root.AddCommand(cmd)
}
