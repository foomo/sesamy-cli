package cmd

import (
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
)

// tagmanagerServerCmd represents the server command
var tagmanagerServerCmd = &cobra.Command{
	Use:               "server",
	Short:             "Provision Google Tag Manager Server Container",
	PersistentPreRunE: preRunReadConfig,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		var opt option.ClientOption
		if cfg.Google.CredentialsFile != "" {
			opt = option.WithCredentialsFile(cfg.Google.CredentialsFile)
		} else {
			opt = option.WithCredentialsFile(cfg.Google.CredentialsJSON)
		}

		c, err := tagmanager.NewClient(
			cfg.Google.GTM.AccountID,
			cfg.Google.GTM.Server.ContainerID,
			cfg.Google.GTM.Server.WorkspaceID,
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

		logger.Info("- GTM client:", logger.Args("name", "Google Tag Manager Web Container"))
		if _, err := c.UpsertGTMClient(ctx, "Google Tag Manager Web Container", cfg.Google.GTM.Web.MeasurementID); err != nil {
			return err
		}

		logger.Info("- GA4 client:", logger.Args("name", "Google Analytics GA4"))
		ga4Client, err := c.UpsertGA4Client(ctx, "Google Analytics GA4")
		if err != nil {
			return err
		}

		logger.Info("- GA4 trigger:", logger.Args("name", "Google Analytics GA4"))
		ga4ClientTrigger, err := c.UpsertClientTrigger(ctx, "ga4", ga4Client.Name)
		if err != nil {
			return err
		}

		logger.Info("- GA4 tag:", logger.Args("name", "Google Analytics GA4"))
		if _, err := c.UpsertGA4ServerTag(ctx, "Google Analytics GA4", measurementID, ga4ClientTrigger); err != nil {
			return err
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
