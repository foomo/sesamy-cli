package googletagmanager

import (
	"context"

	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	commonvariable "github.com/foomo/sesamy-cli/pkg/tagmanager/common/variable"
	"github.com/foomo/sesamy-cli/pkg/tagmanager/web/variable"
)

func Web(ctx context.Context, tm *tagmanager.TagManager, cfg config.GoogleTagManager) error {
	folder, err := tm.UpsertFolder(ctx, "Sesamy - "+Name)
	if err != nil {
		return err
	}

	{ // create variables
		for _, value := range cfg.WebContaienrVariables.DataLayer {
			if _, err := tm.UpsertVariable(ctx, folder, variable.NewDataLayer(value)); err != nil {
				return err
			}
		}
		for key, value := range cfg.WebContaienrVariables.LookupTables {
			if _, err := tm.UpsertVariable(ctx, folder, commonvariable.NewLookupTable(key, value)); err != nil {
				return err
			}
		}
	}

	return nil
}
