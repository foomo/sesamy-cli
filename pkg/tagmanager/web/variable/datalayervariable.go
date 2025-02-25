package variable

import (
	"google.golang.org/api/tagmanager/v2"
)

func DataLayerName(v string) string {
	return "dlv." + v
}

func NewDataLayer(name string) *tagmanager.Variable {
	return &tagmanager.Variable{
		Name: DataLayerName(name),
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
				Value: name,
			},
		},
		Type: "v",
	}
}
