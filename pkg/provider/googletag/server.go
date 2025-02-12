package googletag

import (
	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	commonvariable "github.com/foomo/sesamy-cli/pkg/tagmanager/common/variable"
)

func Server(tm *tagmanager.TagManager, cfg config.GoogleTag) error {
	{ // create folder
		if folder, err := tm.UpsertFolder("Sesamy - " + Name); err != nil {
			return err
		} else {
			tm.SetFolderName(folder.Name)
		}
	}

	{ // create constants
		if _, err := tm.UpsertVariable(commonvariable.NewConstant(NameGoogleTagMeasurementID, cfg.TagID)); err != nil {
			return err
		}
	}

	return nil
}
