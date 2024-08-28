package trigger

import (
	"google.golang.org/api/tagmanager/v2"
)

type (
	GoogleAdsRemarketingEventOptions struct {
		consentMode *tagmanager.Variable
	}
	GoogleAdsRemarketingEventOption func(*GoogleAdsRemarketingEventOptions)
)

func GoogleAdsRemarketingEventWithConsentMode(mode *tagmanager.Variable) GoogleAdsRemarketingEventOption {
	return func(o *GoogleAdsRemarketingEventOptions) {
		o.consentMode = mode
	}
}

func NewGoogleAdsRemarketingEvent(name string, opts ...GoogleAdsRemarketingEventOption) *tagmanager.Trigger {
	o := &GoogleAdsRemarketingEventOptions{}
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
		Type:   "serverPageview",
		Name:   name,
		Filter: filter,
	}
}
