package internal

import (
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/packages"
)

type Package struct {
	// Imports map[string]*Import
	l          *Loader
	pkg        *packages.Package
	exprs      map[string]ast.Expr
	types      map[string]types.Object
	scopeExprs map[string]ast.Expr
	scopeTypes map[string]types.Object
}

// ------------------------------------------------------------------------------------------------
// ~ Constructor
// ------------------------------------------------------------------------------------------------

func NewPackage(l *Loader, pkg *packages.Package) *Package {
	exprs := make(map[string]ast.Expr)
	for expr, value := range pkg.TypesInfo.Defs {
		if value != nil {
			switch value.(type) {
			case *types.Const:
				exprs[value.Name()] = expr
			case *types.Func, *types.TypeName:
				exprs[value.Name()] = expr
			}
		}
	}

	typess := make(map[string]types.Object)
	for _, name := range pkg.Types.Scope().Names() {
		typess[name] = pkg.Types.Scope().Lookup(name)
	}

	inst := &Package{
		l:          l,
		pkg:        pkg,
		types:      typess,
		exprs:      map[string]ast.Expr{},
		scopeExprs: map[string]ast.Expr{},
		scopeTypes: map[string]types.Object{},
	}

	inst.addExprs(pkg.TypesInfo.Defs)

	return inst
}

func (s *Package) Name() string {
	return s.pkg.Name
}

func (s *Package) Path() string {
	return s.pkg.PkgPath
}

func (s *Package) Exprs() map[string]ast.Expr {
	return s.exprs
}

func (s *Package) Types() map[string]types.Object {
	return s.types
}

func (s *Package) ScopeTypes() map[string]types.Object {
	return s.scopeTypes
}

func (s *Package) LookupExpr(name string) ast.Expr {
	return s.exprs[name]
}

func (s *Package) LookupScopeExpr(name string) ast.Expr {
	return s.scopeExprs[name]
}

func (s *Package) FilterExprsByTypeExpr(expr ast.Expr) []ast.Expr {
	var ret []ast.Expr
	if exprIdent := TC[*ast.Ident](expr); exprIdent != nil {
		for _, child := range s.exprs {
			if childIdent := TC[*ast.Ident](child); childIdent != nil && childIdent.Obj != nil {
				if childDecl := TC[*ast.ValueSpec](childIdent.Obj.Decl); childDecl != nil {
					if childDeclType := TC[*ast.Ident](childDecl.Type); childDeclType != nil {
						if childDeclType.Obj == exprIdent.Obj {
							ret = append(ret, child)
						}
					}
				}
			}
		}
	}
	return ret
}

// func (s *Package) GetTypeType(obj types.Object) types.Object {
// 	if objectExprIdent := TC[*ast.Ident](s.LookupExpr(obj.Name())); objectExprIdent != nil {
// 		if objectExprDecl := TC[*ast.TypeSpec](objectExprIdent.Obj.Decl); objectExprDecl != nil {
// 			if objectType := s.LookupType(objectExprDecl.Name.String()).Type(); objectType == obj.Type() {
// 				return objectType
// 			}
//
// 			return s.pkg.TypesInfo.TypeOf(objectExprDecl.Type)
// 		}
// 	}
//
// }

func (s *Package) LookupType(name string) types.Object {
	return s.types[name]
}

func (s *Package) LookupScopeType(name string) types.Object {
	return s.scopeTypes[name]
}

func (s *Package) AddScopeTypes(names ...string) {
	for _, name := range names {
		if _, ok := s.scopeTypes[name]; !ok {
			scopeType := s.LookupType(name)
			scopeExpr := s.LookupExpr(name)
			if scopeType != nil && scopeExpr != nil {
				s.scopeTypes[name] = scopeType
				s.scopeExprs[name] = scopeExpr
				s.addScopeTypeAstExpr(scopeExpr)
			}
		}
	}

	// // var added []string
	// pkgScope := s.pkg.Types.Scope()
	//
	// for _, typeName := range names {
	// 	// check if already within the local scope
	// 	if _, ok := s.Scope[typeName]; !ok {
	// 		// lookup scope object
	// 		if typeObject := pkgScope.Lookup(typeName); typeObject != nil {
	// 			// add type to local scope
	// 			for expr, tav := range s.pkg.TypesInfo.Types {
	// 				if tav.Type == typeObject.Type() {
	// 					s.Scope[typeName] = NewScope(s.l, s.pkg, expr, tav)
	// 				}
	// 			}
	//
	// 			// scan the underlying type
	// 			// s.typesType(typeObject.Type().Underlying())
	// 		}
	// 	}
	// }
}

