package variable

import (
	"google.golang.org/api/tagmanager/v2"
)

func GoogleTagEventModelName(v string) string {
	return "dlv.eventModel." + v
}

func NewGoogleTagEventModel(name string) *tagmanager.Variable {
	return &tagmanager.Variable{
		Name: GoogleTagEventModelName(name),
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "dataLayerVersion",
				Type:  "integer",
				Value: "2",
			},
			{
				Key:   "setDefaultValue",
				Type:  "boolean",
				Value: "false",
			},
			{
				Key:   "name",
				Type:  "template",
				Value: "eventModel." + name,
			},
		},
		Type: "v",
	}
}
