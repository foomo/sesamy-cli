package cmd

import (
	"log/slog"
	"maps"
	"slices"

	"github.com/foomo/sesamy-cli/pkg/provider/conversionlinker"
	"github.com/foomo/sesamy-cli/pkg/provider/cookiebot"
	"github.com/foomo/sesamy-cli/pkg/provider/criteo"
	"github.com/foomo/sesamy-cli/pkg/provider/emarsys"
	"github.com/foomo/sesamy-cli/pkg/provider/facebook"
	"github.com/foomo/sesamy-cli/pkg/provider/googleads"
	"github.com/foomo/sesamy-cli/pkg/provider/googleanalytics"
	"github.com/foomo/sesamy-cli/pkg/provider/googleconsent"
	"github.com/foomo/sesamy-cli/pkg/provider/googletag"
	"github.com/foomo/sesamy-cli/pkg/provider/googletagmanager"
	"github.com/foomo/sesamy-cli/pkg/provider/hotjar"
	"github.com/foomo/sesamy-cli/pkg/provider/microsoftads"
	"github.com/foomo/sesamy-cli/pkg/provider/mixpanel"
	"github.com/foomo/sesamy-cli/pkg/provider/pinterest"
	"github.com/foomo/sesamy-cli/pkg/provider/tracify"
	"github.com/foomo/sesamy-cli/pkg/provider/umami"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// NewTags represents the tags command
func NewTags(l *slog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tags",
		Short: "Print out all available tags",
		RunE: func(cmd *cobra.Command, args []string) error {
			tags := map[string]string{
				conversionlinker.Name: conversionlinker.Tag,
				cookiebot.Name:        cookiebot.Tag,
				criteo.Name:           criteo.Tag,
				emarsys.Name:          emarsys.Tag,
				facebook.Name:         facebook.Tag,
				googleads.Name:        googleads.Tag,
				googleanalytics.Name:  googleanalytics.Tag,
				googletag.Name:        googletag.Tag,
				googleconsent.Name:    googleconsent.Tag,
				googletagmanager.Name: googletagmanager.Tag,
				hotjar.Name:           hotjar.Tag,
				microsoftads.Name:     microsoftads.Tag,
				mixpanel.Name:         mixpanel.Tag,
				tracify.Name:          tracify.Tag,
				umami.Name:            umami.Tag,
				pinterest.Name:        pinterest.Tag,
			}
			// Define the data for the first table
			data := pterm.TableData{{"Name", "Tag"}}
			for _, name := range slices.Sorted(maps.Keys(tags)) {
				data = append(data, []string{name, tags[name]})
			}
			return pterm.DefaultTable.WithHasHeader().WithData(data).Render()
		},
	}

	return cmd
}
