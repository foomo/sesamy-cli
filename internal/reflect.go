package internal

import (
	"context"
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"maps"
	"slices"

	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/pkg/errors"
	"github.com/pterm/pterm"
	"github.com/stoewer/go-strcase"
	"golang.org/x/tools/go/packages"
)

func GetStructTypes(ctx context.Context, cfg config.Packages) (map[string]*types.Struct, error) {
	ret := map[string]*types.Struct{}

	fset := token.NewFileSet()
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName | packages.NeedTypesInfo |
			packages.NeedFiles | packages.NeedImports | packages.NeedDeps |
			packages.NeedModule | packages.NeedTypes | packages.NeedSyntax,
		Context: ctx,
		Logf: func(format string, args ...any) {
			pterm.Debug.Printfln(format, args...)
		},
		Fset: fset,
	}, cfg.PackageNames()...)
	if err != nil {
		return nil, err
	}

	for _, pkg := range pkgs {
		if len(pkg.Errors) > 0 {
			return nil, errors.Wrap(pkg.Errors[0], "packages contain errors")
		}
		packageStructs, err := getPackageStructs(cfg, pkg)
		if err != nil {
			return nil, err
		}
		maps.Copy(ret, packageStructs)
	}

	return ret, nil
}

func GetStructTypeParameters(ctx context.Context, cfg config.Packages) (map[string][]string, error) {
	typs, err := GetStructTypes(ctx, cfg)
	if err != nil {
		return nil, err
	}

	ret := map[string][]string{}
	for name, t := range typs {
		var fields []string
		for i := range t.NumFields() {
			tag, err := ParseTag(t.Tag(i))
			if err != nil {
				return nil, err
			}
			if tag != "" {
				fields = append(fields, tag)
			}
		}
		ret[strcase.SnakeCase(name)] = fields
	}
	return ret, nil
}

type Struct struct {
	Name       string      `json:"name,omitempty"`
	Attributes []Attribute `json:"attributes,omitempty"`
}

type Attribute struct {
	Name string            `json:"name,omitempty"`
	Tags map[string]string `json:"tags,omitempty"`
}

type (
	Scanner struct {
		cfg      *Config
		pkgs     []*packages.Package
		Packages map[string]*ScannerPackage
	}
	ScannerPackage struct {
		pkg     *packages.Package
		Name    string
		PkgPath string
		Imports map[*types.Package][]string
		Scope   ScannerScope
		Values  ScannerValues
	}
	ScannerScope map[string]types.Type
)

func (s ScannerScope) LookupUnderlyingTypeName(name string) map[string]types.Type {
	ret := map[string]types.Type{}
	for i, i2 := range s {
		if x, ok := i2.(*types.Named); ok && i != name && x.Obj().Name() == name {
			ret[i] = i2
		}
	}
	return ret
}

type (
	Config struct {
		Packages []*ConfigPackage
	}
	ConfigPackage struct {
		Path  string
		Names []string
	}
)

func (c *Config) PackageNames(path string) []string {
	for _, configPackage := range c.Packages {
		if configPackage.Path == path {
			return configPackage.Names
		}
	}
	return nil
}

func (c *Config) PackagePaths() []string {
	ret := make([]string, len(c.Packages))
	for i, p := range c.Packages {
		ret[i] = p.Path
	}
	return ret
}

func NewScanner(cfg *Config) *Scanner {
	return &Scanner{
		cfg:      cfg,
		Packages: map[string]*ScannerPackage{},
	}
}

func (s *Scanner) Scan(ctx context.Context) error {
	var err error
	// load packages
	s.pkgs, err = packages.Load(&packages.Config{
		Mode: packages.NeedName | packages.NeedTypesInfo |
			packages.NeedFiles | packages.NeedImports | packages.NeedDeps |
			packages.NeedModule | packages.NeedTypes | packages.NeedSyntax,
		Context: ctx,
		// Fset: token.NewFileSet(),
	}, s.cfg.PackagePaths()...)
	if err != nil {
		return err
	}

	// iterate packages
	for _, pkg := range s.pkgs {
		// retrieve requested packages names
		if names := s.cfg.PackageNames(pkg.PkgPath); len(names) > 0 {
			s.pkgPackage(pkg, names...)
		}
	}

	return nil
}

