package tag

import (
	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/utils"
	"google.golang.org/api/tagmanager/v2"
)

func GoogleAdsConversionTrackingName(v string) string {
	return "GAds Conversion - " + v
}

func NewGoogleAdsConversionTracking(name string, value, currency, conversionID *tagmanager.Variable, settings config.GoogleAdsConversionTracking, triggers ...*tagmanager.Trigger) *tagmanager.Tag {
	tagName := GoogleAdsConversionTrackingName(name)
	tagConversionID := "{{" + conversionID.Name + "}}"
	if settings.ConversionID != "" {
		tagName += " (" + settings.ConversionID + ")"
		tagConversionID = settings.ConversionID
	}

	return &tagmanager.Tag{
		FiringTriggerId: utils.TriggerIDs(triggers),
		Name:            tagName,
		TagFiringOption: "oncePerEvent",
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "enableNewCustomerReporting",
				Type:  "boolean",
				Value: "false",
			},
			{
				Key:   "enableConversionLinker",
				Type:  "boolean",
				Value: "true",
			},
			{
				Key:   "enableProductReporting",
				Type:  "boolean",
				Value: "false",
			},
			{
				Key:   "conversionValue",
				Type:  "template",
				Value: "{{" + value.Name + "}}",
			},
			{
				Key:   "conversionId",
				Type:  "template",
				Value: tagConversionID,
			},
			{
				Key:   "currencyCode",
				Type:  "template",
				Value: "{{" + currency.Name + "}}",
			},
			{
				Key:   "conversionLabel",
				Type:  "template",
				Value: settings.Label,
			},
			{
				Key:   "rdp",
				Type:  "boolean",
				Value: "false",
			},
		},
		Type: "sgtmadsct",
	}
}
