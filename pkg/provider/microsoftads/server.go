package microsoftads

import (
	"context"
	"log/slog"

	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/provider/googleconsent"
	googleconsentvariable "github.com/foomo/sesamy-cli/pkg/provider/googleconsent/server/variable"
	servertagx "github.com/foomo/sesamy-cli/pkg/provider/microsoftads/server/tag"
	"github.com/foomo/sesamy-cli/pkg/provider/microsoftads/server/template"
	"github.com/foomo/sesamy-cli/pkg/provider/microsoftads/server/trigger"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	commonvariable "github.com/foomo/sesamy-cli/pkg/tagmanager/common/variable"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"github.com/pkg/errors"
)

func Server(ctx context.Context, l *slog.Logger, tm *tagmanager.TagManager, cfg config.MicrosoftAds) error {
	folder, err := tm.UpsertFolder("Sesamy - " + Name)
	if err != nil {
		return err
	}

	tagID, err := tm.UpsertVariable(folder, commonvariable.NewConstant(NameTagIDConstant, cfg.TagID))
	if err != nil {
		return err
	}

	if cfg.Conversion.Enabled {
		tagTemplate, err := tm.UpsertCustomTemplate(template.NewConversionTag(NameConversionsTagTemplate))
		if err != nil {
			return err
		}

		{ // create tags
			eventParameters, err := utils.LoadEventParams(ctx, cfg.Conversion.ServerContainer.Config)
			if err != nil {
				return err
			}

			for event := range eventParameters {
				var eventTriggerOpts []trigger.ConversionEventOption
				if cfg.GoogleConsent.Enabled {
					if err := googleconsent.ServerEnsure(tm); err != nil {
						return err
					}
					consentVariable, err := tm.LookupVariable(googleconsentvariable.GoogleConsentModeName(cfg.GoogleConsent.Mode))
					if err != nil {
						return err
					}
					eventTriggerOpts = append(eventTriggerOpts, trigger.ConversionEventWithConsentMode(consentVariable))
				}

				eventTrigger, err := tm.UpsertTrigger(folder, trigger.NewConversionEvent(event, eventTriggerOpts...))
				if err != nil {
					return errors.Wrap(err, "failed to upsert event trigger: "+event)
				}

				if _, err := tm.UpsertTag(folder, servertagx.NewConversion(event, tagID, tagTemplate, cfg.Conversion.ServerContainer.Setting(event), eventTrigger)); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
