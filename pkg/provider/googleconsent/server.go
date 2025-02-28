package googleconsent

import (
	"github.com/foomo/sesamy-cli/pkg/provider/googleconsent/server/template"
	"github.com/foomo/sesamy-cli/pkg/provider/googleconsent/server/variable"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
)

func ServerEnsure(tm *tagmanager.TagManager) error {
	folder, err := tm.UpsertFolder("Sesamy - " + Name)
	if err != nil {
		return err
	}

	{ // create clients
		consentTemplate, err := tm.UpsertCustomTemplate(template.NewGoogleConsentModeCheck(NameGoogleConsentModeCheckVariableTemplate))
		if err != nil {
			return err
		}
		if _, err = tm.UpsertVariable(folder, variable.NewGoogleConsentModeAdStorage(consentTemplate)); err != nil {
			return err
		}
		if _, err = tm.UpsertVariable(folder, variable.NewGoogleConsentModeAnalyticsStorage(consentTemplate)); err != nil {
			return err
		}
	}

	return nil
}
