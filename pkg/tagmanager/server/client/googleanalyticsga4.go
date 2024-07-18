package client

import (
	"google.golang.org/api/tagmanager/v2"
)

func NewGoogleAnalyticsGA4(name string) *tagmanager.Client {
	return &tagmanager.Client{
		Name: name,
		Parameter: []*tagmanager.Parameter{
			{
				Type:  "boolean",
				Key:   "activateGtagSupport",
				Value: "false",
			},
			{
				Type:  "boolean",
				Key:   "activateDefaultPaths",
				Value: "true",
			},
			{
				Type:  "template",
				Key:   "cookieManagement",
				Value: "js",
			},
		},
		Type: "gaaw_client",
	}
}
