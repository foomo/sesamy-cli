package tag

import (
	"strconv"

	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"google.golang.org/api/tagmanager/v2"
)

func ConversionsAPITagName(v string) string {
	return "FB Conversion - " + v
}

func NewConversionsAPITag(name string, pixelID, apiAccessToken, testEventCode *tagmanager.Variable, settings config.FacebookConversionAPITag, template *tagmanager.CustomTemplate, triggers ...*tagmanager.Trigger) *tagmanager.Tag {
	return &tagmanager.Tag{
		FiringTriggerId: utils.TriggerIDs(triggers),
		Name:            ConversionsAPITagName(name),
		TagFiringOption: "oncePerEvent",
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "pixelId",
				Type:  "template",
				Value: "{{" + pixelID.Name + "}}",
			},
			{
				Key:   "apiAccessToken",
				Type:  "template",
				Value: "{{" + apiAccessToken.Name + "}}",
			},
			{
				Key:   "testEventCode",
				Type:  "template",
				Value: "{{" + testEventCode.Name + "}}",
			},
			{
				Key:   "enableEventEnhancement",
				Type:  "boolean",
				Value: strconv.FormatBool(settings.EnableEventEnhancement),
			},
			{
				Key:   "extendCookies",
				Type:  "boolean",
				Value: strconv.FormatBool(settings.ExtendCookies),
			},
			{
				Key:   "actionSource",
				Type:  "template",
				Value: "website",
			},
		},
		Type: utils.TemplateType(template),
	}
}
