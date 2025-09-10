package tag

import (
	"strconv"

	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"google.golang.org/api/tagmanager/v2"
)

func ConversionName(v string) string {
	return "Pinterest Conversion - " + v
}

func NewConversion(name string, cfg config.Pinterest, template *tagmanager.CustomTemplate, triggers ...*tagmanager.Trigger) *tagmanager.Tag {
	return &tagmanager.Tag{
		FiringTriggerId: utils.TriggerIDs(triggers),
		Name:            ConversionName(name),
		TagFiringOption: "oncePerEvent",
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "apiAccessToken",
				Type:  "template",
				Value: cfg.APIAccessToken,
			},
			{
				Key:   "testMode",
				Type:  "boolean",
				Value: strconv.FormatBool(cfg.TestModeEnabled),
			},
			{
				Key:   "eventName",
				Type:  "template",
				Value: "inherit",
			},
			{
				Key:  "logMode",
				Type: "template",
				Value: func(testMode bool) string {
					if testMode {
						return "log"
					}
					return "donotlog"
				}(cfg.TestModeEnabled),
			},
			{
				Key:   "overrideMode",
				Type:  "boolean",
				Value: "false",
			},
			{
				Key:   "advertiserId",
				Type:  "template",
				Value: cfg.AdvertiserID,
			},
		},
		Type: utils.TemplateType(template),
	}
}
