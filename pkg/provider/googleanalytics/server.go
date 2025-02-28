package googleanalytics

import (
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

func Server(tm *tagmanager.TagManager, cfg config.GoogleAnalytics, redactVisitorIP, enableGeoResolution bool) error {
	folder, err := tm.UpsertFolder("Sesamy - " + Name)
	if err != nil {
		return err
	}

	{ // create clients
		{
			measurementID, err := tm.LookupVariable(googletag.NameGoogleTagMeasurementID)
			if err != nil {
				return err
			}

			visitorRegion, err := tm.LookupVariable(googletagmanager.NameGoogleTagManagerVisitorRegion)
			if err != nil {
				return err
			}

			client, err := tm.UpsertClient(folder, googleanalyticsclient.NewGoogleAnalyticsGA4(NameGoogleAnalyticsGA4Client, enableGeoResolution, visitorRegion, measurementID))
			if err != nil {
				return err
			}

			if _, err = tm.UpsertTrigger(folder, servertrigger.NewClient(NameGoogleAnalyticsGA4ClientTrigger, client)); err != nil {
				return err
			}
		}

		{
			client, err := tm.UpsertClient(folder, googleanalyticsclient.NewMeasurementProtocolGA4(NameMeasurementProtocolGA4Client))
			if err != nil {
				return err
			}
			if _, err = tm.UpsertTrigger(folder, servertrigger.NewClient(NameMeasurementProtocolGA4ClientTrigger, client)); err != nil {
				return err
			}

			userDataTemplate, err := tm.UpsertCustomTemplate(servertemplate.NewJSONRequestValue(NameJSONRequestValueTemplate))
			if err != nil {
				return err
			}

			userDataVariable, err := tm.UpsertVariable(folder, servervariable.NewMPv2Data("user_data", userDataTemplate))
			if err != nil {
				return err
			}

			debugModeVariable, err := tm.UpsertVariable(folder, servervariable.NewMPv2Data("debug_mode", userDataTemplate))
			if err != nil {
				return err
			}

			_, err = tm.UpsertTransformation(folder, servertransformation.NewMPv2UserData(NameMPv2UserDataTransformation, map[string]*api.Variable{
				"user_data":  userDataVariable,
				"debug_mode": debugModeVariable,
			}, client))
			if err != nil {
				return err
			}
		}

		if cfg.GoogleGTagJSOverride.Enabled {
			template, err := tm.UpsertCustomTemplate(googleanalyticstemplate.NewGoogleGTagClient(NameGoogleGTagClientTemplate))
			if err != nil {
				return err
			}

			_, err = tm.UpsertClient(folder, googleanalyticsclient.NewGoogleGTag(NameGoogleGTagClient, cfg.GoogleGTagJSOverride, template))
			if err != nil {
				return err
			}
		}
	}

	{ // create tags
		eventParameters, err := utils.LoadEventParams(cfg.ServerContainer)
		if err != nil {
			return err
		}

		for event := range eventParameters {
			var eventTriggerOpts []trigger.GoogleAnalyticsEventOption
			if cfg.GoogleConsent.Enabled {
				if err := googleconsent.ServerEnsure(tm); err != nil {
					return err
				}
				consentVariable, err := tm.LookupVariable(googleconsentvariable.GoogleConsentModeName(cfg.GoogleConsent.Mode))
				if err != nil {
					return err
				}
				eventTriggerOpts = append(eventTriggerOpts, trigger.GoogleAnalyticsEventWithConsentMode(consentVariable))
			}

			eventTrigger, err := tm.UpsertTrigger(folder, trigger.NewGoogleAnalyticsEvent(event, eventTriggerOpts...))
			if err != nil {
				return errors.Wrap(err, "failed to upsert event trigger: "+event)
			}

			if _, err := tm.UpsertTag(folder, containertag.NewGoogleAnalytics(event, redactVisitorIP, eventTrigger)); err != nil {
				return errors.Wrap(err, "failed to upsert google analytics ga4 tag: "+event)
			}
		}
	}

	return nil
}
