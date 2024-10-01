package tag

import (
	"github.com/foomo/sesamy-cli/pkg/utils"
	"google.golang.org/api/tagmanager/v2"
)

func TracifyName(v string) string {
	return "Tracify - " + v
}

func NewTracify(name string, customerSiteID, token *tagmanager.Variable, template *tagmanager.CustomTemplate, triggers ...*tagmanager.Trigger) *tagmanager.Tag {
	return &tagmanager.Tag{
		FiringTriggerId: utils.TriggerIDs(triggers),
		Name:            TracifyName(name),
		TagFiringOption: "oncePerEvent",
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "customerSiteId",
				Type:  "template",
				Value: "{{" + customerSiteID.Name + "}}",
			},
			{
				Key:   "analyticsStorageConsent",
				Type:  "template",
				Value: "optional",
			},
			{
				Key:   "isStagingMode",
				Type:  "boolean",
				Value: "false",
			},
			{
				Key:   "token",
				Type:  "template",
				Value: "{{" + token.Name + "}}",
			},
		},
		Type: utils.TemplateType(template),
	}
}
