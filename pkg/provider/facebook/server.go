package facebook

import (
	"context"
	"log/slog"

	"github.com/foomo/sesamy-cli/pkg/config"
	servertagx "github.com/foomo/sesamy-cli/pkg/provider/facebook/server/tag"
	"github.com/foomo/sesamy-cli/pkg/provider/facebook/server/trigger"
	"github.com/foomo/sesamy-cli/pkg/provider/googleconsent"
	googleconsentvariable "github.com/foomo/sesamy-cli/pkg/provider/googleconsent/server/variable"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	commonvariable "github.com/foomo/sesamy-cli/pkg/tagmanager/common/variable"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"github.com/pkg/errors"
)

func Server(ctx context.Context, l *slog.Logger, tm *tagmanager.TagManager, cfg config.Facebook) error {
	folder, err := tm.UpsertFolder(ctx, "Sesamy - "+Name)
	if err != nil {
		return err
	}

	pixelID, err := tm.UpsertVariable(ctx, folder, commonvariable.NewConstant(NamePixelIDConstant, cfg.PixelID))
	if err != nil {
		return err
	}

	apiAccessToken, err := tm.UpsertVariable(ctx, folder, commonvariable.NewConstant(NameAPIAcessTokenConstant, cfg.APIAccessToken))
	if err != nil {
		return err
	}

	testEventToken, err := tm.UpsertVariable(ctx, folder, commonvariable.NewConstant(NameTestEventTokenConstant, cfg.TestEventToken))
	if err != nil {
		return err
	}

	template, err := tm.LookupTemplate(ctx, NameConversionsAPITagTemplate)
	if errors.Is(err, tagmanager.ErrNotFound) {
		l.Warn("Please install the 'Conversions API Tag' by 'facebookincubator' template manually first")
		return err
	} else if err != nil {
		return err
	}

	{ // create tags
		eventParameters, err := utils.LoadEventParams(ctx, cfg.ServerContainer.Config)
		if err != nil {
			return err
		}

		for event := range eventParameters {
			var eventTriggerOpts []trigger.FacebookEventOption
			if cfg.GoogleConsent.Enabled {
				if err := googleconsent.ServerEnsure(ctx, tm); err != nil {
					return err
				}
				consentVariable, err := tm.LookupVariable(ctx, googleconsentvariable.GoogleConsentModeName(cfg.GoogleConsent.Mode))
				if err != nil {
					return err
				}
				eventTriggerOpts = append(eventTriggerOpts, trigger.FacebookEventWithConsentMode(consentVariable))
			}

			eventTrigger, err := tm.UpsertTrigger(ctx, folder, trigger.NewFacebookEvent(event, eventTriggerOpts...))
			if err != nil {
				return errors.Wrap(err, "failed to upsert event trigger: "+event)
			}

			if _, err := tm.UpsertTag(ctx, folder, servertagx.NewConversionsAPITag(event, pixelID, apiAccessToken, testEventToken, cfg.ServerContainer.Setting(event), template, eventTrigger)); err != nil {
				return err
			}
		}
	}

	return nil
}
