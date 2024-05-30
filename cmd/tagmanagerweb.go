package cmd

import (
	"go/types"
	"strconv"

	"github.com/foomo/gocontemplate/pkg/assume"
	"github.com/foomo/gocontemplate/pkg/contemplate"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	client "github.com/foomo/sesamy-cli/pkg/tagmanager/tag"
	trigger2 "github.com/foomo/sesamy-cli/pkg/tagmanager/trigger"
	"github.com/foomo/sesamy-cli/pkg/tagmanager/variable"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/stoewer/go-strcase"
	"google.golang.org/api/option"
	tagmanager2 "google.golang.org/api/tagmanager/v2"
)

// NewTagmanagerWebCmd represents the web command
func NewTagmanagerWebCmd(root *cobra.Command) {
	getEventParams := func(obj types.Object) ([]string, error) {
		var ret []string
		if eventStruct := assume.T[*types.Struct](obj.Type().Underlying()); eventStruct != nil {
			for i := range eventStruct.NumFields() {
				if eventField := eventStruct.Field(i); eventField.Name() == "Params" {
					if paramsStruct := assume.T[*types.Struct](eventField.Type().Underlying()); paramsStruct != nil {
						for j := range paramsStruct.NumFields() {
							tag, err := utils.ParseTagName(paramsStruct.Tag(j))
							if err != nil {
								return nil, errors.Wrapf(err, "failed to parse tag `%s`", paramsStruct.Tag(j))
							}
							ret = append(ret, tag)
						}
					}
				}
			}
		}
		return ret, nil
	}

	cmd := &cobra.Command{
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

			parser, err := contemplate.Load(&cfg.Tagmanager.Config)
			if err != nil {
				return err
			}

			eventParameters := map[string][]string{}
			for _, pkgCfg := range cfg.Tagmanager.Packages {
				pkg := parser.Package(pkgCfg.Path)
				for _, event := range pkgCfg.Types {
					eventParams, err := getEventParams(pkg.LookupScopeType(event))
					if err != nil {
						return err
					}
					eventParameters[strcase.SnakeCase(event)] = eventParams
				}
			}

			c, err := tagmanager.NewClient(
				cmd.Context(),
				logger,
				cfg.Google.GTM.AccountID,
				cfg.Google.GTM.Web.ContainerID,
				cfg.Google.GTM.Web.WorkspaceID,
				cfg.Google.GA4.MeasurementID,
				tagmanager.ClientWithRequestQuota(cfg.Google.RequestQuota),
				tagmanager.ClientWithClientOptions(clientCredentialsOption),
			)
			if err != nil {
				return err
			}

			p := cfg.Tagmanager.Prefixes

			if _, err := c.UpsertFolder(p.FolderName(c.FolderName())); err != nil {
				return err
			}

			var ga4MeasurementID *tagmanager2.Variable
			{
				name := p.Variables.ConstantName("Google Analytics GA4 ID")
				if ga4MeasurementID, err = c.UpsertVariable(variable.NewConstant(name, c.MeasurementID())); err != nil {
					return err
				}
			}

			var serverContainerURL *tagmanager2.Variable
			{
				name := p.Variables.ConstantName("Server Container URL")
				if serverContainerURL, err = c.UpsertVariable(variable.NewConstant(name, cfg.Google.GT.ServerContainerURL)); err != nil {
					return err
				}
			}

			var googleTagSettings *tagmanager2.Variable
			{
				name := p.Variables.GTSettingsName("Google Tag")
				if googleTagSettings, err = c.UpsertVariable(variable.NewGTSettings(name, map[string]*tagmanager2.Variable{
					"server_container_url": serverContainerURL,
				})); err != nil {
					return err
				}
			}

			{
				name := p.Tags.GoogleTagName("Google Tag")
				if _, err = c.UpsertTag(client.NewGoogleTag(name, ga4MeasurementID, googleTagSettings, map[string]string{
					"enable_page_views": strconv.FormatBool(cfg.Google.GT.EnablePageViews),
				})); err != nil {
					return err
				}
			}

			for event, parameters := range eventParameters {
				var trigger *tagmanager2.Trigger
				{
					name := p.Triggers.CustomEventName(event)
					if trigger, err = c.UpsertTrigger(trigger2.NewCustomEvent(name, event)); err != nil {
						return err
					}
				}

				if cfg.Tagmanager.Tags.GA4Enabled {
					eventSettingsVariables := make(map[string]*tagmanager2.Variable, len(parameters))
					for _, parameter := range parameters {
						name := p.Variables.EventModelName(parameter)
						eventSettingsVariables[parameter], err = c.UpsertVariable(variable.NewEventModel(name, parameter))
						if err != nil {
							return err
						}
					}

					var eventSettings *tagmanager2.Variable
					{
						name := p.Variables.GTEventSettingsName(event)
						if eventSettings, err = c.UpsertVariable(variable.NewGTEventSettings(name, eventSettingsVariables)); err != nil {
							return err
						}
					}

					{
						name := p.Tags.GA4EventName(event)
						if _, err := c.UpsertTag(client.NewGA4Event(name, event, eventSettings, ga4MeasurementID, trigger)); err != nil {
							return err
						}
					}
				}
			}

			return nil
		},
	}

	root.AddCommand(cmd)
}
