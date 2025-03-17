package tracify

import (
	"context"
	"log/slog"

	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/provider/googleconsent"
	googleconsentvariable "github.com/foomo/sesamy-cli/pkg/provider/googleconsent/server/variable"
	servertagx "github.com/foomo/sesamy-cli/pkg/provider/tracify/server/tag"
	"github.com/foomo/sesamy-cli/pkg/provider/tracify/server/template"
	"github.com/foomo/sesamy-cli/pkg/provider/tracify/server/trigger"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	commonvariable "github.com/foomo/sesamy-cli/pkg/tagmanager/common/variable"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"github.com/pkg/errors"
)

func Server(ctx context.Context, l *slog.Logger, tm *tagmanager.TagManager, cfg config.Tracify) error {
	folder, err := tm.UpsertFolder(ctx, "Tracify - "+Name)
	if err != nil {
		return err
	}

	{ // conversion
		token, err := tm.UpsertVariable(ctx, folder, commonvariable.NewConstant(NameTokenConstant, cfg.Token))
		if err != nil {
			return err
		}

		customerSiteID, err := tm.UpsertVariable(ctx, folder, commonvariable.NewConstant(NameCustomerSiteIDConstant, cfg.CustomerSiteID))
		if err != nil {
			return err
		}

		tagTemplate, err := tm.UpsertCustomTemplate(ctx, template.NewTracifyTag(NameTracifyServerTagTemplate))
		if err != nil {
			return err
		}

		{ // create tags
			eventParameters, err := utils.LoadEventParams(ctx, cfg.ServerContainer)
			if err != nil {
				return err
			}

			for event := range eventParameters {
				var eventTriggerOpts []trigger.TracifyEventOption
				if cfg.GoogleConsent.Enabled {
					if err := googleconsent.ServerEnsure(ctx, tm); err != nil {
						return err
					}
					consentVariable, err := tm.LookupVariable(ctx, googleconsentvariable.GoogleConsentModeName(cfg.GoogleConsent.Mode))
					if err != nil {
						return err
					}
					eventTriggerOpts = append(eventTriggerOpts, trigger.TracifyEventWithConsentMode(consentVariable))
				}

				eventTrigger, err := tm.UpsertTrigger(ctx, folder, trigger.NewTracifyEvent(event, eventTriggerOpts...))
				if err != nil {
					return errors.Wrap(err, "failed to upsert event trigger: "+event)
				}

				if _, err := tm.UpsertTag(ctx, folder, servertagx.NewTracify(event, token, customerSiteID, tagTemplate, cfg, eventTrigger)); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
