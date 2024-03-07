package cmd

import (
	"github.com/foomo/sesamy-cli/internal"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
)

// tagmanagerWebCmd represents the web command
var tagmanagerWebCmd = &cobra.Command{
	Use:               "web",
	Short:             "Provision Google Tag Manager Web Container",
	PersistentPreRunE: preRunReadConfig,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		var opt option.ClientOption
		if cfg.Google.CredentialsFile != "" {
			opt = option.WithCredentialsFile(cfg.Google.CredentialsFile)
		} else {
			opt = option.WithCredentialsFile(cfg.Google.CredentialsJSON)
		}

		eventParameters, err := internal.GetEventParameters(cfg)
		if err != nil {
			return err
		}

		c, err := tagmanager.NewClient(
			cfg.Google.GTM.AccountID,
			cfg.Google.GTM.Web.ContainerID,
			cfg.Google.GTM.Web.WorkspaceID,
			cfg.Google.GA4.MeasurementID,
			tagmanager.ClientWithClientOptions(opt),
		)
		if err != nil {
			return err
		}

		logger.Info("- Folder:", logger.Args("name", c.FolderName()))
		if _, err := c.UpsertFolder(ctx, c.FolderName()); err != nil {
			return err
		}

		logger.Info("- Variable:", logger.Args("name", "ga4-measurement-id"))
		measurementID, err := c.UpsertConstantVariable(ctx, "ga4-measurement-id", c.MeasurementID())
		if err != nil {
			return err
		}

		for key, value := range eventParameters {
			logger.Info("- GA4 Trigger:", logger.Args("name", key, "parameters", value))
			trigger, err := c.UpsertCustomEventTrigger(ctx, key)
			if err != nil {
				return err
			}
			logger.Info("- GA4 Tag:", logger.Args("name", key, "parameters", value))
			if _, err := c.UpsertGA4WebTag(ctx, key, value, measurementID, trigger); err != nil {
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