func (s *Package) addScopeTypeAstExpr(input ast.Expr) {
	switch t := input.(type) {
	case *ast.Ident:
		if t.Obj != nil {
			s.addScopeTypeAstObject(t.Obj.Decl)
		} else {
			s.l.addPackageTypeNames(s.pkg, t.Name)
		}
		// } else if info := s.pkg.TypesInfo.TypeOf(input); info != nil {
		// 	if infoNamed := TC[*types.Named](info); infoNamed != nil {
		// 		for ident, object := range s.pkg.TypesInfo.Uses {
		// 			if infoNamed.Obj() == object {
		// 				s.astExpr(ident)
		// 			}
		// 		}
		// 	}
		// }
	case *ast.StructType:
		for _, field := range t.Fields.List {
			s.addScopeTypeAstExpr(field.Type)
		}
	case *ast.IndexExpr:
		s.addScopeTypeAstExpr(t.X)
		s.addScopeTypeAstExpr(t.Index)
	case *ast.SelectorExpr:
		s.addScoptTypeAstSelectorExpr(t)
	}
}

func (s *Package) addScoptTypeAstSelectorExpr(input *ast.SelectorExpr) {
	if x := TC[*ast.Ident](input.X); x != nil {
		if xPkgName := TC[*types.PkgName](s.pkg.TypesInfo.Uses[x]); xPkgName != nil {
			if selIdent := TC[*ast.Ident](input.Sel); selIdent != nil {
				for node, object := range s.pkg.TypesInfo.Implicits {
					if object == xPkgName {
						if nodeImportSepc := TC[*ast.ImportSpec](node); nodeImportSepc != nil {
							v := strings.Trim(nodeImportSepc.Path.Value, "\"")
							s.l.addPackageTypeNames(s.pkg.Imports[v], selIdent.Name)
						}
					}
				}
			}
		}
	}
}

func (s *Package) addScopeTypeAstObject(input any) {
	switch t := input.(type) {
	case *ast.TypeSpec:
		s.addScopeTypeAstExpr(t.Type)
	}
}

func (s *Package) LookupAstIdentDef(typeName string) *ast.Ident {
	for defAstIdent, defTypeObject := range s.pkg.TypesInfo.Defs {
		if defTypeObject != nil && defTypeObject.Name() == typeName {
			return defAstIdent
		}
	}
	return nil
}

// func (s *Package) LookupAstIdentDefsByDeclType(input *ast.Ident) map[*ast.Ident]types.Object {
// 	ret := map[*ast.Ident]types.Object{}
// 	for defAstIdent, defTypeObject := range s.pkg.TypesInfo.Defs {
// 		if defAstIdent != nil && defAstIdent.Obj != nil && defTypeObject != nil {
// 			if declValueSpec := TC[*ast.ValueSpec](defAstIdent.Obj.Decl); declValueSpec != nil {
// 				if declValueSpecIdent := TC[*ast.Ident](declValueSpec.Type); declValueSpecIdent != nil {
// 					if declValueSpecIdent.Obj == input.Obj {
// 						ret[defAstIdent] = defTypeObject
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return ret
// }

// func (s *Package) typesType(v types.Type) {
// 	switch t := v.(type) {
// 	case *types.Struct:
// 		for i := range t.NumFields() {
// 			s.typesVar(t.Field(i))
// 		}
// 	}
// }

// func (s *Package) typesVar(input *types.Var) {
// 	if !input.Exported() {
// 		return
// 	}
// 	switch t := input.Type().(type) {
// 	case *types.Named:
// 		s.addImportTypeName(t.Obj().Pkg().Path(), t.Obj().Name())
// 	default:
// 		pterm.Debug.Println("unhandled typeVar", t)
// 	}
// }

// func (s *Package) addImportTypeName(pkgPath string, typeName string) {
// 	// ignore same package path
// 	if s.pkg.PkgPath == pkgPath {
// 		return
// 	}
//
// 	if pkgImport, ok := s.pkg.Imports[pkgPath]; ok {
// 		s.l.addPackageTypeNames(pkgImport, typeName)
// 		// if _, ok := s.Imports[pkgPath]; !ok {
// 		// 	s.Imports[pkgPath] = NewImport(pkgImport)
// 		// }
// 		// s.Imports[pkgPath].AddScope(typeName)
// 	}
// }

func (s *Package) addExprs(source map[*ast.Ident]types.Object) {
	for expor, object := range source {
		if object != nil {
			switch object.(type) {
			case *types.Func:
				s.exprs[object.Name()] = expor
			case *types.Const:
				s.exprs[object.Name()] = expor
			case *types.TypeName:
				s.exprs[object.Name()] = expor
			}
		}
	}
}
