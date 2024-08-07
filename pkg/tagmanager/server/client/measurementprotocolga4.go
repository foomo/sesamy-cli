package client

import (
	"google.golang.org/api/tagmanager/v2"
)

func NewMeasurementProtocolGA4(name string) *tagmanager.Client {
	return &tagmanager.Client{
		Name: name,
		Parameter: []*tagmanager.Parameter{
			{
				Type:  "template",
				Key:   "activationPath",
				Value: "/mp/collect",
			},
		},
		Type: "mpaw_client",
	}
}
