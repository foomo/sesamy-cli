package variable

import (
	"google.golang.org/api/tagmanager/v2"
)

func EventDataName(v string) string {
	return "event." + v
}

func NewEventData(name string) *tagmanager.Variable {
	return &tagmanager.Variable{
		Name: EventDataName(name),
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "setDefaultValue",
				Type:  "boolean",
				Value: "false",
			},
			{
				Key:   "keyPath",
				Type:  "template",
				Value: name,
			},
		},
		Type: "ed",
	}
}