// ------------------------------------------------------------------------------------------------
// ~ Private methods
// ------------------------------------------------------------------------------------------------

func (s *Scanner) pkgPackage(pkg *packages.Package, names ...string) {
	if _, ok := s.Packages[pkg.PkgPath]; !ok {
		s.Packages[pkg.PkgPath] = NewScannerPackage(pkg)
	}
	// add request scopes
	added := s.Packages[pkg.PkgPath].typesScope(names)

	// check underlying added scopes
	for _, name := range added {
		s.typesType(pkg, s.Packages[pkg.PkgPath].Scope[name].Underlying())
	}
}

func (s *Scanner) typesType(pkg *packages.Package, v types.Type) {
	switch t := v.(type) {
	case *types.Struct:
		// iterate fields
		for i := range t.NumFields() {
			s.typesVar(pkg, t.Field(i))
		}
	default:
		fmt.Println(t)
	}
}

func (s *Scanner) typesVar(pkg *packages.Package, v *types.Var) {
	if !v.Exported() {
		return
	}
	switch t := v.Type().(type) {
	case *types.Named:
		if p, ok := pkg.Imports[v.Pkg().Path()]; ok {
			s.pkgPackage(p, t.Obj().Name())
		}
	}
}

type ScannerValues map[string]*ast.ValueSpec

func (s ScannerValues) Lookup(name string) *ast.ValueSpec {
	return s[name]
}

func NewScannerPackage(pkg *packages.Package) *ScannerPackage {
	values := ScannerValues{}
	for _, file := range pkg.Syntax {
		for _, decl := range file.Decls {
			if genDecl, ok := decl.(*ast.GenDecl); ok {
				for _, spec := range genDecl.Specs {
					if valueSpec, ok := spec.(*ast.ValueSpec); ok && len(valueSpec.Names) > 0 {
						values[valueSpec.Names[0].Name] = valueSpec
					}
				}
			}
		}
	}

	return &ScannerPackage{
		pkg:     pkg,
		Name:    pkg.Name,
		PkgPath: pkg.PkgPath,
		Imports: map[*types.Package][]string{},
		Scope:   ScannerScope{},
		Values:  values,
	}
}

func (s *ScannerPackage) typesScope(names []string) []string {
	var added []string
	scope := s.pkg.Types.Scope()

	// iterate requested names
	for _, name := range names {
		// check if already within the local scope
		if _, ok := s.Scope[name]; !ok {
			// lookup scope object
			if obj := scope.Lookup(name); obj != nil {
				// add type to local scope
				s.Scope[name] = obj.Type()
				// scan the underlying type
				s.typesType(obj.Type().Underlying())
				// append to added scopes
				added = append(added, name)
			}

			// FIXME find underlying type usages e.g. `const Foo <name>`
			for _, i := range scope.Names() {
				child := scope.Lookup(i)
				if x, ok := child.Type().(*types.Named); ok && x.Obj().Name() == name {
					s.Scope[i] = x
					added = append(added, i)
				}
			}
		}
	}
	return added
}

func (s *ScannerPackage) typesType(v types.Type) {
	switch t := v.(type) {
	case *types.Struct:
		for i := range t.NumFields() {
			s.typesVar(t.Field(i))
		}
	}
}

