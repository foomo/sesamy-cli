package googleanalytics

import (
	"github.com/foomo/sesamy-cli/pkg/config"
	client2 "github.com/foomo/sesamy-cli/pkg/provider/googleanalytics/server/client"
	containertag "github.com/foomo/sesamy-cli/pkg/provider/googleanalytics/server/tag"
	template2 "github.com/foomo/sesamy-cli/pkg/provider/googleanalytics/server/template"
	"github.com/foomo/sesamy-cli/pkg/provider/googletag"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	commontrigger "github.com/foomo/sesamy-cli/pkg/tagmanager/common/trigger"
	serverclient "github.com/foomo/sesamy-cli/pkg/tagmanager/server/client"
	servertrigger "github.com/foomo/sesamy-cli/pkg/tagmanager/server/trigger"
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
		}

		if cfg.GoogleGTag.Enabled {
			template, err := tm.UpsertCustomTemplate(template2.NewGoogleGTagClient(NameGoogleGTagClientTemplate))
			if err != nil {
				return err
			}

			_, err = tm.UpsertClient(client2.NewGoogleGTag(NameGoogleGTagClient, cfg.GoogleGTag, template))
			if err != nil {
				return err
			}
		}
	}

	{ // create tags
		eventParameters, err := googletag.CreateServerEventTriggers(tm, cfg.ServerContainer)
		if err != nil {
			return err
		}

		for event := range eventParameters {
			eventTrigger, err := tm.LookupTrigger(commontrigger.EventName(event))
			if err != nil {
				return errors.Wrap(err, "failed to lookup event trigger: "+event)
			}

			if _, err := tm.UpsertTag(containertag.NewGoogleAnalyticsGA4(event, redactVisitorIP, eventTrigger)); err != nil {
				return errors.Wrap(err, "failed to upsert google analytics ga4 tag: "+event)
			}
		}
	}

	return nil
}
