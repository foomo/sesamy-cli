package client

import (
	"github.com/foomo/sesamy-cli/pkg/tagmanager/web/trigger"
	"google.golang.org/api/tagmanager/v2"
)

func NewHotjar(name string, siteID *tagmanager.Variable) *tagmanager.Tag {
	ret := &tagmanager.Tag{
		FiringTriggerId: []string{trigger.IDAllPages},
		Name:            name,
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "hotjar_site_id",
				Type:  "template",
				Value: "{{" + siteID.Name + "}}",
			},
		},
		Type: "hjtc",
	}
	return ret
}
