package trigger

import (
	"google.golang.org/api/tagmanager/v2"
)

func CriteoEventName(v string) string {
	return "Criteo - " + v
}

type (
	CriteoEventOptions struct {
		consentMode *tagmanager.Variable
	}
	CriteoEventOption func(*CriteoEventOptions)
)

func CriteoEventWithConsentMode(mode *tagmanager.Variable) CriteoEventOption {
	return func(o *CriteoEventOptions) {
		o.consentMode = mode
	}
}

func NewCriteoEvent(name string, opts ...CriteoEventOption) *tagmanager.Trigger {
	o := &CriteoEventOptions{}
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
		Name: CriteoEventName(name),
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
