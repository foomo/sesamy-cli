package trigger

import (
	"google.golang.org/api/tagmanager/v2"
)

func EventName(v string) string {
	return "Mixpanel - " + v
}

type (
	EventOptions struct {
		consentMode *tagmanager.Variable
	}
	EventOption func(*EventOptions)
)

func EventWithConsentMode(mode *tagmanager.Variable) EventOption {
	return func(o *EventOptions) {
		o.consentMode = mode
	}
}

func NewEvent(name string, opts ...EventOption) *tagmanager.Trigger {
	o := &EventOptions{}
	for _, opt := range opts {
		if opt != nil {
			opt(o)
		}
	}

	var filter []*tagmanager.Condition
	if o.consentMode != nil {
		filter = append(filter,
			&tagmanager.Condition{
				Parameter: []*tagmanager.Parameter{
					{
						Key:   "arg0",
						Type:  "template",
						Value: "{{" + o.consentMode.Name + "}}",
					},
					{
						Key:   "arg1",
						Type:  "template",
						Value: "granted",
					},
				},
				Type: "equals",
			},
		)
	}

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
		Filter: filter,
	}
}
