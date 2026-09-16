package criteo

import (
	"context"
	"errors"
	"log/slog"

	"github.com/foomo/sesamy-cli/pkg/config"
	client "github.com/foomo/sesamy-cli/pkg/provider/criteo/web/tag"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	commontemplate "github.com/foomo/sesamy-cli/pkg/tagmanager/common/template"
	commonvariable "github.com/foomo/sesamy-cli/pkg/tagmanager/common/variable"
	tagmanager2 "google.golang.org/api/tagmanager/v2"
)

func Web(ctx context.Context, l *slog.Logger, tm *tagmanager.TagManager, cfg config.Criteo) error {
	folder, err := tm.UpsertFolder(ctx, "Sesamy - "+Name)
	if err != nil {
		return err
	}

	template, err := commontemplate.ResolveOrLookup(ctx, tm, cfg.Templates.UserIdentification, NameCriteoUserIdentificationTemplate, func() (*tagmanager2.CustomTemplate, error) {
		return tm.LookupTemplate(ctx, NameCriteoUserIdentificationTemplate)
	})
	if err != nil {
		if errors.Is(err, tagmanager.ErrNotFound) {
			l.Warn("Please install the 'Criteo User Identification' template manually first")
		}

		return err
	}

	{ // setup criteo
		callerID, err := tm.UpsertVariable(ctx, folder, commonvariable.NewConstant(NameCallerID, cfg.CallerID))
		if err != nil {
			return err
		}

		partnerID, err := tm.UpsertVariable(ctx, folder, commonvariable.NewConstant(NamePartnerID, cfg.PartnerID))
		if err != nil {
			return err
		}

		if _, err = tm.UpsertTag(ctx, folder, client.NewUserIdentification(NameCriteoUserIdentificationTag, callerID, partnerID, template)); err != nil {
			return err
		}
	}

	return nil
}
