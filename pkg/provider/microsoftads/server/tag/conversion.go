package tag

import (
	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"google.golang.org/api/tagmanager/v2"
)

func ConversionName(v string) string {
	return "MAds Conversion - " + v
}

func NewConversion(name string, tagID *tagmanager.Variable, template *tagmanager.CustomTemplate, settings config.MicrosoftAdsConversionTag, triggers ...*tagmanager.Trigger) *tagmanager.Tag {
	return &tagmanager.Tag{
		FiringTriggerId: utils.TriggerIDs(triggers),
		Name:            ConversionName(name),
		TagFiringOption: "oncePerEvent",
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "pagetype",
				Type:  "template",
				Value: settings.PageType,
			},
			{
				Key:   "activateLogs",
				Type:  "boolean",
				Value: "false",
			},
			{
				Key:  "pageTitle",
				Type: "template",
			},
			{
				Key:   "prodidGa",
				Type:  "boolean",
				Value: "true",
			},
			{
				Key:  "vid",
				Type: "template",
			},
			{
				Key:  "searchTerm",
				Type: "template",
			},
			{
				Key:  "msclkidCookie",
				Type: "template",
			},
			{
				Key:  "msclkidCookie",
				Type: "template",
			},
			{
				Key:   "itemsGa",
				Type:  "boolean",
				Value: "true",
			},
			{
				Key:   "ea",
				Type:  "template",
				Value: name,
			},
			{
				Key:   "ea",
				Type:  "template",
				Value: name,
			},
			{
				Key:  "ecommCategory",
				Type: "template",
			},
			{
				Key:  "gc",
				Type: "template",
			},
			{
				Key:  "ec",
				Type: "template",
			},
			{
				Key:   "evt",
				Type:  "template",
				Value: settings.EventType,
			},
			{
				Key:   "firstClick",
				Type:  "boolean",
				Value: "false",
			},
			{
				Key:   "spa",
				Type:  "boolean",
				Value: "false",
			},
			{
				Key:  "el",
				Type: "template",
			},
			{
				Key:  "prodid",
				Type: "template",
			},
			{
				Key:  "pagePath",
				Type: "template",
			},
			{
				Key:  "gv",
				Type: "template",
			},
			{
				Key:  "ev",
				Type: "template",
			},
			{
				Key:  "pageLocation",
				Type: "template",
			},
			{
				Key:  "msclkidQuery",
				Type: "template",
			},
			{
				Key:   "createCookie",
				Type:  "boolean",
				Value: "false",
			},
			{
				Key:   "ti",
				Type:  "template",
				Value: "{{" + tagID.Name + "}}",
			},
			{
				Key:  "user_id",
				Type: "template",
			},
			{
				Key:  "pageReferrer",
				Type: "template",
			},
		},
		Type: utils.TemplateType(template),
	}
}
