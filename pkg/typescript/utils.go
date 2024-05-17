package typescript

// func TSType(value types.Type) string {
// 	switch t := value.(type) {
// 	case *types.Basic:
// 		return tsType(t.Info())
// 	case *types.Named:
// 		return TSType(t.Obj().Type().Underlying())
// 	case *types.Slice:
// 		var t string
// 		if s, ok := t.Elem().(*types.Named); ok {
// 			s.Obj()
// 		}
// 		return "Array<" + TSType(t.Elem().Underlying()) + ">"
// 	default:
// 		return "any"
// 	}
// }
//
// func tsType(c types.BasicInfo) string {
// 	switch {
// 	case c&types.IsNumeric != 0:
// 		return "number"
// 	case c&types.IsString != 0:
// 		return "string"
// 	case c&types.IsBoolean != 0:
// 		return "boolean"
// 	default:
// 		return "any"
// 	}
// }
