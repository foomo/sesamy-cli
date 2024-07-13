package conversionlinker

import (
	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	servertag "github.com/foomo/sesamy-cli/pkg/tagmanager/server/tag"
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

	{
		if _, err := tm.UpsertTag(servertag.NewConversionLinker(Name, servertrigger.AllPages)); err != nil {
			return err
		}
	}

	return nil
}
