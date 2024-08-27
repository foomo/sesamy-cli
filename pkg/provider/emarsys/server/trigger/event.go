package trigger

import (
	"google.golang.org/api/tagmanager/v2"
)

func EmarsysEventName(v string) string {
	return "Emarsys - " + v
}

type (
	EmarsysEventOptions struct {
		consentMode *tagmanager.Variable
	}
	EmarsysEventOption func(*EmarsysEventOptions)
)

func EmarsysEventWithConsentMode(mode *tagmanager.Variable) EmarsysEventOption {
	return func(o *EmarsysEventOptions) {
		o.consentMode = mode
	}
}

func NewEmarsysEvent(name string, opts ...EmarsysEventOption) *tagmanager.Trigger {
	o := &EmarsysEventOptions{}
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
		Name: EmarsysEventName(name),
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
