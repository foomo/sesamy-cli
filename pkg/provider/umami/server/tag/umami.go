package tag

import (
	"maps"
	"slices"

	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"google.golang.org/api/tagmanager/v2"
)

func UmamiName(v string) string {
	return "Umami - " + v
}

// NewUmami returns the Umami tag for the given event.
// The keys of params are the event parameter names that the tag sends as Umami event data.
func NewUmami(name string, cfg config.Umami, template *tagmanager.CustomTemplate, params map[string]string, triggers ...*tagmanager.Trigger) *tagmanager.Tag {
	parameter := []*tagmanager.Parameter{
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
	}

	if len(params) > 0 {
		var list []*tagmanager.Parameter

		for _, key := range slices.Sorted(maps.Keys(params)) {
			if key == "" {
				continue
			}

			list = append(list, &tagmanager.Parameter{
				Type: "map",
				Map: []*tagmanager.Parameter{
					{
						Key:   "name",
						Type:  "template",
						Value: key,
					},
				},
			})
		}

		if len(list) > 0 {
			parameter = append(parameter, &tagmanager.Parameter{
				Key:  "eventParameters",
				Type: "list",
				List: list,
			})
		}
	}

	return &tagmanager.Tag{
		FiringTriggerId: utils.TriggerIDs(triggers),
		Name:            UmamiName(name),
		TagFiringOption: "oncePerEvent",
		Parameter:       parameter,
		Type:            utils.TemplateType(template),
	}
}
