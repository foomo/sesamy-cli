package googletagmanager

import (
	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/provider/googletagmanager/server/client"
	"github.com/foomo/sesamy-cli/pkg/provider/googletagmanager/server/variable"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	commonvariable "github.com/foomo/sesamy-cli/pkg/tagmanager/common/variable"
	servervariable "github.com/foomo/sesamy-cli/pkg/tagmanager/server/variable"
)

func Server(tm *tagmanager.TagManager, cfg config.GoogleTagManager, enableGeoResolution bool) error {
	folder, err := tm.UpsertFolder("Sesamy - " + Name)
	if err != nil {
		return err
	}

	{ // enable build in variables
		if _, err := tm.EnableBuiltInVariable("clientName"); err != nil {
			return err
		}
	}

	{ // create client
		visitorRegion, err := tm.UpsertVariable(folder, variable.NewVisitorRegion(NameGoogleTagManagerVisitorRegion))
		if err != nil {
			return err
		}

		if _, err := tm.UpsertClient(folder, client.NewGoogleTagManagerWebContainer(NameGoogleTagManagerWebContainerClient, cfg.WebContainer.TagID, enableGeoResolution, visitorRegion)); err != nil {
			return err
		}
	}

	{ // create variables
		for _, value := range cfg.ServerContaienrVariables.EventData {
			if _, err := tm.UpsertVariable(folder, servervariable.NewEventData(value)); err != nil {
				return err
			}
		}
		for key, value := range cfg.ServerContaienrVariables.LookupTables {
			if _, err := tm.UpsertVariable(folder, commonvariable.NewLookupTable(key, value)); err != nil {
				return err
			}
		}
	}

	return nil
}
