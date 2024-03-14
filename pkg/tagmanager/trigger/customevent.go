package trigger

import (
	"google.golang.org/api/tagmanager/v2"
)

func NewCustomEvent(name, eventName string) *tagmanager.Trigger {
	return &tagmanager.Trigger{
		Type: "customEvent",
		Name: name,
		CustomEventFilter: []*tagmanager.Condition{
			{
				Parameter: []*tagmanager.Parameter{
					{
						Key:   "arg0",
						Type:  "template",
						Value: "{{_event}}",
					},
					{
						Key:   "arg1",
						Type:  "template",
						Value: eventName,
					},
				},
				Type: "equals",
			},
		},
	}
}
