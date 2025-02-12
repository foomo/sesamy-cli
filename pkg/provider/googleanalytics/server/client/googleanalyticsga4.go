package client

import (
	"strconv"

	"google.golang.org/api/tagmanager/v2"
)

func NewGoogleAnalyticsGA4(name string, enableGeoResolution bool, visitorRegion, measurementID *tagmanager.Variable) *tagmanager.Client {
	return &tagmanager.Client{
		Name: name,
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "activateResponseCompression",
				Type:  "template",
				Value: "true",
			},
			{
				Key:   "activateGeoResolution",
				Type:  "boolean",
				Value: strconv.FormatBool(enableGeoResolution),
			},
			{
				Key:   "activateGtagSupport",
				Type:  "boolean",
				Value: "true",
			},
			{
				Key:   "activateDependencyServing",
				Type:  "boolean",
				Value: "true",
			},
			{
				Key:   "activateDefaultPaths",
				Type:  "boolean",
				Value: "true",
			},
			{
				Key:   "region",
				Type:  "template",
				Value: "{{" + visitorRegion.Name + "}}",
			},
			{
				Key:   "cookieManagement",
				Type:  "template",
				Value: "js",
			},
			{
				Key:  "gtagMeasurementIds",
				Type: "list",
				List: []*tagmanager.Parameter{
					{
						Type: "map",
						Map: []*tagmanager.Parameter{
							{
								Key:   "measurementId",
								Type:  "template",
								Value: "{{" + measurementID.Name + "}}",
							},
						},
					},
				},
			},
		},
		Type: "gaaw_client",
	}
}
