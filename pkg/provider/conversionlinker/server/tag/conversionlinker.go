package tag

import (
	"github.com/foomo/sesamy-cli/pkg/utils"
	"google.golang.org/api/tagmanager/v2"
)

func NewConversionLinker(name string, triggers ...*tagmanager.Trigger) *tagmanager.Tag {
	return &tagmanager.Tag{
		FiringTriggerId: utils.TriggerIDs(triggers),
		Name:            name,
		TagFiringOption: "oncePerEvent",
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "enableLinkerParams",
				Type:  "boolean",
				Value: "false",
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
