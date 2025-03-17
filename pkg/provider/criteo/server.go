package criteo

import (
	"context"
	"log/slog"

	"github.com/foomo/sesamy-cli/pkg/config"
	servertagx "github.com/foomo/sesamy-cli/pkg/provider/criteo/server/tag"
	"github.com/foomo/sesamy-cli/pkg/provider/criteo/server/trigger"
	"github.com/foomo/sesamy-cli/pkg/provider/googleconsent"
	googleconsentvariable "github.com/foomo/sesamy-cli/pkg/provider/googleconsent/server/variable"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	commonvariable "github.com/foomo/sesamy-cli/pkg/tagmanager/common/variable"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"github.com/pkg/errors"
)

func Server(ctx context.Context, l *slog.Logger, tm *tagmanager.TagManager, cfg config.Criteo) error {
	folder, err := tm.UpsertFolder(ctx, "Sesamy - "+Name)
	if err != nil {
		return err
	}

	template, err := tm.LookupTemplate(ctx, NameCriteoEventsAPITemplate)
	if err != nil {
		if errors.Is(err, tagmanager.ErrNotFound) {
			l.Warn("Please install the 'Criteo Events API' template manually first")
		}
		return err
	}

	{ // create tags
		callerID, err := tm.UpsertVariable(ctx, folder, commonvariable.NewConstant(NameCallerID, cfg.CallerID))
		if err != nil {
			return err
		}

		partnerID, err := tm.UpsertVariable(ctx, folder, commonvariable.NewConstant(NamePartnerID, cfg.PartnerID))
		if err != nil {
			return err
		}

		applicationID, err := tm.UpsertVariable(ctx, folder, commonvariable.NewConstant(NameApplicationID, cfg.ApplicationID))
		if err != nil {
			return err
		}

		eventParameters, err := utils.LoadEventParams(ctx, cfg.ServerContainer)
		if err != nil {
			return err
		}

		for event := range eventParameters {
			var eventTriggerOpts []trigger.CriteoEventOption
			if cfg.GoogleConsent.Enabled {
				if err := googleconsent.ServerEnsure(ctx, tm); err != nil {
					return err
				}
				consentVariable, err := tm.LookupVariable(ctx, googleconsentvariable.GoogleConsentModeName(cfg.GoogleConsent.Mode))
				if err != nil {
					return err
				}
				eventTriggerOpts = append(eventTriggerOpts, trigger.CriteoEventWithConsentMode(consentVariable))
			}

			eventTrigger, err := tm.UpsertTrigger(ctx, folder, trigger.NewCriteoEvent(event, eventTriggerOpts...))
			if err != nil {
				return errors.Wrap(err, "failed to upsert event trigger: "+event)
			}

			if _, err := tm.UpsertTag(ctx, folder, servertagx.NewEventsAPITag(event, callerID, partnerID, applicationID, template, eventTrigger)); err != nil {
				return err
			}
		}
	}

	return nil
}
