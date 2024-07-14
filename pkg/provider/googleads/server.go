package googleads

import (
	"log/slog"

	"github.com/foomo/sesamy-cli/pkg/config"
	servertagx "github.com/foomo/sesamy-cli/pkg/provider/googleads/server/tag"
	"github.com/foomo/sesamy-cli/pkg/provider/googletag"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	commontrigger "github.com/foomo/sesamy-cli/pkg/tagmanager/common/trigger"
	commonvariable "github.com/foomo/sesamy-cli/pkg/tagmanager/common/variable"
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

	{ // conversion
		conversionID, err := tm.UpsertVariable(commonvariable.NewConstant(NameConversionIDConstant, cfg.Conversion.ConversionID))
		if err != nil {
			return err
		}

		conversionLabel, err := tm.UpsertVariable(commonvariable.NewConstant(NameConversionLabelConstant, cfg.Conversion.ConversionLabel))
		if err != nil {
			return err
		}

		{ // create tags
			eventParameters, err := googletag.CreateServerEventTriggers(tm, cfg.Conversion.ServerContainer)
			if err != nil {
				return err
			}

			for event := range eventParameters {
				eventTrigger, err := tm.LookupTrigger(commontrigger.EventName(event))
				if err != nil {
					return errors.Wrap(err, "failed to lookup event trigger: "+event)
				}

				if _, err := tm.UpsertTag(servertagx.NewGoogleAdsConversionTracking(event, conversionID, conversionLabel, eventTrigger)); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
