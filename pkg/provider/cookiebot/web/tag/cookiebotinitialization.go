package tag

import (
	"strconv"

	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/tagmanager/web/trigger"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"google.golang.org/api/tagmanager/v2"
)

func NewCookiebotInitialization(name string, cfg config.Cookiebot, template *tagmanager.CustomTemplate) *tagmanager.Tag {
	return &tagmanager.Tag{
		Name:            name,
		FiringTriggerId: []string{trigger.IDConsentInitializtion},
		TagFiringOption: "oncePerEvent",
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "adsDataRedaction",
				Type:  "template",
				Value: "dynamic",
			},
			{
				Key:   "addGeoRegion",
				Type:  "boolean",
				Value: "false",
			},
			{
				Key:   "serial",
				Type:  "template",
				Value: cfg.CookiebotID,
			},
			{
				Key:   "iabFramework",
				Type:  "boolean",
				Value: "false",
			},
			{
				Key:   "cdnRegion",
				Type:  "template",
				Value: cfg.CDNRegion,
			},
			{
				Key:   "advertiserConsentModeEnabled",
				Type:  "boolean",
				Value: "true",
			},
			{
				Key:   "language",
				Type:  "template",
				Value: "auto",
			},
			{
				Key:   "urlPassthrough",
				Type:  "boolean",
				Value: strconv.FormatBool(cfg.URLPassthrough),
			},
			{
				Key:   "consentModeEnabled",
				Type:  "boolean",
				Value: "true",
			},
			{
				Key:   "waitForUpdate",
				Type:  "template",
				Value: "2000",
			},
		},
		Type: utils.TemplateType(template),
	}
}
