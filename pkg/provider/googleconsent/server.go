package googleconsent

import (
	"github.com/foomo/sesamy-cli/pkg/provider/googleconsent/server/template"
	"github.com/foomo/sesamy-cli/pkg/provider/googleconsent/server/variable"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
)

func ServerEnsure(tm *tagmanager.TagManager) error {
	folderName := tm.FolderName()
	defer tm.SetFolderName(folderName)

	{ // create folder
		if folder, err := tm.UpsertFolder("Sesamy - " + Name); err != nil {
			return err
		} else {
			tm.SetFolderName(folder.Name)
		}
	}

	{ // create clients
		consentTemplate, err := tm.UpsertCustomTemplate(template.NewGoogleConsentModeCheck(NameGoogleConsentModeCheckVariableTemplate))
		if err != nil {
			return err
		}
		if _, err = tm.UpsertVariable(variable.NewGoogleConsentModeAdStorage(consentTemplate)); err != nil {
			return err
		}
		if _, err = tm.UpsertVariable(variable.NewGoogleConsentModeAnalyticsStorage(consentTemplate)); err != nil {
			return err
		}
	}

	return nil
}
