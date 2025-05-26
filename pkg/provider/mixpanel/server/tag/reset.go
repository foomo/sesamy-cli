package tag

import (
	"github.com/foomo/sesamy-cli/pkg/utils"
	"google.golang.org/api/tagmanager/v2"
)

func ResetName(v string) string {
	return "Mixpanel Reset - " + v
}

func NewReset(name string, projectToken *tagmanager.Variable, template *tagmanager.CustomTemplate, triggers ...*tagmanager.Trigger) *tagmanager.Tag {
	return &tagmanager.Tag{
		FiringTriggerId: utils.TriggerIDs(triggers),
		Name:            ResetName(name),
		TagFiringOption: "oncePerEvent",
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "serverEU",
				Type:  "boolean",
				Value: "true",
			},
			{
				Key:   "logType",
				Type:  "template",
				Value: "debug",
			},
			{
				Key:   "type",
				Type:  "template",
				Value: "reset",
			},
			{
				Key:   "token",
				Type:  "template",
				Value: "{{" + projectToken.Name + "}}",
			},
		},
		Type: utils.TemplateType(template),
	}
}
