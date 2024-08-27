package tag

import (
	"strconv"

	"github.com/foomo/sesamy-cli/pkg/utils"
	"google.golang.org/api/tagmanager/v2"
)

func GoogleAnalyticsName(v string) string {
	return "GA4 - " + v
}

func NewGoogleAnalytics(name string, redactVisitorIP bool, triggers ...*tagmanager.Trigger) *tagmanager.Tag {
	return &tagmanager.Tag{
		FiringTriggerId: utils.TriggerIDs(triggers),
		Name:            GoogleAnalyticsName(name),
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
