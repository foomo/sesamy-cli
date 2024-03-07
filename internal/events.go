package internal

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"

	"github.com/fatih/structtag"
	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/pterm/pterm"
	"github.com/stoewer/go-strcase"
	"golang.org/x/tools/go/packages"
)

func GetEventParameters(cfg *config.Config) (map[string][]string, error) {
	ret := map[string][]string{}

	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedSyntax | packages.NeedFiles,
	}, cfg.Events.PackageNames()...)
	if err != nil {
		return nil, err
	}

	for i, pkg := range pkgs {
		if len(pkg.Errors) > 0 {
			return nil, fmt.Errorf("%+v", pkg.Errors)
		}

		if len(pkg.GoFiles) == 0 {
			return nil, fmt.Errorf("no input go files for package index %d", i)
		}

		conf := cfg.Events.PackageConfig(pkg.ID)

		for i, file := range pkg.Syntax {
			if conf.IsFileIgnored(pkg.GoFiles[i]) {
				continue
			}

			ast.Inspect(file, func(n ast.Node) bool {
				// GenDecl can be an import, type, var, or const expression
				if x, ok := n.(*ast.GenDecl); ok {
					if x.Tok == token.IMPORT {
						return false
					}

					for _, spec := range x.Specs {
						// e.g. "type Foo struct {}" or "type Bar = string"
						if elem, ok := spec.(*ast.TypeSpec); ok && elem.Name.IsExported() {
							if strct, ok := elem.Type.(*ast.StructType); ok {
								for _, field := range strct.Fields.List {
									tags, err := structtag.Parse(strings.Trim(field.Tag.Value, "`"))
									if err != nil {
										pterm.Warning.Println(err.Error(), field.Tag.Value)
										return false
									}
									tag, err := tags.Get("json")
									if err != nil {
										pterm.Warning.Println(err.Error())
										return false
									}
									if tag.Value() != "" && tag.Value() != "-" {
										name := strcase.SnakeCase(elem.Name.String())
										ret[name] = append(ret[name], strings.Split(tag.Value(), ",")[0])
									}
								}
							}
						}
					}
					return false
				}
				return true
			})
		}
	}

	return ret, nil
}
