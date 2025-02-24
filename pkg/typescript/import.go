package typescript

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

type Import struct {
	def      string
	location string
	// types indicates if the given module is a type
	types map[string]bool
	// modules map of map[name]alias
	modules map[string]string
}

// ------------------------------------------------------------------------------------------------
// ~ Constructor
// ------------------------------------------------------------------------------------------------

func NewImport(location string) *Import {
	return &Import{
		location: location,
		types:    make(map[string]bool),
		modules:  make(map[string]string),
	}
}

// ------------------------------------------------------------------------------------------------
// ~ Getter
// ------------------------------------------------------------------------------------------------

func (i *Import) Default() string {
	return i.def
}

func (i *Import) Location() string {
	return i.def
}

func (i *Import) SetDefault(v string) {
	i.def = v
}

func (i *Import) Modules() []string {
	ret := slices.AppendSeq(make([]string, 0, len(i.modules)), maps.Keys(i.modules))
	slices.Sort(ret)
	return ret
}

// ------------------------------------------------------------------------------------------------
// ~ Public methods
// ------------------------------------------------------------------------------------------------

func (i *Import) IsType(name string) bool {
	v, ok := i.types[name]
	return ok && v
}

func (i *Import) Alias(name string) string {
	if v, ok := i.modules[name]; ok {
		return v
	}
	return ""
}

func (i *Import) AddModule(name string) {
	i.modules[name] = ""
}

func (i *Import) AddModuleWithAlias(name, alias string) {
	i.modules[name] = alias
}

func (i *Import) AddTypeModule(name string) {
	i.types[name] = true
	i.modules[name] = ""
}

func (i *Import) AddTypeModuleWithAlias(name, alias string) {
	i.types[name] = true
	i.modules[name] = alias
}

func (i *Import) String() string {
	var defModules []string

	if i.def != "" {
		defModules = append(defModules, i.def)
	}

	if len(i.modules) > 0 {
		var modules []string
		for name, alias := range i.modules {
			if alias != "" {
				name += " as " + alias
			}
			if i.IsType(name) {
				name = "type " + name
			}
			modules = append(modules, name)
		}
		defModules = append(defModules, "{ "+strings.Join(modules, ", ")+" }")
	}

	return fmt.Sprintf("import %s from '%s';", strings.Join(defModules, ", "), i.location)
}
