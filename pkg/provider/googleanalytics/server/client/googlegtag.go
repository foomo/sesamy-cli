package client

import (
	"strconv"

	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"google.golang.org/api/tagmanager/v2"
)

func NewGoogleGTag(name string, cfg config.GoogleGTag, template *tagmanager.CustomTemplate) *tagmanager.Client {
	return &tagmanager.Client{
		Name:     name,
		Priority: cfg.Priority,
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "patchEcommerceItems",
				Type:  "boolean",
				Value: strconv.FormatBool(cfg.EcommerceItems),
			},
		},
		Type: utils.TemplateType(template),
	}
}
