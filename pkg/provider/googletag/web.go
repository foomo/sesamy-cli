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
	{ // create folder
		if folder, err := tm.UpsertFolder("Sesamy - " + Name); err != nil {
			return err
		} else {
			tm.SetFolderName(folder.Name)
		}
	}

	{ // setup google tag
		settings := map[string]string{
			"server_container_url": "https://{{Page Hostname}}",
		}
		if cfg.DebugMode {
			settings["debug_mode"] = "true"
		}
		if !cfg.SendPageView {
			settings["send_page_view"] = "false"
		}

		tagID, err := tm.UpsertVariable(commonvariable.NewConstant(NameGoogleTagID, cfg.TagID))
		if err != nil {
			return err
		}

		settingsVariable, err := tm.UpsertVariable(containervariable.NewGoogleTagConfigurationSettings(NameGoogleTagSettings, settings))
		if err != nil {
			return err
		}
		if _, err = tm.UpsertTag(webtag.NewGoogleTag(NameGoogleTag, tagID, settingsVariable)); err != nil {
			return err
		}
	}

	return nil
}

func CreateWebEventTriggers(tm *tagmanager.TagManager, cfg contemplate.Config) (map[string]map[string]string, error) {
	previousFolderName := tm.FolderName()
	tm.SetFolderName("Sesamy - " + Name)
	defer tm.SetFolderName(previousFolderName)

	eventParameters, err := utils.LoadEventParams(cfg)
	if err != nil {
		return nil, err
	}

	for event, parameters := range eventParameters {
		if _, err = tm.UpsertTrigger(commontrigger.NewEvent(event)); err != nil {
			return nil, err
		}

		variables := make(map[string]*api.Variable, len(parameters))
		for parameterName, parameterValue := range parameters {
			if variables[parameterName], err = tm.UpsertVariable(variable.NewDataLayerVariable(parameterValue)); err != nil {
				return nil, err
			}
		}

		if _, err := tm.UpsertVariable(containervariable.NewGoogleTagEventSettings(event, variables)); err != nil {
			return nil, err
		}
	}

	return eventParameters, nil
}
