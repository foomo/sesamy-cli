package tag

import (
	"fmt"

	"google.golang.org/api/tagmanager/v2"
)

func NewUmami(name, websiteID, domain, endpointURL string, template *tagmanager.CustomTemplate, triggers ...*tagmanager.Trigger) *tagmanager.Tag {
	triggerIDs := make([]string, len(triggers))
	for i, trigger := range triggers {
		triggerIDs[i] = trigger.TriggerId
	}

	return &tagmanager.Tag{
		FiringTriggerId: triggerIDs,
		Name:            name,
		TagFiringOption: "oncePerEvent",
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "websiteId",
				Type:  "template",
				Value: websiteID,
			},
			{
				Key:   "endpointUrl",
				Type:  "template",
				Value: endpointURL,
			},
			{
				Key:   "domain",
				Type:  "template",
				Value: domain,
			},
			{
				Key:   "timeout",
				Type:  "template",
				Value: "1000",
			},
		},
		Type: fmt.Sprintf("cvt_%s_%s", template.ContainerId, template.TemplateId),
	}
}
