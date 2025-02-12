package googletagmanager

import (
	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/provider/googletagmanager/server/client"
	"github.com/foomo/sesamy-cli/pkg/provider/googletagmanager/server/variable"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
)

func Server(tm *tagmanager.TagManager, cfg config.GoogleTagManager, enableGeoResolution bool) error {
	{ // create folder
		if folder, err := tm.UpsertFolder("Sesamy - " + Name); err != nil {
			return err
		} else {
			tm.SetFolderName(folder.Name)
		}
	}

	{ // enable build in variables
		if _, err := tm.EnableBuiltInVariable("clientName"); err != nil {
			return err
		}
	}

	{ // create client
		visitorRegion, err := tm.UpsertVariable(variable.NewVisitorRegion(NameGoogleTagManagerVisitorRegion))
		if err != nil {
			return err
		}

		if _, err := tm.UpsertClient(client.NewGoogleTagManagerWebContainer(NameGoogleTagManagerWebContainerClient, cfg.WebContainer.TagID, enableGeoResolution, visitorRegion)); err != nil {
			return err
		}
	}

	return nil
}
