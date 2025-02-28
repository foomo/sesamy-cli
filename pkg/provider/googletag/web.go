package googletag

import (
	"github.com/foomo/gocontemplate/pkg/contemplate"
	"github.com/foomo/sesamy-cli/pkg/config"
	webtag "github.com/foomo/sesamy-cli/pkg/provider/googletag/web/tag"
	containervariable "github.com/foomo/sesamy-cli/pkg/provider/googletag/web/variable"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	commontrigger "github.com/foomo/sesamy-cli/pkg/tagmanager/common/trigger"
	commonvariable "github.com/foomo/sesamy-cli/pkg/tagmanager/common/variable"
	"github.com/foomo/sesamy-cli/pkg/tagmanager/web/variable"
	"github.com/foomo/sesamy-cli/pkg/utils"
	api "google.golang.org/api/tagmanager/v2"
)

func Web(tm *tagmanager.TagManager, cfg config.GoogleTag) error {
	folder, err := tm.UpsertFolder("Sesamy - " + Name)
	if err != nil {
		return err
	}

	{ // setup google tag
		configSettings := map[string]string{
			"server_container_url": "https://{{Page Hostname}}",
		}
		if !cfg.SendPageView {
			configSettings["send_page_view"] = "false"
		}

		eventSettings := map[string]*api.Variable{}
		for k, v := range cfg.DataLayerVariables {
			dlv, err := tm.UpsertVariable(folder, variable.NewDataLayer(v))
			if err != nil {
				return err
			}
			eventSettings[k] = dlv
		}

		tagID, err := tm.UpsertVariable(folder, commonvariable.NewConstant(NameGoogleTagID, cfg.TagID))
		if err != nil {
			return err
		}

		settingsVariable, err := tm.UpsertVariable(folder, containervariable.NewGoogleTagConfigurationSettings(NameGoogleTagSettings, configSettings))
		if err != nil {
			return err
		}
		if _, err = tm.UpsertTag(folder, webtag.NewGoogleTag(NameGoogleTag, tagID, settingsVariable, eventSettings)); err != nil {
			return err
		}
	}

	return nil
}

func CreateWebEventTriggers(tm *tagmanager.TagManager, cfg contemplate.Config) (map[string]map[string]string, error) {
	folder, err := tm.LookupFolder("Sesamy - " + Name)
	if err != nil {
		return nil, err
	}

	eventParameters, err := utils.LoadEventParams(cfg)
	if err != nil {
		return nil, err
	}

	for event, parameters := range eventParameters {
		if _, err = tm.UpsertTrigger(folder, commontrigger.NewEvent(event)); err != nil {
			return nil, err
		}

		variables, err := CreateWebDatalayerVariables(tm, parameters)
		if err != nil {
			return nil, err
		}

		if _, err := tm.UpsertVariable(folder, containervariable.NewGoogleTagEventSettings(event, variables)); err != nil {
			return nil, err
		}
	}

	return eventParameters, nil
}

func CreateWebDatalayerVariables(tm *tagmanager.TagManager, parameters map[string]string) (map[string]*api.Variable, error) {
	folder, err := tm.LookupFolder("Sesamy - " + Name)
	if err != nil {
		return nil, err
	}

	variables := make(map[string]*api.Variable, len(parameters))
	for parameterName, parameterValue := range parameters {
		if variables[parameterName], err = tm.UpsertVariable(folder, variable.NewDataLayer(parameterValue)); err != nil {
			return nil, err
		}
	}
	return variables, nil
}
