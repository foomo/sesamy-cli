package template

import (
	"fmt"

	"google.golang.org/api/tagmanager/v2"
)

func NewJSONRequestValue(name string) *tagmanager.CustomTemplate {
	return &tagmanager.CustomTemplate{
		Name:         name,
		TemplateData: fmt.Sprintf(JSONRequestValueData, name),
	}
}