package client

import (
	"github.com/foomo/sesamy-cli/pkg/tagmanager/web/trigger"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"google.golang.org/api/tagmanager/v2"
)

func NewUserIdentification(name string, callerID, partnerID *tagmanager.Variable, template *tagmanager.CustomTemplate) *tagmanager.Tag {
	ret := &tagmanager.Tag{
		FiringTriggerId: []string{trigger.IDInitialization},
		Name:            name,
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "psbEnabled",
				Type:  "boolean",
				Value: "true",
			},
			{
				Key:   "callerId",
				Type:  "template",
				Value: "{{" + callerID.Name + "}}",
			},
			{
				Key:   "partnerId",
				Type:  "template",
				Value: "{{" + partnerID.Name + "}}",
			},
		},
		Type: utils.TemplateType(template),
	}
	return ret
}
