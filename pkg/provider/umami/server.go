package umami

import (
	"context"

	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/provider/googleconsent"
	googleconsentvariable "github.com/foomo/sesamy-cli/pkg/provider/googleconsent/server/variable"
	containertag "github.com/foomo/sesamy-cli/pkg/provider/umami/server/tag"
	containertemplate "github.com/foomo/sesamy-cli/pkg/provider/umami/server/template"
	"github.com/foomo/sesamy-cli/pkg/provider/umami/server/trigger"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"github.com/pkg/errors"
)

func Server(ctx context.Context, tm *tagmanager.TagManager, cfg config.Umami) error {
	folder, err := tm.UpsertFolder(ctx, "Sesamy - "+Name)
	if err != nil {
		return err
	}

	template, err := tm.UpsertCustomTemplate(ctx, containertemplate.NewUmami(Name))
	if err != nil {
		return err
	}

	{ // create tags
		eventParameters, err := utils.LoadEventParams(ctx, cfg.ServerContainer)
		if err != nil {
			return err
		}

		for event := range eventParameters {
			var eventTriggerOpts []trigger.UmamiEventOption
			if cfg.GoogleConsent.Enabled {
				if err := googleconsent.ServerEnsure(ctx, tm); err != nil {
					return err
				}
				consentVariable, err := tm.LookupVariable(ctx, googleconsentvariable.GoogleConsentModeName(cfg.GoogleConsent.Mode))
				if err != nil {
					return err
				}
				eventTriggerOpts = append(eventTriggerOpts, trigger.UmamiEventWithConsentMode(consentVariable))
			}

			eventTrigger, err := tm.UpsertTrigger(ctx, folder, trigger.NewUmamiEvent(event, eventTriggerOpts...))
			if err != nil {
				return errors.Wrap(err, "failed to upsert event trigger: "+event)
			}

			if _, err := tm.UpsertTag(ctx, folder, containertag.NewUmami(event, cfg, template, eventTrigger)); err != nil {
				return err
			}
		}
	}

	return nil
}
