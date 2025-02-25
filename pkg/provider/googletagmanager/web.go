package googletagmanager

import (
	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	commonvariable "github.com/foomo/sesamy-cli/pkg/tagmanager/common/variable"
	"github.com/foomo/sesamy-cli/pkg/tagmanager/web/variable"
)

func Web(tm *tagmanager.TagManager, cfg config.GoogleTagManager) error {
	{ // create folder
		if folder, err := tm.UpsertFolder("Sesamy - " + Name); err != nil {
			return err
		} else {
			tm.SetFolderName(folder.Name)
		}
	}

	{ // create variables
		for _, value := range cfg.WebContaienrVariables.DataLayer {
			if _, err := tm.UpsertVariable(variable.NewDataLayer(value)); err != nil {
				return err
			}
		}
		for key, value := range cfg.WebContaienrVariables.LookupTables {
			if _, err := tm.UpsertVariable(commonvariable.NewLookupTable(key, value)); err != nil {
				return err
			}
		}
	}

	return nil
}
