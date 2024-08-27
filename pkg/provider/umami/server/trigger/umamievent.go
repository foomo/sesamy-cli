package trigger

import (
	"google.golang.org/api/tagmanager/v2"
)

func UmamiEventName(v string) string {
	return "Umami - " + v
}

type (
	UmamiEventOptions struct {
		consentMode *tagmanager.Variable
	}
	UmamiEventOption func(*UmamiEventOptions)
)

func UmamiEventWithConsentMode(mode *tagmanager.Variable) UmamiEventOption {
	return func(o *UmamiEventOptions) {
		o.consentMode = mode
	}
}

func NewUmamiEvent(name string, opts ...UmamiEventOption) *tagmanager.Trigger {
	o := &UmamiEventOptions{}
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
		Name: UmamiEventName(name),
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
