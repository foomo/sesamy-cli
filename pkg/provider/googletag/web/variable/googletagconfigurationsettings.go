package variable

import (
	"maps"
	"slices"

	"google.golang.org/api/tagmanager/v2"
)

func NewGoogleTagConfigurationSettings(name string, variables map[string]string) *tagmanager.Variable {
	variableKeys := slices.AppendSeq(make([]string, 0, len(variables)), maps.Keys(variables))
	slices.Sort(variableKeys)

	list := make([]*tagmanager.Parameter, len(variables))
	for i, k := range variableKeys {
		list[i] = &tagmanager.Parameter{
			Type: "map",
			Map: []*tagmanager.Parameter{
				{
					Key:   "parameter",
					Type:  "template",
					Value: k,
				},
				{
					Key:   "parameterValue",
					Type:  "template",
					Value: variables[k],
				},
			},
		}
	}

	return &tagmanager.Variable{
		Name: name,
		Parameter: []*tagmanager.Parameter{
			{
				Key:  "configSettingsTable",
				Type: "list",
				List: list,
			},
		},
		Type: "gtcs",
	}
}
