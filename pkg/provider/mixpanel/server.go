package mixpanel

import (
	"context"
	"log/slog"

	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/provider/googleconsent"
	googleconsentvariable "github.com/foomo/sesamy-cli/pkg/provider/googleconsent/server/variable"
	"github.com/foomo/sesamy-cli/pkg/provider/googletag"
	servertagx "github.com/foomo/sesamy-cli/pkg/provider/mixpanel/server/tag"
	"github.com/foomo/sesamy-cli/pkg/provider/mixpanel/server/trigger"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	commonvariable "github.com/foomo/sesamy-cli/pkg/tagmanager/common/variable"
	"github.com/foomo/sesamy-cli/pkg/tagmanager/server/variable"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"github.com/pkg/errors"
	tagmanager2 "google.golang.org/api/tagmanager/v2"
)

func Server(ctx context.Context, l *slog.Logger, tm *tagmanager.TagManager, cfg config.Mixpanel) error {
	folder, err := Folder(ctx, tm)
	if err != nil {
		return err
	}

	gtagFolder, err := googletag.Folder(ctx, tm)
	if err != nil {
		return err
	}

	template, err := tm.LookupTemplate(ctx, NameTagTemplate)
	if err != nil {
		if errors.Is(err, tagmanager.ErrNotFound) {
			l.Warn("Please install the 'Mixpanel' by stape-io Tag Template manually first")
		}
		return err
	}

	projectToken, err := tm.UpsertVariable(ctx, folder, commonvariable.NewConstant(NamePrjectTokenConstant, cfg.ProjectToken))
	if err != nil {
		return err
	}

	{ // create track tags
		eventParameters, err := utils.LoadEventParams(ctx, cfg.ServerContainer.Track)
		if err != nil {
			return err
		}

		for event, params := range eventParameters {
			var eventTriggerOpts []trigger.EventOption
			if cfg.GoogleConsent.Enabled {
				if err := googleconsent.ServerEnsure(ctx, tm); err != nil {
					return err
				}
				consentVariable, err := tm.LookupVariable(ctx, googleconsentvariable.GoogleConsentModeName(cfg.GoogleConsent.Mode))
				if err != nil {
					return err
				}
				eventTriggerOpts = append(eventTriggerOpts, trigger.EventWithConsentMode(consentVariable))
			}

			eventParams := map[string]*tagmanager2.Variable{}
			for k := range params {
				if value, err := tm.UpsertVariable(ctx, gtagFolder, variable.NewEventData(k)); err != nil {
					return err
				} else {
					eventParams[k] = value
				}
			}

			eventTrigger, err := tm.UpsertTrigger(ctx, folder, trigger.NewEvent(event, eventTriggerOpts...))
			if err != nil {
				return errors.Wrap(err, "failed to upsert event trigger: "+event)
			}

			if _, err := tm.UpsertTag(ctx, folder, servertagx.NewEvent(servertagx.TypeTrack, event, projectToken, template, eventParams, eventTrigger)); err != nil {
				return err
			}
		}
	}

	return nil
}
