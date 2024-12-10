package tag

import (
	"strconv"

	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"google.golang.org/api/tagmanager/v2"
)

func TracifyName(v string) string {
	return "Tracify - " + v
}

func NewTracify(name string, token, customerSiteID *tagmanager.Variable, template *tagmanager.CustomTemplate, cfg config.Tracify, triggers ...*tagmanager.Trigger) *tagmanager.Tag {
	return &tagmanager.Tag{
		FiringTriggerId: utils.TriggerIDs(triggers),
		Name:            TracifyName(name),
		TagFiringOption: "oncePerEvent",
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "customerSiteId",
				Type:  "template",
				Value: "{{" + customerSiteID.Name + "}}",
			},
			{
				Key:   "analyticsStorageConsent",
				Type:  "template",
				Value: "optional",
			},
			{
				Key:   "isStagingMode",
				Type:  "boolean",
				Value: strconv.FormatBool(cfg.StagingModeEnabled),
			},
			{
				Key:   "token",
				Type:  "template",
				Value: "{{" + token.Name + "}}",
			},
		},
		Type: utils.TemplateType(template),
	}
}
