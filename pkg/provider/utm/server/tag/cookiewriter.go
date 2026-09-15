package tag

import (
	"github.com/foomo/sesamy-cli/pkg/utils"
	"google.golang.org/api/tagmanager/v2"
)

func NewUtmAttributionCookieWriter(name string, template *tagmanager.CustomTemplate, triggers ...*tagmanager.Trigger) *tagmanager.Tag {
	return &tagmanager.Tag{
		FiringTriggerId: utils.TriggerIDs(triggers),
		Name:            name,
		TagFiringOption: "oncePerEvent",
		Priority: &tagmanager.Parameter{
			Type:  "integer",
			Value: "100",
		},
		Type: utils.TemplateType(template),
	}
}
