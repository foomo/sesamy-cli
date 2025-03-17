package conversionlinker

import (
	"context"

	"github.com/foomo/sesamy-cli/pkg/config"
	containertag "github.com/foomo/sesamy-cli/pkg/provider/conversionlinker/server/tag"
	"github.com/foomo/sesamy-cli/pkg/provider/conversionlinker/server/trigger"
	"github.com/foomo/sesamy-cli/pkg/provider/googleconsent"
	googleconsentvariable "github.com/foomo/sesamy-cli/pkg/provider/googleconsent/server/variable"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	"github.com/pkg/errors"
)

func Server(ctx context.Context, tm *tagmanager.TagManager, cfg config.ConversionLinker) error {
	folder, err := tm.UpsertFolder(ctx, "Sesamy - "+Name)
	if err != nil {
		return err
	}

	var eventTriggerOpts []trigger.ConversionLinkerEventOption
	if cfg.GoogleConsent.Enabled {
		if err := googleconsent.ServerEnsure(ctx, tm); err != nil {
			return err
		}
		consentVariable, err := tm.LookupVariable(ctx, googleconsentvariable.GoogleConsentModeName(cfg.GoogleConsent.Mode))
		if err != nil {
			return err
		}
		eventTriggerOpts = append(eventTriggerOpts, trigger.ConversionLinkerEventWithConsentMode(consentVariable))
	}

	eventTrigger, err := tm.UpsertTrigger(ctx, folder, trigger.NewConversionLinkerEvent(NameConversionLinkerTrigger, eventTriggerOpts...))
	if err != nil {
		return errors.Wrap(err, "failed to upsert event trigger: "+NameConversionLinkerTrigger)
	}

	if _, err := tm.UpsertTag(ctx, folder, containertag.NewConversionLinker(Name, eventTrigger)); err != nil {
		return err
	}

	return nil
}
