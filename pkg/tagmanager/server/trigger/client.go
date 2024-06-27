package trigger

import (
	"google.golang.org/api/tagmanager/v2"
)

func NewClient(name string, client *tagmanager.Client) *tagmanager.Trigger {
	return &tagmanager.Trigger{
		Type: "always",
		Name: name,
		Filter: []*tagmanager.Condition{
			{
				Parameter: []*tagmanager.Parameter{
					{
						Key:   "arg0",
						Type:  "template",
						Value: "{{Client Name}}",
					},
					{
						Key:   "arg1",
						Type:  "template",
						Value: client.Name,
					},
				},
				Type: "equals",
			},
		},
	}
}
