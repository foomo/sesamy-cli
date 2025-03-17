package googleconsent

import (
	"context"

	"github.com/foomo/sesamy-cli/pkg/provider/googleconsent/server/template"
	"github.com/foomo/sesamy-cli/pkg/provider/googleconsent/server/variable"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
)

func ServerEnsure(ctx context.Context, tm *tagmanager.TagManager) error {
	folder, err := tm.UpsertFolder(ctx, "Sesamy - "+Name)
	if err != nil {
		return err
	}

	{ // create clients
		consentTemplate, err := tm.UpsertCustomTemplate(ctx, template.NewGoogleConsentModeCheck(NameGoogleConsentModeCheckVariableTemplate))
		if err != nil {
			return err
		}
		if _, err = tm.UpsertVariable(ctx, folder, variable.NewGoogleConsentModeAdStorage(consentTemplate)); err != nil {
			return err
		}
		if _, err = tm.UpsertVariable(ctx, folder, variable.NewGoogleConsentModeAnalyticsStorage(consentTemplate)); err != nil {
			return err
		}
	}

	return nil
}
