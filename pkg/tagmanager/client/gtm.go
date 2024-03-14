package client

import (
	"google.golang.org/api/tagmanager/v2"
)

func NewGTM(name string, webContainerMeasurementID *tagmanager.Variable) *tagmanager.Client {
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
								Type:  "template",
								Key:   "containerId",
								Value: "{{" + webContainerMeasurementID.Name + "}}",
							},
						},
					},
				},
			},
		},
		Type: "gtm_client",
	}
}
