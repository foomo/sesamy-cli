package hotjar

import (
	"context"

	"github.com/foomo/sesamy-cli/pkg/config"
	client "github.com/foomo/sesamy-cli/pkg/provider/hotjar/web/tag"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	commonvariable "github.com/foomo/sesamy-cli/pkg/tagmanager/common/variable"
)

func Web(ctx context.Context, tm *tagmanager.TagManager, cfg config.Hotjar) error {
	folder, err := tm.UpsertFolder("Sesamy - " + Name)
	if err != nil {
		return err
	}

	{ // setup hotjar
		siteID, err := tm.UpsertVariable(folder, commonvariable.NewConstant(NameSiteID, cfg.SiteID))
		if err != nil {
			return err
		}

		if _, err = tm.UpsertTag(folder, client.NewHotjar(NameHotjarTag, siteID)); err != nil {
			return err
		}
	}

	return nil
}
