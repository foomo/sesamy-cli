package cmd

import (
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	"github.com/foomo/sesamy-cli/pkg/tagmanager/client"
	client2 "github.com/foomo/sesamy-cli/pkg/tagmanager/tag"
	"github.com/foomo/sesamy-cli/pkg/tagmanager/trigger"
	"github.com/foomo/sesamy-cli/pkg/tagmanager/variable"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
	tagmanager2 "google.golang.org/api/tagmanager/v2"
)

// tagmanagerServerCmd represents the server command
// TODO flags workspace, dry, diff
// TODO google user auth
var tagmanagerServerCmd = &cobra.Command{
	Use:               "server",
	Short:             "Provision Google Tag Manager Server Container",
	PersistentPreRunE: preRunReadConfig,
	RunE: func(cmd *cobra.Command, args []string) error {
		var clientCredentialsOption option.ClientOption
		if cfg.Google.CredentialsFile != "" {
			clientCredentialsOption = option.WithCredentialsFile(cfg.Google.CredentialsFile)
		} else {
			clientCredentialsOption = option.WithCredentialsJSON([]byte(cfg.Google.CredentialsJSON))
		}

		c, err := tagmanager.NewClient(
			cmd.Context(),
			logger,
			cfg.Google.GTM.AccountID,
			cfg.Google.GTM.Server.ContainerID,
			cfg.Google.GTM.Server.WorkspaceID,
			cfg.Google.GA4.MeasurementID,
			tagmanager.ClientWithRequestQuota(cfg.Google.RequestQuota),
			tagmanager.ClientWithClientOptions(clientCredentialsOption),
		)
		if err != nil {
			return err
		}

		p := cfg.Tagmanager.Prefixes

		{
			name := p.FolderName(c.FolderName())
			if _, err := c.UpsertFolder(name); err != nil {
				return err
			}
		}

		{
			if _, err := c.EnableBuiltInVariable("clientName"); err != nil {
				return err
			}
		}

		var ga4MeasurementID *tagmanager2.Variable
		{
			name := p.Variables.ConstantName("Google Analytics GA4 ID")
			if ga4MeasurementID, err = c.UpsertVariable(variable.NewConstant(name, c.MeasurementID())); err != nil {
				return err
			}
		}

		var webContainerMeasurementID *tagmanager2.Variable
		{
			name := p.Variables.ConstantName("Google Tag Mangager Web Container ID")
			if webContainerMeasurementID, err = c.UpsertVariable(variable.NewConstant(name, cfg.Google.GTM.Web.MeasurementID)); err != nil {
				return err
			}
		}

		{
			name := p.ClientName("Google Tag Manager Web Container")
			if _, err := c.UpsertClient(client.NewGTM(name, webContainerMeasurementID)); err != nil {
				return err
			}
		}

		var ga4Client *tagmanager2.Client
		{
			name := p.ClientName("Google Analytics GA4")
			if ga4Client, err = c.UpsertClient(client.NewGA4(name)); err != nil {
				return err
			}
		}

		var ga4ClientTrigger *tagmanager2.Trigger
		{
			name := p.Triggers.ClientName("Google Analytics GA4 Client")
			if ga4ClientTrigger, err = c.UpsertTrigger(trigger.NewClient(name, ga4Client)); err != nil {
				return err
			}
		}

		{
			name := p.Tags.ServerGA4EventName("Google Analytics GA4")
			if _, err := c.UpsertTag(client2.NewServerGA4Event(name, ga4MeasurementID, ga4ClientTrigger)); err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	tagmanagerCmd.AddCommand(tagmanagerServerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tagmanagerServerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tagmanagerServerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
