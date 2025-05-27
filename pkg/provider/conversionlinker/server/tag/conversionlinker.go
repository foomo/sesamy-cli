package tag

import (
	"strconv"

	"github.com/foomo/sesamy-cli/pkg/utils"
	"google.golang.org/api/tagmanager/v2"
)

func NewConversionLinker(name string, enableLinkerParams bool, triggers ...*tagmanager.Trigger) *tagmanager.Tag {
	return &tagmanager.Tag{
		FiringTriggerId: utils.TriggerIDs(triggers),
		Name:            name,
		TagFiringOption: "oncePerEvent",
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "enableLinkerParams",
				Type:  "boolean",
				Value: strconv.FormatBool(enableLinkerParams),
			},
			{
				Key:   "enableCookieOverrides",
				Type:  "boolean",
				Value: "false",
			},
		},
		Type: "sgtmadscl",
	}
}
