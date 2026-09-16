package openaiads

import (
	"context"
	"log/slog"

	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/provider/googleconsent"
	googleconsentvariable "github.com/foomo/sesamy-cli/pkg/provider/googleconsent/server/variable"
	servertagx "github.com/foomo/sesamy-cli/pkg/provider/openaiads/server/tag"
	"github.com/foomo/sesamy-cli/pkg/provider/openaiads/server/trigger"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	commontemplate "github.com/foomo/sesamy-cli/pkg/tagmanager/common/template"
	commonvariable "github.com/foomo/sesamy-cli/pkg/tagmanager/common/variable"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"github.com/pkg/errors"
	tagmanager2 "google.golang.org/api/tagmanager/v2"
)

func Server(ctx context.Context, l *slog.Logger, tm *tagmanager.TagManager, cfg config.OpenAIAds) error {
	folder, err := tm.UpsertFolder(ctx, "Sesamy - "+Name)
	if err != nil {
		return err
	}

	pixelID, err := tm.UpsertVariable(ctx, folder, commonvariable.NewConstant(NamePixelIDConstant, cfg.PixelID))
	if err != nil {
		return err
	}

	apiKey, err := tm.UpsertVariable(ctx, folder, commonvariable.NewConstant(NameAPIKeyConstant, cfg.APIKey))
	if err != nil {
		return err
	}

	if cfg.Conversion.Enabled {
		tagTemplate, err := commontemplate.ResolveOrLookup(ctx, tm, cfg.Templates.ConversionsAPITag, NameConversionsAPITagTemplate, func() (*tagmanager2.CustomTemplate, error) {
			return tm.LookupTemplate(ctx, NameConversionsAPITagTemplate)
		})
		if err != nil {
			if errors.Is(err, tagmanager.ErrNotFound) {
				l.Warn("Please install the '" + NameConversionsAPITagTemplate + "' Tag Template manually first")
			}

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
					if err := googleconsent.ServerEnsure(ctx, tm); err != nil {
						return err
					}

					consentVariable, err := tm.LookupVariable(ctx, googleconsentvariable.GoogleConsentModeName(cfg.GoogleConsent.Mode))
					if err != nil {
						return err
					}

					eventTriggerOpts = append(eventTriggerOpts, trigger.ConversionEventWithConsentMode(consentVariable))
				}

				eventTrigger, err := tm.UpsertTrigger(ctx, folder, trigger.NewConversionEvent(event, eventTriggerOpts...))
				if err != nil {
					return errors.Wrap(err, "failed to upsert event trigger: "+event)
				}

				if _, err := tm.UpsertTag(ctx, folder, servertagx.NewConversion(event, pixelID, apiKey, tagTemplate, cfg.Conversion.ValidateOnly, cfg.Conversion.ServerContainer.Setting(event), eventTrigger)); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
