package tag

import (
	"github.com/foomo/sesamy-cli/pkg/utils"
	"google.golang.org/api/tagmanager/v2"
)

func EmarsysName(v string) string {
	return "Emarsys - " + v
}

func NewEmarsys(name string, merchantID *tagmanager.Variable, template *tagmanager.CustomTemplate, triggers ...*tagmanager.Trigger) *tagmanager.Tag {
	return &tagmanager.Tag{
		FiringTriggerId: utils.TriggerIDs(triggers),
		Name:            EmarsysName(name),
		TagFiringOption: "oncePerEvent",
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "adStorageConsent",
				Type:  "template",
				Value: "optional",
			},
			{
				Key:   "merchantId",
				Type:  "template",
				Value: "{{" + merchantID.Name + "}}",
			},
			{
				Key:   "isTestMode",
				Type:  "boolean",
				Value: "false",
			},
			{
				Key:   "isDebugMode",
				Type:  "boolean",
				Value: "false",
			},
		},
		Type: utils.TemplateType(template),
	}
}
