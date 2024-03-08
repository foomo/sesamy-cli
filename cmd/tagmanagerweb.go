package cmd

import (
	"github.com/foomo/sesamy-cli/internal"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
	tagmanager2 "google.golang.org/api/tagmanager/v2"
)

// tagmanagerWebCmd represents the web command
var tagmanagerWebCmd = &cobra.Command{
	Use:               "web",
	Short:             "Provision Google Tag Manager Web Container",
	PersistentPreRunE: preRunReadConfig,
	RunE: func(cmd *cobra.Command, args []string) error {
		var clientCredentialsOption option.ClientOption
		if cfg.Google.CredentialsFile != "" {
			clientCredentialsOption = option.WithCredentialsFile(cfg.Google.CredentialsFile)
		} else {
			clientCredentialsOption = option.WithCredentialsJSON([]byte(cfg.Google.CredentialsJSON))
		}

		eventParameters, err := internal.GetEventParameters(cfg.Tagmanager)
		if err != nil {
			return err
		}

		c, err := tagmanager.NewClient(
			cmd.Context(),
			cfg.Google.GTM.AccountID,
			cfg.Google.GTM.Web.ContainerID,
			cfg.Google.GTM.Web.WorkspaceID,
			cfg.Google.GA4.MeasurementID,
			tagmanager.ClientWithRequestQuota(15),
			tagmanager.ClientWithClientOptions(clientCredentialsOption),
		)
		if err != nil {
			return err
		}

		logger.Info("- Folder:", logger.Args("name", c.FolderName()))
		if _, err := c.UpsertFolder(c.FolderName()); err != nil {
			return err
		}

		logger.Info("- Variable:", logger.Args("name", "ga4-measurement-id"))
		measurementID, err := c.UpsertConstantVariable("ga4-measurement-id", c.MeasurementID())
		if err != nil {
			return err
		}

		for event, parameters := range eventParameters {
			logger.Info("- GA4 Event Trigger:", logger.Args("name", event))
			trigger, err := c.UpsertCustomEventTrigger(event)
			if err != nil {
				return err
			}

			eventSettingsVariables := make(map[string]*tagmanager2.Variable, len(parameters))
			for _, parameter := range parameters {
				logger.Info("- Event Model Variable:", logger.Args("name", parameter))
				eventSettingsVariables[parameter], err = c.UpsertEventModelVariable(parameter)
				if err != nil {
					return err
				}
			}

			logger.Info("- GT Event Settings Variable:", logger.Args("name", event))
			eventSettings, err := c.UpsertGTEventSettingsVariable(event, eventSettingsVariables)
			if err != nil {
				return err
			}

			logger.Info("- GA4 Tag:", logger.Args("name", event, "parameters", parameters))
			if _, err := c.UpsertGA4WebTag(event, eventSettings, measurementID, trigger); err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	tagmanagerCmd.AddCommand(tagmanagerWebCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tagmanagerWebCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tagmanagerWebCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
