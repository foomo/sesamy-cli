package cmd

import (
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	containervariable "github.com/foomo/sesamy-cli/pkg/tagmanager/common/variable"
	containerclient "github.com/foomo/sesamy-cli/pkg/tagmanager/server/client"
	containertag "github.com/foomo/sesamy-cli/pkg/tagmanager/server/tag"
	containertemplate "github.com/foomo/sesamy-cli/pkg/tagmanager/server/template"
	containertrigger "github.com/foomo/sesamy-cli/pkg/tagmanager/server/trigger"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
	tagmanager2 "google.golang.org/api/tagmanager/v2"
)

// NewTagManagerServerCmd represents the server command
// TODO flags workspace, dry, diff
// TODO google user auth
func NewTagManagerServerCmd(root *cobra.Command) {
	var cmd = &cobra.Command{
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

			var webContainerMeasurementID *tagmanager2.Variable
			{
				name := p.Variables.ConstantName("Google Tag Mangager Web Container ID")
				if webContainerMeasurementID, err = c.UpsertVariable(containervariable.NewConstant(name, cfg.Google.GTM.Web.MeasurementID)); err != nil {
					return err
				}
			}

			{
				name := p.ClientName("Google Tag Manager Web Container")
				if _, err := c.UpsertClient(containerclient.NewGTM(name, webContainerMeasurementID)); err != nil {
					return err
				}
			}

			// --- MPv2 Client ---
			var mpv2Client *tagmanager2.Client
			{
				name := p.ClientName("Measurement Protocol GA4")
				if mpv2Client, err = c.UpsertClient(containerclient.NewMPv2(name)); err != nil {
					return err
				}
			}

			var mpv2ClientTrigger *tagmanager2.Trigger
			{
				name := p.Triggers.ClientName("Measurement Protocol GA4 Client")
				if mpv2ClientTrigger, err = c.UpsertTrigger(containertrigger.NewClient(name, mpv2Client)); err != nil {
					return err
				}
			}

			// --- GA4 Client ---
			var ga4Client *tagmanager2.Client
			{
				name := p.ClientName("Google Analytics GA4")
				if ga4Client, err = c.UpsertClient(containerclient.NewGA4(name)); err != nil {
					return err
				}
			}

			var ga4ClientTrigger *tagmanager2.Trigger
			{
				name := p.Triggers.ClientName("Google Analytics GA4 Client")
				if ga4ClientTrigger, err = c.UpsertTrigger(containertrigger.NewClient(name, ga4Client)); err != nil {
					return err
				}
			}

			// --- Tags ---
			if cfg.Tagmanager.Tags.GA4.Enabled {
				var ga4MeasurementID *tagmanager2.Variable
				{
					name := p.Variables.ConstantName("Google Analytics GA4 ID")
					if ga4MeasurementID, err = c.UpsertVariable(containervariable.NewConstant(name, c.MeasurementID())); err != nil {
						return err
					}
				}
				name := p.Tags.ServerGA4EventName("Google Analytics GA4")
				if _, err := c.UpsertTag(containertag.NewGoogleAnalyticsGA4(name, ga4MeasurementID, ga4ClientTrigger, mpv2ClientTrigger)); err != nil {
					return err
				}
			}

			if cfg.Tagmanager.Tags.Umami.Enabled {
				var umamiTemplate *tagmanager2.CustomTemplate
				if umamiTemplate, err = c.UpsertCustomTemplate(containertemplate.NewUmami("Sesamy Umami")); err != nil {
					return err
				}
				if _, err := c.UpsertTag(containertag.NewUmami(
					"Umami",
					cfg.Tagmanager.Tags.Umami.WebsiteID,
					cfg.Tagmanager.Tags.Umami.Domain,
					cfg.Tagmanager.Tags.Umami.EndpointURL,
					umamiTemplate,
					ga4ClientTrigger,
					mpv2ClientTrigger,
				)); err != nil {
					return err
				}
			}

			return nil
		},
	}

	root.AddCommand(cmd)
}
