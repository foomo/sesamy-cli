package tagmanager

import (
	conversionlinkerprovider "github.com/foomo/sesamy-cli/pkg/provider/conversionlinker"
	criteoprovider "github.com/foomo/sesamy-cli/pkg/provider/criteo"
	emarsysprovider "github.com/foomo/sesamy-cli/pkg/provider/emarsys"
	facebookprovider "github.com/foomo/sesamy-cli/pkg/provider/facebook"
	googleadsprovider "github.com/foomo/sesamy-cli/pkg/provider/googleads"
	googleanalyticsprovider "github.com/foomo/sesamy-cli/pkg/provider/googleanalytics"
	googletagmanagerprovider "github.com/foomo/sesamy-cli/pkg/provider/googletagmanager"
	microsoftadsprovider "github.com/foomo/sesamy-cli/pkg/provider/microsoftads"
	tracifyprovider "github.com/foomo/sesamy-cli/pkg/provider/tracify"
	umamiprovider "github.com/foomo/sesamy-cli/pkg/provider/umami"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// NewTags represents the tags command
func NewTags(root *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tags",
		Short: "Print out all available tags",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Define the data for the first table
			data := pterm.TableData{
				{"Name", "Tag"},
				{conversionlinkerprovider.Name, conversionlinkerprovider.Tag},
				{criteoprovider.Name, criteoprovider.Tag},
				{emarsysprovider.Name, emarsysprovider.Tag},
				{facebookprovider.Name, facebookprovider.Tag},
				{googleadsprovider.Name, googleadsprovider.Tag},
				{googleanalyticsprovider.Name, googleanalyticsprovider.Tag},
				{googletagmanagerprovider.Name, googletagmanagerprovider.Tag},
				{microsoftadsprovider.Name, microsoftadsprovider.Tag},
				{tracifyprovider.Name, tracifyprovider.Tag},
				{umamiprovider.Name, umamiprovider.Tag},
			}

			return pterm.DefaultTable.WithHasHeader().WithData(data).Render()
		},
	}

	root.AddCommand(cmd)

	return cmd
}
