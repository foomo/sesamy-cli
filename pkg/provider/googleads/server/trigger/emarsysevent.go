package trigger

import (
	"google.golang.org/api/tagmanager/v2"
)

func GoogleAdsEventName(v string) string {
	return "GAds - " + v
}

type (
	GoogleAdsEventOptions struct {
		consentMode *tagmanager.Variable
	}
	GoogleAdsEventOption func(*GoogleAdsEventOptions)
)

func GoogleAdsEventWithConsentMode(mode *tagmanager.Variable) GoogleAdsEventOption {
	return func(o *GoogleAdsEventOptions) {
		o.consentMode = mode
	}
}

func NewGoogleAdsEvent(name string, opts ...GoogleAdsEventOption) *tagmanager.Trigger {
	o := &GoogleAdsEventOptions{}
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
		Name: GoogleAdsEventName(name),
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
