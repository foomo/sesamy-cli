package variable

import (
	"github.com/foomo/sesamy-cli/pkg/utils"
	"google.golang.org/api/tagmanager/v2"
)

func MPv2DataName(v string) string {
	return "mpv2." + v
}

func NewMPv2Data(name string, template *tagmanager.CustomTemplate) *tagmanager.Variable {
	return &tagmanager.Variable{
		Name: name,
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "key",
				Type:  "template",
				Value: name,
			},
		},
		Type: utils.TemplateType(template),
	}
}
