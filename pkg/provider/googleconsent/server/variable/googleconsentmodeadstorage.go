package variable

import (
	"github.com/foomo/sesamy-cli/pkg/utils"
	"google.golang.org/api/tagmanager/v2"
)

func NewGoogleConsentModeAdStorage(template *tagmanager.CustomTemplate) *tagmanager.Variable {
	return &tagmanager.Variable{
		Name: GoogleConsentModeName("ad_storage"),
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "consentType",
				Type:  "template",
				Value: "ad_storage",
			},
		},
		Type: utils.TemplateType(template),
	}
}
