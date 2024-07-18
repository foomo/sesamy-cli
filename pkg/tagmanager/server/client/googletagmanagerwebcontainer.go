package client

import (
	"google.golang.org/api/tagmanager/v2"
)

func NewGoogleTagManagerWebContainer(name string, tagID string) *tagmanager.Client {
	return &tagmanager.Client{
		Name: name,
		Parameter: []*tagmanager.Parameter{
			{
				Type:  "boolean",
				Key:   "activateResponseCompression",
				Value: "true",
			},
			{
				Type:  "boolean",
				Key:   "activateGeoResolution",
				Value: "false",
			},
			{
				Type:  "boolean",
				Key:   "activateDependencyServing",
				Value: "true",
			},
			{
				Type: "list",
				Key:  "allowedContainerIds",
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
		},
		Type: "gtm_client",
	}
}
