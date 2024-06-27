package client

import (
	"google.golang.org/api/tagmanager/v2"
)

func NewGA4Event(name, eventName string, eventSettings *tagmanager.Variable, measurementID *tagmanager.Variable, trigger *tagmanager.Trigger) *tagmanager.Tag {
	return &tagmanager.Tag{
		FiringTriggerId: []string{trigger.TriggerId},
		Name:            name,
		Parameter: []*tagmanager.Parameter{
			{
				Type:  "boolean",
				Key:   "sendEcommerceData",
				Value: "false",
			},
			{
				Type:  "boolean",
				Key:   "enhancedUserId",
				Value: "false",
			},
			{
				Type:  "template",
				Key:   "eventName",
				Value: eventName,
			},
			{
				Type:  "template",
				Key:   "measurementIdOverride",
				Value: "{{" + measurementID.Name + "}}",
			},
			{
				Type:  "template",
				Key:   "eventSettingsVariable",
				Value: "{{" + eventSettings.Name + "}}",
			},
		},
		Type: "gaawe",
	}
}
