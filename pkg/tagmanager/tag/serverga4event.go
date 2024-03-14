package client

import (
	"google.golang.org/api/tagmanager/v2"
)

func NewServerGA4Event(name string, measurementID *tagmanager.Variable, trigger *tagmanager.Trigger) *tagmanager.Tag {
	return &tagmanager.Tag{
		FiringTriggerId: []string{trigger.TriggerId},
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
