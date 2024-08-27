package trigger

import (
	"google.golang.org/api/tagmanager/v2"
)

func FacebookEventName(v string) string {
	return "FB Conversion - " + v
}

type (
	FacebookEventOptions struct {
		consentMode *tagmanager.Variable
	}
	FacebookEventOption func(*FacebookEventOptions)
)

func FacebookEventWithConsentMode(mode *tagmanager.Variable) FacebookEventOption {
	return func(o *FacebookEventOptions) {
		o.consentMode = mode
	}
}

func NewFacebookEvent(name string, opts ...FacebookEventOption) *tagmanager.Trigger {
	o := &FacebookEventOptions{}
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
		Name: FacebookEventName(name),
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
