package tag

import (
	"google.golang.org/api/tagmanager/v2"
)

func GoogleAnalyticsGA4EventName(v string) string {
	return "GA4 Event - " + v
}

func NewGoogleAnalyticsGA4Event(name string, tagID, settings *tagmanager.Variable, trigger *tagmanager.Trigger) *tagmanager.Tag {
	return &tagmanager.Tag{
		FiringTriggerId: []string{trigger.TriggerId},
		Name:            GoogleAnalyticsGA4EventName(name),
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
				Value: name,
			},
			{
				Type:  "template",
				Key:   "measurementIdOverride",
				Value: "{{" + tagID.Name + "}}",
			},
			{
				Type:  "template",
				Key:   "eventSettingsVariable",
				Value: "{{" + settings.Name + "}}",
			},
		},
		Type: "gaawe",
	}
}
