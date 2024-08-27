package trigger

import (
	"google.golang.org/api/tagmanager/v2"
)

func GoogleAnalyticsEventName(v string) string {
	return "GA4 - " + v
}

type (
	GoogleAnalyticsEventOptions struct {
		consentMode *tagmanager.Variable
	}
	GoogleAnalyticsEventOption func(*GoogleAnalyticsEventOptions)
)

func GoogleAnalyticsEventWithConsentMode(mode *tagmanager.Variable) GoogleAnalyticsEventOption {
	return func(o *GoogleAnalyticsEventOptions) {
		o.consentMode = mode
	}
}

func NewGoogleAnalyticsEvent(name string, opts ...GoogleAnalyticsEventOption) *tagmanager.Trigger {
	o := &GoogleAnalyticsEventOptions{}
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
		Name: GoogleAnalyticsEventName(name),
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
