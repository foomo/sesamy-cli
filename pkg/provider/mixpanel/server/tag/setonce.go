package tag

import (
	"maps"
	"slices"

	"github.com/foomo/sesamy-cli/pkg/utils"
	"google.golang.org/api/tagmanager/v2"
)

func SetOnceName(v string) string {
	return "Mixpanel SetOnce - " + v
}

func NewSetOnce(name string, projectToken *tagmanager.Variable, template *tagmanager.CustomTemplate, params map[string]*tagmanager.Variable, triggers ...*tagmanager.Trigger) *tagmanager.Tag {
	parameter := []*tagmanager.Parameter{
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
			Key:   "identifyAuto",
			Type:  "boolean",
			Value: "true",
		},
		{
			Key:   "trackName",
			Type:  "template",
			Value: name,
		},
		{
			Key:   "type",
			Type:  "template",
			Value: "profile-set-once",
		},
		{
			Key:   "token",
			Type:  "template",
			Value: "{{" + projectToken.Name + "}}",
		},
		{
			Key:   "trackFromVariable",
			Type:  "boolean",
			Value: "false",
		},
	}

	if len(params) > 0 {
		var list []*tagmanager.Parameter
		for _, key := range slices.Sorted(maps.Keys(params)) {
			param := params[key]
			list = append(list, &tagmanager.Parameter{
				Type: "map",
				Map: []*tagmanager.Parameter{
					{
						Key:   "name",
						Type:  "template",
						Value: key,
					},
					{
						Key:   "value",
						Type:  "template",
						Value: "{{" + param.Name + "}}",
					},
				},
			})
		}
		parameter = append(parameter, &tagmanager.Parameter{
			Key:  "userPropertiesToSetOnce",
			Type: "list",
			List: list,
		})
	}

	return &tagmanager.Tag{
		FiringTriggerId: utils.TriggerIDs(triggers),
		Name:            SetOnceName(name),
		TagFiringOption: "oncePerEvent",
		Parameter:       parameter,
		Type:            utils.TemplateType(template),
	}
}
