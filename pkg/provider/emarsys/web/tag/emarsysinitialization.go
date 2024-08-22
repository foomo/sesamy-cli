package tag

import (
	"github.com/foomo/sesamy-cli/pkg/tagmanager/web/trigger"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"google.golang.org/api/tagmanager/v2"
)

func NewEmarsysInitialization(name string, template *tagmanager.CustomTemplate) *tagmanager.Tag {
	ret := &tagmanager.Tag{
		Name:            name,
		FiringTriggerId: []string{trigger.IDConsentInitializtion},
		TagFiringOption: "oncePerEvent",
		Type:            utils.TemplateType(template),
	}
	return ret
}
