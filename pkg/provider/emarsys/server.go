package emarsys

import (
	"log/slog"

	"github.com/foomo/sesamy-cli/pkg/config"
	serverclientx "github.com/foomo/sesamy-cli/pkg/provider/emarsys/server/client"
	servertagx "github.com/foomo/sesamy-cli/pkg/provider/emarsys/server/tag"
	"github.com/foomo/sesamy-cli/pkg/provider/emarsys/server/template"
	"github.com/foomo/sesamy-cli/pkg/provider/googletag"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	commontrigger "github.com/foomo/sesamy-cli/pkg/tagmanager/common/trigger"
	commonvariable "github.com/foomo/sesamy-cli/pkg/tagmanager/common/variable"
	"github.com/pkg/errors"
)

func Server(l *slog.Logger, tm *tagmanager.TagManager, cfg config.Emarsys) error {
	{ // create folder
		if folder, err := tm.UpsertFolder("Sesamy - " + Name); err != nil {
			return err
		} else {
			tm.SetFolderName(folder.Name)
		}
	}

	{ // conversion
		merchantID, err := tm.UpsertVariable(commonvariable.NewConstant(NameMerchantIDConstant, cfg.MerchantID))
		if err != nil {
			return err
		}

		tagTemplate, err := tm.UpsertCustomTemplate(template.NewEmarsysWebExtendTag(NameServerEmarsysWebExtendTagTemplate))
		if err != nil {
			return err
		}

		clientTemplate, err := tm.UpsertCustomTemplate(template.NewEmarsysInitializationClient(NameServerEmarsysInitalizationClientTemplate))
		if err != nil {
			return err
		}

		_, err = tm.UpsertClient(serverclientx.NewEmarsys(NameServerEmarsysClient, cfg, clientTemplate))
		if err != nil {
			return err
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

				if _, err := tm.UpsertTag(servertagx.NewEmarsys(event, cfg.NewPageViewEvent == eventTrigger.Name, merchantID, tagTemplate, eventTrigger)); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
