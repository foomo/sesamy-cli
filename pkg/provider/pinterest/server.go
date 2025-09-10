package pinterest

import (
	"context"
	"log/slog"

	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/provider/googleconsent"
	googleconsentvariable "github.com/foomo/sesamy-cli/pkg/provider/googleconsent/server/variable"
	containertag "github.com/foomo/sesamy-cli/pkg/provider/pinterest/server/tag"
	"github.com/foomo/sesamy-cli/pkg/provider/pinterest/server/trigger"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"github.com/pkg/errors"
)

func Server(ctx context.Context, l *slog.Logger, tm *tagmanager.TagManager, cfg config.Pinterest) error {
	folder, err := tm.UpsertFolder(ctx, "Sesamy - "+Name)
	if err != nil {
		return err
	}

	template, err := tm.LookupTemplate(ctx, NameTagTemplate)
	if err != nil {
		if errors.Is(err, tagmanager.ErrNotFound) {
			l.Warn("Please install the 'Pinterest API for Conversions Ta' Tag Template manually first")
		}
		return err
	}

	{ // create tags
		eventParameters, err := utils.LoadEventParams(ctx, cfg.ServerContainer)
		if err != nil {
			return err
		}

		for event := range eventParameters {
			var eventTriggerOpts []trigger.EventOption
			if cfg.GoogleConsent.Enabled {
				if err := googleconsent.ServerEnsure(ctx, tm); err != nil {
					return err
				}
				consentVariable, err := tm.LookupVariable(ctx, googleconsentvariable.GoogleConsentModeName(cfg.GoogleConsent.Mode))
				if err != nil {
					return err
				}
				eventTriggerOpts = append(eventTriggerOpts, trigger.ConversionWithConsentMode(consentVariable))
			}

			eventTrigger, err := tm.UpsertTrigger(ctx, folder, trigger.NewConversion(event, eventTriggerOpts...))
			if err != nil {
				return errors.Wrap(err, "failed to upsert event trigger: "+event)
			}

			if _, err := tm.UpsertTag(ctx, folder, containertag.NewConversion(event, cfg, template, eventTrigger)); err != nil {
				return err
			}
		}
	}

	return nil
}
