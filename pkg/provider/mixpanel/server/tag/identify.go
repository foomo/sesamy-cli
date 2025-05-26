package tag

import (
	"github.com/foomo/sesamy-cli/pkg/utils"
	"google.golang.org/api/tagmanager/v2"
)

func IdentifyName(v string) string {
	return "Mixpanel Identify - " + v
}

func NewIdentify(name string, identifier, projectToken *tagmanager.Variable, template *tagmanager.CustomTemplate, triggers ...*tagmanager.Trigger) *tagmanager.Tag {
	return &tagmanager.Tag{
		FiringTriggerId: utils.TriggerIDs(triggers),
		Name:            IdentifyName(name),
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
				Key:   "identifier",
				Type:  "template",
				Value: "{{" + identifier.Name + "}}",
			},
			{
				Key:   "identifyAuto",
				Type:  "boolean",
				Value: "true",
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
