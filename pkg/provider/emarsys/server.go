package emarsys

import (
	"context"
	"log/slog"

	"github.com/foomo/sesamy-cli/pkg/config"
	serverclientx "github.com/foomo/sesamy-cli/pkg/provider/emarsys/server/client"
	servertagx "github.com/foomo/sesamy-cli/pkg/provider/emarsys/server/tag"
	"github.com/foomo/sesamy-cli/pkg/provider/emarsys/server/template"
	"github.com/foomo/sesamy-cli/pkg/provider/emarsys/server/trigger"
	"github.com/foomo/sesamy-cli/pkg/provider/googleconsent"
	googleconsentvariable "github.com/foomo/sesamy-cli/pkg/provider/googleconsent/server/variable"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	commonvariable "github.com/foomo/sesamy-cli/pkg/tagmanager/common/variable"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"github.com/pkg/errors"
)

func Server(ctx context.Context, l *slog.Logger, tm *tagmanager.TagManager, cfg config.Emarsys) error {
	folder, err := tm.UpsertFolder("Sesamy - " + Name)
	if err != nil {
		return err
	}

	{ // conversion
		merchantID, err := tm.UpsertVariable(folder, commonvariable.NewConstant(NameMerchantIDConstant, cfg.MerchantID))
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

		_, err = tm.UpsertClient(folder, serverclientx.NewEmarsys(NameServerEmarsysClient, cfg, clientTemplate))
		if err != nil {
			return err
		}

		{ // create tags
			eventParameters, err := utils.LoadEventParams(ctx, cfg.ServerContainer)
			if err != nil {
				return err
			}

			for event := range eventParameters {
				var eventTriggerOpts []trigger.EmarsysEventOption
				if cfg.GoogleConsent.Enabled {
					if err := googleconsent.ServerEnsure(tm); err != nil {
						return err
					}
					consentVariable, err := tm.LookupVariable(googleconsentvariable.GoogleConsentModeName(cfg.GoogleConsent.Mode))
					if err != nil {
						return err
					}
					eventTriggerOpts = append(eventTriggerOpts, trigger.EmarsysEventWithConsentMode(consentVariable))
				}

				eventTrigger, err := tm.UpsertTrigger(folder, trigger.NewEmarsysEvent(event, eventTriggerOpts...))
				if err != nil {
					return errors.Wrap(err, "failed to upsert event trigger: "+event)
				}

				if _, err := tm.UpsertTag(folder, servertagx.NewEmarsys(event, merchantID, cfg.TestMode, cfg.DebugMode, tagTemplate, eventTrigger)); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
