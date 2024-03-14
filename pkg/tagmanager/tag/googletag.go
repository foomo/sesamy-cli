package client

import (
	"google.golang.org/api/tagmanager/v2"
)

func NewGoogleTag(name string, measurementID *tagmanager.Variable, configSettings *tagmanager.Variable) *tagmanager.Tag {
	return &tagmanager.Tag{
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
		},
		Type: "googtag",
	}
}
