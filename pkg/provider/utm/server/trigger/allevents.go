package trigger

import (
	"google.golang.org/api/tagmanager/v2"
)

type (
	UtmAllEventsOptions struct {
		consentMode *tagmanager.Variable
	}
	UtmAllEventsOption func(*UtmAllEventsOptions)
)

func UtmAllEventsWithConsentMode(mode *tagmanager.Variable) UtmAllEventsOption {
	return func(o *UtmAllEventsOptions) {
		o.consentMode = mode
	}
}

func NewUtmAllEvents(name string, opts ...UtmAllEventsOption) *tagmanager.Trigger {
	o := &UtmAllEventsOptions{}

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
