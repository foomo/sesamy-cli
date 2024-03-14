package variable

import (
	"sort"

	"google.golang.org/api/tagmanager/v2"
)

func NewGTEventSettings(name string, variables map[string]*tagmanager.Variable) *tagmanager.Variable {
	parameters := make([]string, 0, len(variables))
	for k := range variables {
		parameters = append(parameters, k)
	}
	sort.Strings(parameters)

	list := make([]*tagmanager.Parameter, len(parameters))
	for i, parameter := range parameters {
		list[i] = &tagmanager.Parameter{
			Type: "map",
			Map: []*tagmanager.Parameter{
				{
					Key:   "parameter",
					Type:  "template",
					Value: parameter,
				},
				{
					Key:   "parameterValue",
					Type:  "template",
					Value: "{{" + variables[parameter].Name + "}}",
				},
			},
		}
	}

	return &tagmanager.Variable{
		Name: name,
		Parameter: []*tagmanager.Parameter{
			{
				Key:  "eventSettingsTable",
				Type: "list",
				List: list,
			},
		},
		Type: "gtes",
	}
}
