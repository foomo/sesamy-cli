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
		var clientCredentialsOption option.ClientOption
		if cfg.Google.CredentialsFile != "" {
			clientCredentialsOption = option.WithCredentialsFile(cfg.Google.CredentialsFile)
		} else {
			clientCredentialsOption = option.WithCredentialsJSON([]byte(cfg.Google.CredentialsJSON))
		}

		c, err := tagmanager.NewClient(
			cmd.Context(),
			cfg.Google.GTM.AccountID,
			cfg.Google.GTM.Server.ContainerID,
			cfg.Google.GTM.Server.WorkspaceID,
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

		logger.Info("- GTM client:", logger.Args("name", "Google Tag Manager Web Container"))
		if _, err := c.UpsertGTMClient("Google Tag Manager Web Container", cfg.Google.GTM.Web.MeasurementID); err != nil {
			return err
		}

		logger.Info("- GA4 client:", logger.Args("name", "Google Analytics GA4"))
		ga4Client, err := c.UpsertGA4Client("Google Analytics GA4")
		if err != nil {
			return err
		}

		logger.Info("- GA4 trigger:", logger.Args("name", "Google Analytics GA4"))
		ga4ClientTrigger, err := c.UpsertClientTrigger("ga4", ga4Client)
		if err != nil {
			return err
		}

		logger.Info("- GA4 tag:", logger.Args("name", "Google Analytics GA4"))
		if _, err := c.UpsertGA4ServerTag("Google Analytics GA4", measurementID, ga4ClientTrigger); err != nil {
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
