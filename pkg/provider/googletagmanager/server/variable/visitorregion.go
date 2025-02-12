package variable

import (
	"google.golang.org/api/tagmanager/v2"
)

func NewVisitorRegion(name string) *tagmanager.Variable {
	return &tagmanager.Variable{
		Name: name,
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "resolutionOptions",
				Type:  "template",
				Value: "requestHeaders",
			},
		},
		Type: "vr",
	}
}
