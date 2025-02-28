package googletag

import (
	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	commonvariable "github.com/foomo/sesamy-cli/pkg/tagmanager/common/variable"
)

func Server(tm *tagmanager.TagManager, cfg config.GoogleTag) error {
	folder, err := tm.UpsertFolder("Sesamy - " + Name)
	if err != nil {
		return err
	}

	{ // create constants
		if _, err := tm.UpsertVariable(folder, commonvariable.NewConstant(NameGoogleTagMeasurementID, cfg.TagID)); err != nil {
			return err
		}
	}

	return nil
}
