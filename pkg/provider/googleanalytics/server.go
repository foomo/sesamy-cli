package googleanalytics

import (
	"github.com/foomo/sesamy-cli/pkg/config"
	googleanalyticsclient "github.com/foomo/sesamy-cli/pkg/provider/googleanalytics/server/client"
	containertag "github.com/foomo/sesamy-cli/pkg/provider/googleanalytics/server/tag"
	googleanalyticstemplate "github.com/foomo/sesamy-cli/pkg/provider/googleanalytics/server/template"
	"github.com/foomo/sesamy-cli/pkg/provider/googleanalytics/server/trigger"
	"github.com/foomo/sesamy-cli/pkg/provider/googleconsent"
	googleconsentvariable "github.com/foomo/sesamy-cli/pkg/provider/googleconsent/server/variable"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	serverclient "github.com/foomo/sesamy-cli/pkg/tagmanager/server/client"
	servertemplate "github.com/foomo/sesamy-cli/pkg/tagmanager/server/template"
	servertransformation "github.com/foomo/sesamy-cli/pkg/tagmanager/server/transformation"
	servertrigger "github.com/foomo/sesamy-cli/pkg/tagmanager/server/trigger"
	servervariable "github.com/foomo/sesamy-cli/pkg/tagmanager/server/variable"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"github.com/pkg/errors"
)

func Server(tm *tagmanager.TagManager, cfg config.GoogleAnalytics, redactVisitorIP bool) error {
	{ // create folder
		if folder, err := tm.UpsertFolder("Sesamy - " + Name); err != nil {
			return err
		} else {
			tm.SetFolderName(folder.Name)
		}
	}

	{ // create clients
		{
			client, err := tm.UpsertClient(serverclient.NewGoogleAnalyticsGA4(NameGoogleAnalyticsGA4Client))
			if err != nil {
				return err
			}
			if _, err = tm.UpsertTrigger(servertrigger.NewClient(NameGoogleAnalyticsGA4ClientTrigger, client)); err != nil {
				return err
			}
		}

		{
			client, err := tm.UpsertClient(serverclient.NewMeasurementProtocolGA4(NameMeasurementProtocolGA4Client))
			if err != nil {
				return err
			}
			if _, err = tm.UpsertTrigger(servertrigger.NewClient(NameMeasurementProtocolGA4ClientTrigger, client)); err != nil {
				return err
			}

			userDataTemplate, err := tm.UpsertCustomTemplate(servertemplate.NewJSONRequestValue(NameJSONRequestValueTemplate))
			if err != nil {
				return err
			}

			userDataVariable, err := tm.UpsertVariable(servervariable.NewMPv2Data("user_data", userDataTemplate))
			if err != nil {
				return err
			}

			_, err = tm.UpsertTransformation(servertransformation.NewMPv2UserData(NameMPv2UserDataTransformation, userDataVariable, client))
			if err != nil {
				return err
			}
		}

		if cfg.GoogleGTag.Enabled {
			template, err := tm.UpsertCustomTemplate(googleanalyticstemplate.NewGoogleGTagClient(NameGoogleGTagClientTemplate))
			if err != nil {
				return err
			}

			_, err = tm.UpsertClient(googleanalyticsclient.NewGoogleGTag(NameGoogleGTagClient, cfg.GoogleGTag, template))
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

			eventTrigger, err := tm.UpsertTrigger(trigger.NewGoogleAnalyticsEvent(event, eventTriggerOpts...))
			if err != nil {
				return errors.Wrap(err, "failed to upsert event trigger: "+event)
			}

			if _, err := tm.UpsertTag(containertag.NewGoogleAnalytics(event, redactVisitorIP, eventTrigger)); err != nil {
				return errors.Wrap(err, "failed to upsert google analytics ga4 tag: "+event)
			}
		}
	}

	return nil
}
