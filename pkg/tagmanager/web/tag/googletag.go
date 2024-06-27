package client

import (
	"google.golang.org/api/tagmanager/v2"
)

func NewGoogleTag(name string, measurementID *tagmanager.Variable, configSettings *tagmanager.Variable, extraConfigSettings map[string]string) *tagmanager.Tag {
	ret := &tagmanager.Tag{
		FiringTriggerId: []string{"2147479573"}, // TODO
		Name:            name,
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "tagId",
				Type:  "template",
				Value: "{{" + measurementID.Name + "}}",
			},
			{
				Key:   "configSettingsVariable",
				Type:  "template",
				Value: "{{" + configSettings.Name + "}}",
			},
			{
				Key:   "configSettingsTable",
				Type:  "template",
				Value: "{{" + configSettings.Name + "}}",
			},
		},
		Type: "googtag",
	}

	if len(extraConfigSettings) > 0 {
		var list []*tagmanager.Parameter
		for k, v := range extraConfigSettings {
			list = append(list, &tagmanager.Parameter{
				Type: "map",
				Map: []*tagmanager.Parameter{
					{
						Key:   "parameter",
						Type:  "template",
						Value: k,
					},
					{
						Key:   "parameterValue",
						Type:  "template",
						Value: v,
					},
				},
			})
		}
		ret.Parameter = append(ret.Parameter, &tagmanager.Parameter{
			Key:  "configSettingsTable",
			Type: "list",
			List: list,
		})
	}

	return ret
}
