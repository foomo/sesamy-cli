package variable

import (
	"maps"
	"slices"

	"github.com/foomo/sesamy-cli/pkg/config"
	"google.golang.org/api/tagmanager/v2"
)

func LookupTableName(v string) string {
	return "lookup_table." + v
}

func NewLookupTable(name string, data config.LookupTable) *tagmanager.Variable {
	var list []*tagmanager.Parameter
	{
		keys := slices.AppendSeq(make([]string, 0, len(data.KeyTable)), maps.Keys(data.KeyTable))
		slices.Sort(keys)
		for _, k := range keys {
			v := data.KeyTable[k]
			list = append(list, &tagmanager.Parameter{
				Type: "map",
				Map: []*tagmanager.Parameter{
					{
						Key:   "key",
						Type:  "template",
						Value: k,
					},
					{
						Key:   "value",
						Type:  "template",
						Value: v,
					},
				},
			})
		}
	}
	keys := slices.AppendSeq(make([]string, 0, len(data.ValueTable)), maps.Keys(data.ValueTable))
	slices.Sort(keys)
	for _, k := range keys {
		v := data.ValueTable[k]
		list = append(list, &tagmanager.Parameter{
			Type: "map",
			Map: []*tagmanager.Parameter{
				{
					Key:   "key",
					Type:  "template",
					Value: v,
				},
				{
					Key:   "value",
					Type:  "template",
					Value: k,
				},
			},
		})
	}
	return &tagmanager.Variable{
		Name: LookupTableName(name),
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "setDefaultValue",
				Type:  "boolean",
				Value: "false",
			},
			{
				Key:   "input",
				Type:  "template",
				Value: data.Input,
			},
			{
				Key:  "map",
				Type: "list",
				List: list,
			},
		},
		Type: "smm",
	}
}
