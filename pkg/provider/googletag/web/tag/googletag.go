package client

import (
	"github.com/foomo/sesamy-cli/pkg/tagmanager/web/trigger"
	"google.golang.org/api/tagmanager/v2"
)

func NewGoogleTag(name string, tagID *tagmanager.Variable, settings *tagmanager.Variable) *tagmanager.Tag {
	ret := &tagmanager.Tag{
		FiringTriggerId: []string{trigger.IDInitialization},
		Name:            name,
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "tagId",
				Type:  "template",
				Value: "{{" + tagID.Name + "}}",
			},
			{
				Key:   "configSettingsVariable",
				Type:  "template",
				Value: "{{" + settings.Name + "}}",
			},
		},
		Type: "googtag",
	}
	return ret
}
