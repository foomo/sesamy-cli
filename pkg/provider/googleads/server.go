package googleads

import (
	"log/slog"

	"github.com/foomo/sesamy-cli/pkg/config"
	servertagx "github.com/foomo/sesamy-cli/pkg/provider/googleads/server/tag"
	"github.com/foomo/sesamy-cli/pkg/provider/googleads/server/trigger"
	"github.com/foomo/sesamy-cli/pkg/provider/googleconsent"
	googleconsentvariable "github.com/foomo/sesamy-cli/pkg/provider/googleconsent/server/variable"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	commonvariable "github.com/foomo/sesamy-cli/pkg/tagmanager/common/variable"
	"github.com/foomo/sesamy-cli/pkg/tagmanager/server/variable"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"github.com/pkg/errors"
)

func Server(l *slog.Logger, tm *tagmanager.TagManager, cfg config.GoogleAds) error {
	{ // create folder
		if folder, err := tm.UpsertFolder("Sesamy - " + Name); err != nil {
			return err
		} else {
			tm.SetFolderName(folder.Name)
		}
	}

	conversionID, err := tm.UpsertVariable(commonvariable.NewConstant(NameConversionIDConstant, cfg.ConversionID))
	if err != nil {
		return err
	}

	// conversion
	if cfg.Conversion.Enabled {
		conversionLabel, err := tm.UpsertVariable(commonvariable.NewConstant(NameConversionLabelConstant, cfg.Conversion.ConversionLabel))
		if err != nil {
			return err
		}

		value, err := tm.UpsertVariable(variable.NewEventData("value"))
		if err != nil {
			return err
		}

		currency, err := tm.UpsertVariable(variable.NewEventData("currency"))
		if err != nil {
			return err
		}

		{ // create tags
			eventParameters, err := utils.LoadEventParams(cfg.Conversion.ServerContainer)
			if err != nil {
				return err
			}

			for event := range eventParameters {
				var eventTriggerOpts []trigger.GoogleAdsEventOption
				if cfg.GoogleConsent.Enabled {
					if err := googleconsent.ServerEnsure(tm); err != nil {
						return err
					}
					consentVariable, err := tm.LookupVariable(googleconsentvariable.GoogleConsentModeName(cfg.GoogleConsent.Mode))
					if err != nil {
						return err
					}
					eventTriggerOpts = append(eventTriggerOpts, trigger.GoogleAdsEventWithConsentMode(consentVariable))
				}

				eventTrigger, err := tm.UpsertTrigger(trigger.NewGoogleAdsEvent(event, eventTriggerOpts...))
				if err != nil {
					return errors.Wrap(err, "failed to upsert event trigger: "+event)
				}

				if _, err := tm.UpsertTag(servertagx.NewGoogleAdsConversionTracking(event, value, currency, conversionID, conversionLabel, eventTrigger)); err != nil {
					return err
				}
			}
		}

		// remarketing
		if cfg.Remarketing.Enabled {
			var eventTriggerOpts []trigger.GoogleAdsRemarketingEventOption
			if cfg.GoogleConsent.Enabled {
				if err := googleconsent.ServerEnsure(tm); err != nil {
					return err
				}
				consentVariable, err := tm.LookupVariable(googleconsentvariable.GoogleConsentModeName(cfg.GoogleConsent.Mode))
				if err != nil {
					return err
				}
				eventTriggerOpts = append(eventTriggerOpts, trigger.GoogleAdsRemarketingEventWithConsentMode(consentVariable))
			}

			eventTrigger, err := tm.UpsertTrigger(trigger.NewGoogleAdsRemarketingEvent(NameGoogleAdsRemarketingTrigger, eventTriggerOpts...))
			if err != nil {
				return errors.Wrap(err, "failed to upsert event trigger: "+NameGoogleAdsRemarketingTrigger)
			}

			if _, err := tm.UpsertTag(servertagx.NewGoogleAdsRemarketing(NameGoogleAdsRemarketingTag, conversionID, cfg.Remarketing, eventTrigger)); err != nil {
				return err
			}
		}
	}

	return nil
}
