package trigger

import (
	"google.golang.org/api/tagmanager/v2"
)

func EventName(v string) string {
	return "Event - " + v
}

func NewEvent(name string) *tagmanager.Trigger {
	return &tagmanager.Trigger{
		Type: "customEvent",
		Name: EventName(name),
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
						Value: name,
					},
				},
				Type: "equals",
			},
		},
	}
}
