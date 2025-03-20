package utils

import (
	"context"
	"go/types"

	"github.com/foomo/gocontemplate/pkg/assume"
	"github.com/foomo/gocontemplate/pkg/contemplate"
	"github.com/pkg/errors"
	"github.com/stoewer/go-strcase"
)

func LoadEventParams(ctx context.Context, cfg contemplate.Config) (map[string]map[string]string, error) {
	ret := map[string]map[string]string{}
	if len(cfg.Packages) == 0 {
		return ret, nil
	}

	parser, err := contemplate.Load(ctx, &cfg)
	if err != nil {
		return nil, err
	}

	for _, cfgPkg := range cfg.Packages {
		pkg := parser.Package(cfgPkg.Path)
		for _, typ := range cfgPkg.Types {
			eventParams, err := getEventParams(pkg.LookupScopeType(typ))
			if err != nil {
				return nil, errors.Wrap(err, "failed to load event params: "+cfgPkg.Path+"."+typ)
			}
			ret[strcase.SnakeCase(typ)] = eventParams
		}
	}

	return ret, nil
}

func getEventParams(obj types.Object) (map[string]string, error) {
	ret := map[string]string{}
	if obj == nil {
		return nil, errors.New("obj is nil")
	}
	if obj.Type() == nil {
		return nil, errors.New("object is not a type: " + obj.String())
	}
	if obj.Type().Underlying() == nil {
		return nil, errors.New("underlying object is not a type: " + obj.Type().String())
	}
	if eventStruct := assume.T[*types.Struct](obj.Type().Underlying()); eventStruct != nil {
		for i := range eventStruct.NumFields() {
			if eventField := eventStruct.Field(i); eventField.Name() == "Params" {
				if paramsStruct := assume.T[*types.Struct](eventField.Type().Underlying()); paramsStruct != nil {
					for j := range paramsStruct.NumFields() {
						var name string
						var value string

						tag, err := ParseStructTagName(paramsStruct.Tag(j), "json")
						if err != nil {
							return nil, errors.Wrapf(err, "failed to parse tag `%s`", paramsStruct.Tag(j))
						}
						name = tag
						value = "eventModel." + tag

						// check if there is a custom dlv tag
						if tag, err := ParseStructTagName(paramsStruct.Tag(j), "dlv"); err == nil {
							value = tag
						}
						ret[name] = value
					}
				}
			}
		}
	}
	return ret, nil
}
