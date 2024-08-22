package client

import (
	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"google.golang.org/api/tagmanager/v2"
)

func NewEmarsys(name string, cfg config.Emarsys, template *tagmanager.CustomTemplate) *tagmanager.Client {
	return &tagmanager.Client{
		Name: name,
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "merchantId",
				Type:  "template",
				Value: cfg.MerchantID,
			},
		},
		Type: utils.TemplateType(template),
	}
}
