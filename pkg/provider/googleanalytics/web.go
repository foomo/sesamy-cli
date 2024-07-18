package googleanalytics

import (
	"github.com/foomo/sesamy-cli/pkg/config"
	containertag "github.com/foomo/sesamy-cli/pkg/provider/googleanalytics/web/tag"
	"github.com/foomo/sesamy-cli/pkg/provider/googletag"
	commonvariable "github.com/foomo/sesamy-cli/pkg/provider/googletag/web/variable"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	commontrigger "github.com/foomo/sesamy-cli/pkg/tagmanager/common/trigger"
	"github.com/pkg/errors"
)

func Web(tm *tagmanager.TagManager, cfg config.GoogleAnalytics) error {
	{ // create folder
		if folder, err := tm.UpsertFolder("Sesamy - " + Name); err != nil {
			return err
		} else {
			tm.SetFolderName(folder.Name)
		}
	}

	{ // create event tags
		tagID, err := tm.LookupVariable(googletag.NameGoogleTagID)
		if err != nil {
			return err
		}

		eventParameters, err := googletag.CreateWebEventTriggers(tm, cfg.WebContainer)
		if err != nil {
			return err
		}

		for event := range eventParameters {
			eventTrigger, err := tm.LookupTrigger(commontrigger.EventName(event))
			if err != nil {
				return errors.Wrap(err, "failed to lookup event trigger: "+event)
			}

			eventSettings, err := tm.LookupVariable(commonvariable.GoogleTagEventSettingsName(event))
			if err != nil {
				return errors.Wrap(err, "failed to lookup google tag event setting: "+event)
			}

			if _, err := tm.UpsertTag(containertag.NewGoogleAnalyticsGA4Event(event, tagID, eventSettings, eventTrigger)); err != nil {
				return err
			}
		}
	}

	return nil
}
