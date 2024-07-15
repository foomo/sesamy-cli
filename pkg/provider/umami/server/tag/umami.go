package tag

import (
	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"google.golang.org/api/tagmanager/v2"
)

func UmamiName(v string) string {
	return "Umami - " + v
}

func NewUmami(name string, cfg config.Umami, template *tagmanager.CustomTemplate, triggers ...*tagmanager.Trigger) *tagmanager.Tag {
	return &tagmanager.Tag{
		FiringTriggerId: utils.TriggerIDs(triggers),
		Name:            UmamiName(name),
		TagFiringOption: "oncePerEvent",
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "websiteId",
				Type:  "template",
				Value: cfg.WebsiteID,
			},
			{
				Key:   "endpointUrl",
				Type:  "template",
				Value: cfg.EndpointURL,
			},
			{
				Key:   "domain",
				Type:  "template",
				Value: cfg.Domain,
			},
			{
				Key:   "timeout",
				Type:  "template",
				Value: "1000",
			},
		},
		Type: utils.TemplateType(template),
	}
}
