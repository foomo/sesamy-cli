package tag

import (
	"github.com/foomo/sesamy-cli/pkg/utils"
	"google.golang.org/api/tagmanager/v2"
)

func EventsAPITagName(v string) string {
	return "Criteo - " + v
}

func NewEventsAPITag(name string, callerID, partnerID *tagmanager.Variable, template *tagmanager.CustomTemplate, triggers ...*tagmanager.Trigger) *tagmanager.Tag {
	return &tagmanager.Tag{
		FiringTriggerId: utils.TriggerIDs(triggers),
		Name:            EventsAPITagName(name),
		TagFiringOption: "oncePerEvent",
		Parameter: []*tagmanager.Parameter{
			{
				Key:  "country",
				Type: "template",
			},
			{
				Key:  "language",
				Type: "template",
			},
			{
				Key:   "partnerId",
				Type:  "template",
				Value: "{{" + partnerID.Name + "}}",
			},
			{
				Key:   "enableDising",
				Type:  "boolean",
				Value: "false",
			},
			{
				Key:  "applicationId",
				Type: "template",
			},
		},
		Type: utils.TemplateType(template),
	}
}
