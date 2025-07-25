package googleads

import (
	"context"
	"log/slog"

	"github.com/foomo/sesamy-cli/pkg/config"
	servertagx "github.com/foomo/sesamy-cli/pkg/provider/googleads/server/tag"
	"github.com/foomo/sesamy-cli/pkg/provider/googleads/server/trigger"
	"github.com/foomo/sesamy-cli/pkg/provider/googleconsent"
	googleconsentvariable "github.com/foomo/sesamy-cli/pkg/provider/googleconsent/server/variable"
	"github.com/foomo/sesamy-cli/pkg/provider/googletagmanager"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	commonvariable "github.com/foomo/sesamy-cli/pkg/tagmanager/common/variable"
	"github.com/foomo/sesamy-cli/pkg/tagmanager/server/variable"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"github.com/pkg/errors"
)

func Server(ctx context.Context, l *slog.Logger, tm *tagmanager.TagManager, cfg config.GoogleAds) error {
	folder, err := Folder(ctx, tm)
	if err != nil {
		return err
	}

	gtmFolder, err := googletagmanager.Folder(ctx, tm)
	if err != nil {
		return err
	}

	conversionID, err := tm.UpsertVariable(ctx, folder, commonvariable.NewConstant(NameConversionIDConstant, cfg.ConversionID))
	if err != nil {
		return err
	}

	// conversion
	if cfg.Conversion.Enabled {
		value, err := tm.UpsertVariable(ctx, gtmFolder, variable.NewEventData("value"))
		if err != nil {
			return err
		}

		currency, err := tm.UpsertVariable(ctx, gtmFolder, variable.NewEventData("currency"))
		if err != nil {
			return err
		}

		{ // create tags
			eventParameters, err := utils.LoadEventParams(ctx, cfg.Conversion.ServerContainer.Config)
			if err != nil {
				return err
			}

			for event := range eventParameters {
				var eventTriggerOpts []trigger.GoogleAdsEventOption
				if cfg.GoogleConsent.Enabled {
					if err := googleconsent.ServerEnsure(ctx, tm); err != nil {
						return err
					}
					consentVariable, err := tm.LookupVariable(ctx, googleconsentvariable.GoogleConsentModeName(cfg.GoogleConsent.Mode))
					if err != nil {
						return err
					}
					eventTriggerOpts = append(eventTriggerOpts, trigger.GoogleAdsEventWithConsentMode(consentVariable))
				}

				eventTrigger, err := tm.UpsertTrigger(ctx, folder, trigger.NewGoogleAdsEvent(event, eventTriggerOpts...))
				if err != nil {
					return errors.Wrap(err, "failed to upsert event trigger: "+event)
				}

				for _, setting := range cfg.Conversion.ServerContainer.Setting(event) {
					if _, err := tm.UpsertTag(ctx, folder, servertagx.NewGoogleAdsConversionTracking(event, value, currency, conversionID, setting, eventTrigger)); err != nil {
						return err
					}
				}
			}
		}

		// remarketing
		if cfg.Remarketing.Enabled {
			var eventTriggerOpts []trigger.GoogleAdsRemarketingEventOption
			if cfg.GoogleConsent.Enabled {
				if err := googleconsent.ServerEnsure(ctx, tm); err != nil {
					return err
				}
				consentVariable, err := tm.LookupVariable(ctx, googleconsentvariable.GoogleConsentModeName(cfg.GoogleConsent.Mode))
				if err != nil {
					return err
				}
				eventTriggerOpts = append(eventTriggerOpts, trigger.GoogleAdsRemarketingEventWithConsentMode(consentVariable))
			}

			eventTrigger, err := tm.UpsertTrigger(ctx, folder, trigger.NewGoogleAdsRemarketingEvent(NameGoogleAdsRemarketingTrigger, eventTriggerOpts...))
			if err != nil {
				return errors.Wrap(err, "failed to upsert event trigger: "+NameGoogleAdsRemarketingTrigger)
			}

			if _, err := tm.UpsertTag(ctx, folder, servertagx.NewGoogleAdsRemarketing(NameGoogleAdsRemarketingTag, conversionID, cfg.Remarketing, eventTrigger)); err != nil {
				return err
			}
		}
	}

	return nil
}
