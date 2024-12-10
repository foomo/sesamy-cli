package transformation

import (
	"google.golang.org/api/tagmanager/v2"
)

func NewMPv2UserData(name string, variable *tagmanager.Variable, client *tagmanager.Client) *tagmanager.Transformation {
	return &tagmanager.Transformation{
		Name: name,
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "matchingConditionsEnabled",
				Type:  "boolean",
				Value: "true",
			},
			{
				Key:   "allTagsExcept",
				Type:  "boolean",
				Value: "true",
			},
			{
				Key:  "booleanExpressionString",
				Type: "template",
			},
			{
				Key: "augmentEventTable",
				List: []*tagmanager.Parameter{
					{
						IsWeakReference: false,
						Map: []*tagmanager.Parameter{
							{
								Key:   "paramName",
								Type:  "template",
								Value: "user_data",
							},
							{
								Key:   "paramValue",
								Type:  "template",
								Value: "{{" + variable.Name + "}}",
							},
						},
						Type: "map",
					},
				},
				Type: "list",
			},
			{
				Key:  "affectedTags",
				Type: "list",
			},
			{
				Key:  "affectedTagTypes",
				Type: "list",
			},
			{
				Key: "matchingConditionsTable",
				List: []*tagmanager.Parameter{
					{
						Map: []*tagmanager.Parameter{
							{
								Key:   "variableName",
								Type:  "template",
								Value: "Client Name",
							},
							{
								Key:   "variableReference",
								Type:  "template",
								Value: "{{Client Name}}",
							},
							{
								Key:   "expressionType",
								Type:  "template",
								Value: "EQUALS",
							},
							{
								Key:   "expressionValue",
								Type:  "template",
								Value: client.Name,
							},
						},
						Type: "map",
					},
				},
				Type: "list",
			},
		},
		Type: "tf_augment_event",
	}
}
