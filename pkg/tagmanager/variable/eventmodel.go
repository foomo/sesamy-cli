package variable

import (
	"google.golang.org/api/tagmanager/v2"
)

func NewEventModel(name, parameterName string) *tagmanager.Variable {
	return &tagmanager.Variable{
		Name: name,
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
				Value: "eventModel." + parameterName,
			},
		},
		Type: "v",
	}
}
