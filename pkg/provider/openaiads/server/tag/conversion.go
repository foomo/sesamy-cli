package tag

import (
	"strconv"

	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"google.golang.org/api/tagmanager/v2"
)

func ConversionName(v string) string {
	return "OpenAI Ads Conversion - " + v
}

func NewConversion(name string, pixelID, apiKey *tagmanager.Variable, template *tagmanager.CustomTemplate, validateOnly bool, settings config.OpenAIAdsConversionTag, triggers ...*tagmanager.Trigger) *tagmanager.Tag {
	eventName := settings.EventName
	eventNameSetup := "standard"
	if eventName == "" {
		eventName = name
		eventNameSetup = "inherit"
	}

	return &tagmanager.Tag{
		FiringTriggerId: utils.TriggerIDs(triggers),
		Name:            ConversionName(name),
		TagFiringOption: "oncePerEvent",
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "pixelId",
				Type:  "template",
				Value: "{{" + pixelID.Name + "}}",
			},
			{
				Key:   "apiKey",
				Type:  "template",
				Value: "{{" + apiKey.Name + "}}",
			},
			{
				Key:   "eventNameSetup",
				Type:  "template",
				Value: eventNameSetup,
			},
			{
				Key:   "eventNameStandard",
				Type:  "template",
				Value: eventName,
			},
			{
				Key:   "actionSource",
				Type:  "template",
				Value: "web",
			},
			{
				Key:   "adStorageConsent",
				Type:  "template",
				Value: "optional",
			},
			{
				Key:   "validateOnly",
				Type:  "boolean",
				Value: strconv.FormatBool(validateOnly),
			},
			{
				Key:   "autoMapEventParameters",
				Type:  "boolean",
				Value: "true",
			},
			{
				Key:   "autoMapUserDataParameters",
				Type:  "boolean",
				Value: "true",
			},
			{
				Key:   "autoMapServerEventDataParameters",
				Type:  "boolean",
				Value: "true",
			},
			{
				Key:   "useOptimisticScenario",
				Type:  "boolean",
				Value: "false",
			},
			{
				Key:  "itemIdKey",
				Type: "template",
			},
			{
				Key:   "setBrowserIdCookie",
				Type:  "boolean",
				Value: "true",
			},
			{
				Key:   "setClickIdCookie",
				Type:  "boolean",
				Value: "true",
			},
			{
				Key:   "cookieDomain",
				Type:  "template",
				Value: "auto",
			},
			{
				Key:   "cookieExpiration",
				Type:  "template",
				Value: "30",
			},
			{
				Key:   "cookieExpirationBrowserId",
				Type:  "template",
				Value: "365",
			},
			{
				Key:   "cookieHttpOnly",
				Type:  "boolean",
				Value: "true",
			},
			{
				Key:   "cookieSameSite",
				Type:  "template",
				Value: "lax",
			},
		},
		Type: utils.TemplateType(template),
	}
}
