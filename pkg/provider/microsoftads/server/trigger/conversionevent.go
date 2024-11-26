package trigger

import (
	"google.golang.org/api/tagmanager/v2"
)

func ConversionEventName(v string) string {
	return "MAds Conversion - " + v
}

type (
	ConversionEventOptions struct {
		consentMode *tagmanager.Variable
	}
	ConversionEventOption func(*ConversionEventOptions)
)

func ConversionEventWithConsentMode(mode *tagmanager.Variable) ConversionEventOption {
	return func(o *ConversionEventOptions) {
		o.consentMode = mode
	}
}

func NewConversionEvent(name string, opts ...ConversionEventOption) *tagmanager.Trigger {
	o := &ConversionEventOptions{}
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
		Name: ConversionEventName(name),
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
