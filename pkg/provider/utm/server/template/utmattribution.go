package template

import (
	"fmt"

	"google.golang.org/api/tagmanager/v2"
)

func NewUtmAttribution(name string) *tagmanager.CustomTemplate {
	return &tagmanager.CustomTemplate{
		Name:         name,
		TemplateData: fmt.Sprintf(UtmAttributionData, name),
	}
}