func (s *ScannerPackage) typesVar(v *types.Var) {
	if v.Exported() {
		if p, ok := v.Type().(*types.Named); ok {
			pkg := p.Obj().Pkg().Path()
			if s.pkg.PkgPath != pkg {
				if tv, ok := v.Type().(*types.Named); ok {
					typeName := tv.Obj().Name()
					if !slices.Contains(s.Imports[v.Pkg()], typeName) {
						s.Imports[v.Pkg()] = append(s.Imports[v.Pkg()], typeName)
					}
				}
			}
		} else {
			pterm.Debug.Println("unhandled typeVar")
		}
	}
}

// --------------------

type Package struct {
	pkg   *packages.Package
	Types []*Type `json:"s,omitempty"`
}

type Type struct {
	Spec     *ast.TypeSpec
	TypeInfo types.TypeAndValue
}

func NewPackage(pkg *packages.Package) *Package {
	inst := &Package{
		pkg: pkg,
	}
	for _, file := range pkg.Syntax {
		ast.Inspect(file, inst.astNode)
	}
	return inst
}

func (s *Package) Name() string {
	return s.pkg.Name
}

func (s *Package) Path() string {
	return s.pkg.PkgPath
}

// ------------------------------------------------------------------------------------------------
// ~ Private methods
// ------------------------------------------------------------------------------------------------

func (s *Package) astNode(n ast.Node) bool {
	switch t := n.(type) {
	case *ast.File:
		for _, d := range t.Decls {
			s.astDecl(d)
		}
	}
	return false
}

func (s *Package) astDecl(v ast.Decl) {
	switch t := v.(type) {
	case *ast.GenDecl:
		s.astGenDecl(t)
	}
}

func (s *Package) astGenDecl(v *ast.GenDecl) {
	switch v.Tok {
	case token.IMPORT:
		return
	default:
		for _, spec := range v.Specs {
			s.astSpec(spec)
		}
	}
}

func (s *Package) astSpec(v ast.Spec) {
	switch t := v.(type) {
	case *ast.TypeSpec:
		if t.Name.IsExported() {
			s.astTypeSpec(t)
		}
	}
}

func (s *Package) astTypeSpec(v *ast.TypeSpec) {
	r := &Type{Spec: v}
	ast.Inspect(v, s.astNode)
	switch t := v.Type.(type) {
	case *ast.IndexExpr:
		if value, ok := s.typeInfo(t.Index); ok {
			r.TypeInfo = value
		}
	default:
		return
	}
	s.Types = append(s.Types, r)
}

func (s *Package) typeInfo(n ast.Expr) (types.TypeAndValue, bool) {
	v, ok := s.pkg.TypesInfo.Types[n]
	return v, ok
}

func getPackageStructs(cfg config.Packages, pkg *packages.Package) (map[string]*types.Struct, error) {
	ret := map[string]*types.Struct{}
	if len(pkg.GoFiles) == 0 {
		return nil, fmt.Errorf("no input go files for package index")
	}

	pkgCfg, err := cfg.PackageConfig(pkg.ID)
	if err != nil {
		return nil, err
	}

	for _, file := range pkg.Syntax {
		ast.Inspect(file, func(n ast.Node) bool {
			// GenDecl can be an import, type, var, or const expression
			if genDecl, ok := n.(*ast.GenDecl); ok {
				if genDecl.Tok == token.IMPORT {
					return false
				}
				for _, spec := range genDecl.Specs {
					// e.g. "type Foo struct {}" or "type Bar = string"
					if typeSpec, ok := spec.(*ast.TypeSpec); ok && typeSpec.Name.IsExported() {
						if t, ok := typeSpec.Type.(*ast.IndexExpr); ok {
							x := pkg.TypesInfo.Types[t]
							fmt.Println(x)
						}
						if !pkgCfg.ExportEvent(typeSpec.Name.String()) {
							continue
						}
						if indexExpr, ok := typeSpec.Type.(*ast.IndexExpr); ok {
							if indexType, ok := pkg.TypesInfo.Types[indexExpr.Index]; ok {
								if s, ok := indexType.Type.Underlying().(*types.Struct); ok {
									ret[typeSpec.Name.String()] = s
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

	return ret, nil
}
