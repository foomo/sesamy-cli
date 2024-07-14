package facebook

import (
	"log/slog"

	"github.com/foomo/sesamy-cli/pkg/config"
	servertagx "github.com/foomo/sesamy-cli/pkg/provider/facebook/server/tag"
	"github.com/foomo/sesamy-cli/pkg/provider/googletag"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	commontrigger "github.com/foomo/sesamy-cli/pkg/tagmanager/common/trigger"
	commonvariable "github.com/foomo/sesamy-cli/pkg/tagmanager/common/variable"
	"github.com/pkg/errors"
)

func Server(l *slog.Logger, tm *tagmanager.TagManager, cfg config.Facebook) error {
	{ // create folder
		if folder, err := tm.UpsertFolder("Sesamy - " + Name); err != nil {
			return err
		} else {
			tm.SetFolderName(folder.Name)
		}
	}

	pixelID, err := tm.UpsertVariable(commonvariable.NewConstant(NamePixelIDConstant, cfg.PixelID))
	if err != nil {
		return err
	}

	apiAccessToken, err := tm.UpsertVariable(commonvariable.NewConstant(NameAPIAcessTokenConstant, cfg.APIAccessToken))
	if err != nil {
		return err
	}

	testEventToken, err := tm.UpsertVariable(commonvariable.NewConstant(NameTestEventTokenConstant, cfg.TestEventToken))
	if err != nil {
		return err
	}

	template, err := tm.LookupTemplate(NameConversionsAPITagTemplate)
	if err != nil {
		if errors.Is(err, tagmanager.ErrNotFound) {
			l.Warn("Please install the 'Conversion API' template manually first")
		}
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

			if _, err := tm.UpsertTag(servertagx.NewConversionsAPITag(event, pixelID, apiAccessToken, testEventToken, template, eventTrigger)); err != nil {
				return err
			}
		}
	}

	return nil
}
