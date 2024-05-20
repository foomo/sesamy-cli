package typescript

import (
	"context"
	"strings"

	"github.com/foomo/sesamy-cli/internal"
)

type Builder struct {
	parser          *internal.Loader
	packages        map[string]*Package
	pkgNameReplacer *strings.Replacer
}

func NewBuilder(parser *internal.Loader) *Builder {
	inst := &Builder{
		parser:          parser,
		packages:        map[string]*Package{},
		pkgNameReplacer: strings.NewReplacer(".", "_", "/", "_", "-", "_"),
	}

	for path, pkg := range parser.Packages {
		inst.packages[path] = NewPackage(parser, pkg, inst.pkgNameReplacer)
	}

	return inst
}

// func (b *Builder) AddStruct(name string, s *types.Struct) {
// 	b.structs[name] = s
// }
//
// func (b *Builder) AddStructs(v map[string]*types.Struct) {
// 	maps.Copy(b.structs, v)
// }
//
// func (b *Builder) AddImport(pkg string, names ...string) {
// 	b.imports[pkg] = append(b.imports[pkg], names...)
// }

// type Imports map[string][]string

// func (i Imports) String() string {
// 	var ret string
// 	keys := maps.Keys(i)
// 	sort.Strings(keys)
// 	for _, name := range keys {
// 		vals := i[name]
// 		slices.Sort(vals)
// 		ret += fmt.Sprintf("import { %s } from '%s';\n", strings.Join(vals, ", "), name)
// 	}
// 	return ret
// }

func (b *Builder) Build(ctx context.Context) (map[string]*File, error) {
	ret := make(map[string]*File, len(b.packages))
	for pkgPath, pkg := range b.packages {
		if err := pkg.Build(ctx); err != nil {
			return nil, err
		}
		if len(pkg.Code().Body().Parts) > 0 {
			ret[b.pkgNameReplacer.Replace(pkgPath)+".ts"] = pkg.Code()
		}
	}
	return ret, nil
}

// func (b *Builder) String(pkgs []*internal.Package) (map[string]string, error) {
// 	ret := make(map[string]string)
// 	for _, pkg := range pkgs {
// 		s, err := b.Package(pkg)
// 		if err != nil {
// 			return nil, err
// 		}
// 		ret[b.packageNameReplacer.Replace(pkg.Path)] = s
// 	}
// 	return ret, nil
// }
