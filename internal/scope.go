package internal

import (
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/packages"
)

type Scope struct {
	l            *Loader
	pkg          *packages.Package
	Expr         ast.Expr
	TypeAndValue types.TypeAndValue
}

func NewScope(l *Loader, pkg *packages.Package, expr ast.Expr, tav types.TypeAndValue) *Scope {
	inst := &Scope{
		l:            l,
		pkg:          pkg,
		Expr:         expr,
		TypeAndValue: tav,
	}
	inst.astExpr(expr)

	return inst
}

func (s *Scope) astExpr(input ast.Expr) {
	switch t := input.(type) {
	case *ast.Ident:
		if t.Obj != nil {
			s.astDecl(t.Obj.Decl)
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
			s.astExpr(field.Type)
		}
	case *ast.IndexExpr:
		s.astExpr(t.X)
	case *ast.SelectorExpr:
		s.selectorExpr(t)
	}
}

func (s *Scope) selectorExpr(input *ast.SelectorExpr) {
	if x := TC[*ast.Ident](input.X); x != nil {
		if xPkgName := TC[*types.PkgName](s.pkg.TypesInfo.Uses[x]); xPkgName != nil {
			if selIdent := TC[*ast.Ident](input.Sel); selIdent != nil {
				// s.pkg.TypesInfo.
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

func (s *Scope) astDecl(input any) {
	switch t := input.(type) {
	case *ast.TypeSpec:
		s.astExpr(t.Type)
	}
}
