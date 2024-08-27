package trigger

import (
	"google.golang.org/api/tagmanager/v2"
)

type (
	ConversionLinkerEventOptions struct {
		consentMode *tagmanager.Variable
	}
	ConversionLinkerEventOption func(*ConversionLinkerEventOptions)
)

func ConversionLinkerEventWithConsentMode(mode *tagmanager.Variable) ConversionLinkerEventOption {
	return func(o *ConversionLinkerEventOptions) {
		o.consentMode = mode
	}
}

func NewConversionLinkerEvent(name string, opts ...ConversionLinkerEventOption) *tagmanager.Trigger {
	o := &ConversionLinkerEventOptions{}
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
