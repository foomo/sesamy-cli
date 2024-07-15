package conversionlinker

import (
	"github.com/foomo/sesamy-cli/pkg/config"
	containertag "github.com/foomo/sesamy-cli/pkg/provider/conversionlinker/server/tag"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	servertrigger "github.com/foomo/sesamy-cli/pkg/tagmanager/server/trigger"
)

func Server(tm *tagmanager.TagManager, events config.ConversionLinker) error {
	{
		if folder, err := tm.UpsertFolder("Sesamy - " + Name); err != nil {
			return err
		} else {
			tm.SetFolderName(folder.Name)
		}
	}

	if _, err := tm.UpsertTag(containertag.NewConversionLinker(Name, servertrigger.AllPages)); err != nil {
		return err
	}

	return nil
}
