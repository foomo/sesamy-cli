package tag

import (
	"strconv"

	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"google.golang.org/api/tagmanager/v2"
)

func NewGoogleAdsRemarketing(name string, conversionID *tagmanager.Variable, cfg config.GoogleAdsRemarketing, triggers ...*tagmanager.Trigger) *tagmanager.Tag {
	return &tagmanager.Tag{
		FiringTriggerId: utils.TriggerIDs(triggers),
		Name:            name,
		TagFiringOption: "oncePerEvent",
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "enableConversionLinker",
				Type:  "boolean",
				Value: strconv.FormatBool(cfg.EnableConversionLinker),
			},
			{
				Key:   "enableDynamicRemarketing",
				Type:  "boolean",
				Value: "false",
			},
			{
				Key:   "enableCustomParams",
				Type:  "boolean",
				Value: "false",
			},
			{
				Key:   "enableUserId",
				Type:  "boolean",
				Value: "false",
			},
			{
				Key:   "conversionId",
				Type:  "template",
				Value: "{{" + conversionID.Name + "}}",
			},
			{
				Key:   "rdp",
				Type:  "boolean",
				Value: "false",
			},
		},
		Type: "sgtmadsremarket",
	}
}
