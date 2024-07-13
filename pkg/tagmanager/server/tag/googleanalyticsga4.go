package tag

import (
	"strconv"

	"github.com/foomo/sesamy-cli/pkg/utils"
	"google.golang.org/api/tagmanager/v2"
)

func GoogleAnalyticsGA4Name(v string) string {
	return "GA4 - " + v
}

func NewGoogleAnalyticsGA4(name string, redactVisitorIP bool, triggers ...*tagmanager.Trigger) *tagmanager.Tag {
	return &tagmanager.Tag{
		FiringTriggerId: utils.TriggerIDs(triggers),
		Name:            GoogleAnalyticsGA4Name(name),
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "redactVisitorIp",
				Type:  "boolean",
				Value: strconv.FormatBool(redactVisitorIP),
			},
			{
				Key:   "epToIncludeDropdown",
				Type:  "template",
				Value: "all",
			},
			{
				Key:   "upToIncludeDropdown",
				Type:  "template",
				Value: "all",
			},
		},
		Type: "sgtmgaaw",
	}
}
