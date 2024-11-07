package hotjar

import (
	"github.com/foomo/sesamy-cli/pkg/config"
	client "github.com/foomo/sesamy-cli/pkg/provider/hotjar/web/tag"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	commonvariable "github.com/foomo/sesamy-cli/pkg/tagmanager/common/variable"
)

func Web(tm *tagmanager.TagManager, cfg config.Hotjar) error {
	{ // create folder
		if folder, err := tm.UpsertFolder("Sesamy - " + Name); err != nil {
			return err
		} else {
			tm.SetFolderName(folder.Name)
		}
	}

	{ // setup hotjar
		siteID, err := tm.UpsertVariable(commonvariable.NewConstant(NameSiteID, cfg.SiteID))
		if err != nil {
			return err
		}

		if _, err = tm.UpsertTag(client.NewHotjar(NameHotjarTag, siteID)); err != nil {
			return err
		}
	}

	return nil
}
