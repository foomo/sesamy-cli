package tag

import (
	"google.golang.org/api/tagmanager/v2"
)

func NewGoogleAnalyticsGA4(name string, measurementID *tagmanager.Variable, triggers ...*tagmanager.Trigger) *tagmanager.Tag {
	triggerIDs := make([]string, len(triggers))
	for i, trigger := range triggers {
		triggerIDs[i] = trigger.TriggerId
	}

	return &tagmanager.Tag{
		FiringTriggerId: triggerIDs,
		Name:            name,
		Parameter: []*tagmanager.Parameter{
			{
				Type:  "boolean",
				Key:   "redactVisitorIp",
				Value: "false",
			},
			{
				Type:  "template",
				Key:   "epToIncludeDropdown",
				Value: "all",
			},
			{
				Type:  "boolean",
				Key:   "enableGoogleSignals",
				Value: "false",
			},
			{
				Type:  "template",
				Key:   "upToIncludeDropdown",
				Value: "all",
			},
			{
				Type:  "template",
				Key:   "measurementId",
				Value: "{{" + measurementID.Name + "}}",
			},
			{
				Type:  "boolean",
				Key:   "enableEuid",
				Value: "false",
			},
		},
		Type: "sgtmgaaw",
	}
}
