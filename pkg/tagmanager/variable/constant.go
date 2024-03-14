package variable

import (
	"google.golang.org/api/tagmanager/v2"
)

func NewConstant(name, value string) *tagmanager.Variable {
	return &tagmanager.Variable{
		Name: name,
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "value",
				Type:  "template",
				Value: value,
			},
		},
		Type: "c",
	}
}
