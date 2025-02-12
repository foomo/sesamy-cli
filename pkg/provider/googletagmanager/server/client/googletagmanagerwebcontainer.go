package client

import (
	"strconv"

	"google.golang.org/api/tagmanager/v2"
)

func NewGoogleTagManagerWebContainer(name string, tagID string, enableGeoResolution bool, visitorRegion *tagmanager.Variable) *tagmanager.Client {
	return &tagmanager.Client{
		Name: name,
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "activateResponseCompression",
				Type:  "boolean",
				Value: "true",
			},
			{
				Key:   "activateGeoResolution",
				Type:  "boolean",
				Value: strconv.FormatBool(enableGeoResolution),
			},
			{
				Key:   "activateDependencyServing",
				Type:  "boolean",
				Value: "true",
			},
			{
				Key:  "allowedContainerIds",
				Type: "list",
				List: []*tagmanager.Parameter{
					{
						Type: "map",
						Map: []*tagmanager.Parameter{
							{
								Key:   "containerId",
								Type:  "template",
								Value: tagID,
							},
						},
					},
				},
			},
			{
				Key:   "region",
				Type:  "template",
				Value: "{{" + visitorRegion.Name + "}}",
			},
		},
		Type: "gtm_client",
	}
}
