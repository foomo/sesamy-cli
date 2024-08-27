package variable

import (
	"github.com/foomo/sesamy-cli/pkg/utils"
	"google.golang.org/api/tagmanager/v2"
)

func NewGoogleConsentModeAnalyticsStorage(template *tagmanager.CustomTemplate) *tagmanager.Variable {
	return &tagmanager.Variable{
		Name: GoogleConsentModeName("analytics_storage"),
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "consentType",
				Type:  "template",
				Value: "analytics_storage",
			},
		},
		Type: utils.TemplateType(template),
	}
}
