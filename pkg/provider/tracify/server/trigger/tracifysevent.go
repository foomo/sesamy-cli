package trigger

import (
	"google.golang.org/api/tagmanager/v2"
)

func TracifyEventName(v string) string {
	return "Tracify - " + v
}

type (
	TracifyEventOptions struct {
		consentMode *tagmanager.Variable
	}
	TracifyEventOption func(*TracifyEventOptions)
)

func TracifyEventWithConsentMode(mode *tagmanager.Variable) TracifyEventOption {
	return func(o *TracifyEventOptions) {
		o.consentMode = mode
	}
}

func NewTracifyEvent(name string, opts ...TracifyEventOption) *tagmanager.Trigger {
	o := &TracifyEventOptions{}
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
		Name: TracifyEventName(name),
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
