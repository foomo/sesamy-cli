package variable

import (
	"github.com/foomo/sesamy-cli/pkg/utils"
	"google.golang.org/api/tagmanager/v2"
)

func NewUtmAttribution(name string, template *tagmanager.CustomTemplate) *tagmanager.Variable {
	return &tagmanager.Variable{
		Name: name,
		Type: utils.TemplateType(template),
	}
}
