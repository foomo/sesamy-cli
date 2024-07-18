package googletagmanager

import (
	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	containerclient "github.com/foomo/sesamy-cli/pkg/tagmanager/server/client"
)

func Server(tm *tagmanager.TagManager, cfg config.GoogleTagManager) error {
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
		if _, err := tm.UpsertClient(containerclient.NewGoogleTagManagerWebContainer(NameGoogleTagManagerWebContainerClient, cfg.WebContainer.TagID)); err != nil {
			return err
		}
	}

	return nil
}
