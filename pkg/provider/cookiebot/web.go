package cookiebot

import (
	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/provider/cookiebot/web/tag"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	"github.com/pkg/errors"
)

func Web(tm *tagmanager.TagManager, cfg config.Cookiebot) error {
	{ // create folder
		if folder, err := tm.UpsertFolder("Sesamy - " + Name); err != nil {
			return err
		} else {
			tm.SetFolderName(folder.Name)
		}
	}

	{ // create event tags
		temmplate, err := tm.LookupTemplate(cfg.TemplateName)
		if err != nil {
			return errors.Wrapf(err, "Failed to lookup `%s`, please install the `%s` gallery tag template first (%s)", cfg.TemplateName, "Cookiebot CMP", "https://tagmanager.google.com/gallery/#/owners/cybotcorp/templates/gtm-templates-cookiebot-cmp")
		}

		if _, err := tm.UpsertTag(tag.NewCookiebotInitialization(NameCookiebotTag, cfg, temmplate)); err != nil {
			return err
		}
	}

	return nil
}
