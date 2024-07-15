package umami

import (
	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/provider/googletag"
	containertag "github.com/foomo/sesamy-cli/pkg/provider/umami/server/tag"
	containertemplate "github.com/foomo/sesamy-cli/pkg/provider/umami/server/template"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	commontrigger "github.com/foomo/sesamy-cli/pkg/tagmanager/common/trigger"
	"github.com/pkg/errors"
)

func Server(tm *tagmanager.TagManager, cfg config.Umami) error {
	{ // create folder
		if folder, err := tm.UpsertFolder("Sesamy - " + Name); err != nil {
			return err
		} else {
			tm.SetFolderName(folder.Name)
		}
	}

	template, err := tm.UpsertCustomTemplate(containertemplate.NewUmami(Name))
	if err != nil {
		return err
	}

	{ // create tags
		eventParameters, err := googletag.CreateServerEventTriggers(tm, cfg.ServerContainer)
		if err != nil {
			return err
		}

		for event := range eventParameters {
			eventTrigger, err := tm.LookupTrigger(commontrigger.EventName(event))
			if err != nil {
				return errors.Wrap(err, "failed to lookup event trigger: "+event)
			}

			if _, err := tm.UpsertTag(containertag.NewUmami(event, cfg, template, eventTrigger)); err != nil {
				return err
			}
		}
	}

	return nil
}
