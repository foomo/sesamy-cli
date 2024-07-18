package utils

import (
	"go/types"

	"github.com/foomo/gocontemplate/pkg/assume"
	"github.com/foomo/gocontemplate/pkg/contemplate"
	"github.com/pkg/errors"
	"github.com/stoewer/go-strcase"
)

func LoadEventParams(cfg contemplate.Config) (map[string][]string, error) {
	parser, err := contemplate.Load(&cfg)
	if err != nil {
		return nil, err
	}

	ret := map[string][]string{}
	for _, cfgPkg := range cfg.Packages {
		pkg := parser.Package(cfgPkg.Path)
		for _, typ := range cfgPkg.Types {
			eventParams, err := getEventParams(pkg.LookupScopeType(typ))
			if err != nil {
				return nil, err
			}
			ret[strcase.SnakeCase(typ)] = eventParams
		}
	}

	return ret, nil
}

func getEventParams(obj types.Object) ([]string, error) {
	var ret []string
	if eventStruct := assume.T[*types.Struct](obj.Type().Underlying()); eventStruct != nil {
		for i := range eventStruct.NumFields() {
			if eventField := eventStruct.Field(i); eventField.Name() == "Params" {
				if paramsStruct := assume.T[*types.Struct](eventField.Type().Underlying()); paramsStruct != nil {
					for j := range paramsStruct.NumFields() {
						tag, err := ParseStructTagName(paramsStruct.Tag(j))
						if err != nil {
							return nil, errors.Wrapf(err, "failed to parse tag `%s`", paramsStruct.Tag(j))
						}
						ret = append(ret, tag)
					}
				}
			}
		}
	}
	return ret, nil
}
