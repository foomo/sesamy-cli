package googleanalytics

import (
	"context"

	"github.com/foomo/sesamy-cli/pkg/config"
	googleanalyticsclient "github.com/foomo/sesamy-cli/pkg/provider/googleanalytics/server/client"
	containertag "github.com/foomo/sesamy-cli/pkg/provider/googleanalytics/server/tag"
	googleanalyticstemplate "github.com/foomo/sesamy-cli/pkg/provider/googleanalytics/server/template"
	"github.com/foomo/sesamy-cli/pkg/provider/googleanalytics/server/trigger"
	"github.com/foomo/sesamy-cli/pkg/provider/googleconsent"
	googleconsentvariable "github.com/foomo/sesamy-cli/pkg/provider/googleconsent/server/variable"
	"github.com/foomo/sesamy-cli/pkg/provider/googletag"
	"github.com/foomo/sesamy-cli/pkg/provider/googletagmanager"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	servertemplate "github.com/foomo/sesamy-cli/pkg/tagmanager/server/template"
	servertransformation "github.com/foomo/sesamy-cli/pkg/tagmanager/server/transformation"
	servertrigger "github.com/foomo/sesamy-cli/pkg/tagmanager/server/trigger"
	servervariable "github.com/foomo/sesamy-cli/pkg/tagmanager/server/variable"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"github.com/pkg/errors"
	api "google.golang.org/api/tagmanager/v2"
)

func Server(ctx context.Context, tm *tagmanager.TagManager, cfg config.GoogleAnalytics, redactVisitorIP, enableGeoResolution bool) error {
	folder, err := tm.UpsertFolder(ctx, "Sesamy - "+Name)
	if err != nil {
		return err
	}

	{ // create clients
		{
			measurementID, err := tm.LookupVariable(ctx, googletag.NameGoogleTagMeasurementID)
			if err != nil {
				return err
			}

			visitorRegion, err := tm.LookupVariable(ctx, googletagmanager.NameGoogleTagManagerVisitorRegion)
			if err != nil {
				return err
			}

			client, err := tm.UpsertClient(ctx, folder, googleanalyticsclient.NewGoogleAnalyticsGA4(NameGoogleAnalyticsGA4Client, enableGeoResolution, visitorRegion, measurementID))
			if err != nil {
				return err
			}

			if _, err = tm.UpsertTrigger(ctx, folder, servertrigger.NewClient(NameGoogleAnalyticsGA4ClientTrigger, client)); err != nil {
				return err
			}
		}

		{
			client, err := tm.UpsertClient(ctx, folder, googleanalyticsclient.NewMeasurementProtocolGA4(NameMeasurementProtocolGA4Client))
			if err != nil {
				return err
			}
			if _, err = tm.UpsertTrigger(ctx, folder, servertrigger.NewClient(NameMeasurementProtocolGA4ClientTrigger, client)); err != nil {
				return err
			}

			if cfg.EnableMPv2UserDataTransformation {
				userDataTemplate, err := tm.UpsertCustomTemplate(ctx, servertemplate.NewJSONRequestValue(NameJSONRequestValueTemplate))
				if err != nil {
					return err
				}

				userDataVariable, err := tm.UpsertVariable(ctx, folder, servervariable.NewMPv2Data("user_data", userDataTemplate))
				if err != nil {
					return err
				}

				debugModeVariable, err := tm.UpsertVariable(ctx, folder, servervariable.NewMPv2Data("debug_mode", userDataTemplate))
				if err != nil {
					return err
				}

				_, err = tm.UpsertTransformation(ctx, folder, servertransformation.NewMPv2UserData(NameMPv2UserDataTransformation, map[string]*api.Variable{
					"user_data":  userDataVariable,
					"debug_mode": debugModeVariable,
				}, client))
				if err != nil {
					return err
				}
			}
		}

		if cfg.GoogleGTagJSOverride.Enabled {
			template, err := tm.UpsertCustomTemplate(ctx, googleanalyticstemplate.NewGoogleGTagClient(NameGoogleGTagClientTemplate))
			if err != nil {
				return err
			}

			_, err = tm.UpsertClient(ctx, folder, googleanalyticsclient.NewGoogleGTag(NameGoogleGTagClient, cfg.GoogleGTagJSOverride, template))
			if err != nil {
				return err
			}
		}
	}

	{ // create tags
		eventParameters, err := utils.LoadEventParams(ctx, cfg.ServerContainer)
		if err != nil {
			return err
		}

		for event := range eventParameters {
			var eventTriggerOpts []trigger.GoogleAnalyticsEventOption
			if cfg.GoogleConsent.Enabled {
				if err := googleconsent.ServerEnsure(ctx, tm); err != nil {
					return err
				}
				consentVariable, err := tm.LookupVariable(ctx, googleconsentvariable.GoogleConsentModeName(cfg.GoogleConsent.Mode))
				if err != nil {
					return err
				}
				eventTriggerOpts = append(eventTriggerOpts, trigger.GoogleAnalyticsEventWithConsentMode(consentVariable))
			}

			eventTrigger, err := tm.UpsertTrigger(ctx, folder, trigger.NewGoogleAnalyticsEvent(event, eventTriggerOpts...))
			if err != nil {
				return errors.Wrap(err, "failed to upsert event trigger: "+event)
			}

			if _, err := tm.UpsertTag(ctx, folder, containertag.NewGoogleAnalytics(event, redactVisitorIP, eventTrigger)); err != nil {
				return errors.Wrap(err, "failed to upsert google analytics ga4 tag: "+event)
			}
		}
	}

	return nil
}
