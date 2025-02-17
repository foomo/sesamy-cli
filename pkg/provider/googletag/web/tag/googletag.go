package client

import (
	"github.com/foomo/sesamy-cli/pkg/tagmanager/web/trigger"
	"google.golang.org/api/tagmanager/v2"
)

func NewGoogleTag(name string, tagID *tagmanager.Variable, configSettings *tagmanager.Variable, eventSettings map[string]*tagmanager.Variable) *tagmanager.Tag {
	var eventSettingsList []*tagmanager.Parameter
	for k, v := range eventSettings {
		eventSettingsList = append(eventSettingsList, &tagmanager.Parameter{
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
					Value: "{{" + v.Name + "}}",
				},
			},
		})
	}

	ret := &tagmanager.Tag{
		FiringTriggerId: []string{trigger.IDInitialization},
		Name:            name,
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "tagId",
				Type:  "template",
				Value: "{{" + tagID.Name + "}}",
			},
			{
				Key:  "eventSettingsTable",
				Type: "list",
				List: eventSettingsList,
			},
			{
				Key:   "configSettingsVariable",
				Type:  "template",
				Value: "{{" + configSettings.Name + "}}",
			},
		},
		Type: "googtag",
	}
	return ret
}
