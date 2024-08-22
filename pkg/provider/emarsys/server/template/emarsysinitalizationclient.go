package template

import (
	"fmt"

	"google.golang.org/api/tagmanager/v2"
)

func NewEmarsysInitializationClient(name string) *tagmanager.CustomTemplate {
	return &tagmanager.CustomTemplate{
		Name:         name,
		TemplateData: fmt.Sprintf(EmarsysInitializationClientData, name),
	}
}
