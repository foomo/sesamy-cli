package googletag

import (
	"github.com/foomo/gocontemplate/pkg/contemplate"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	commontrigger "github.com/foomo/sesamy-cli/pkg/tagmanager/common/trigger"
	"github.com/foomo/sesamy-cli/pkg/utils"
)

func Server(tm *tagmanager.TagManager) error {
	{ // create folder
		if folder, err := tm.UpsertFolder("Sesamy - " + Name); err != nil {
			return err
		} else {
			tm.SetFolderName(folder.Name)
		}
	}
	return nil
}

func CreateServerEventTriggers(tm *tagmanager.TagManager, cfg contemplate.Config) (map[string][]string, error) {
	previousFolderName := tm.FolderName()
	tm.SetFolderName("Sesamy - " + Name)
	defer tm.SetFolderName(previousFolderName)

	eventParameters, err := utils.LoadEventParams(cfg)
	if err != nil {
		return nil, err
	}

	for event := range eventParameters {
		if _, err = tm.UpsertTrigger(commontrigger.NewEvent(event)); err != nil {
			return nil, err
		}
	}

	return eventParameters, nil
}
