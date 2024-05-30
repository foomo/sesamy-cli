package typescript

import (
	"github.com/foomo/sesamy-cli/pkg/code"
)

type Imports struct {
	imports map[string]*Import
}

// ------------------------------------------------------------------------------------------------
// ~ Constructor
// ------------------------------------------------------------------------------------------------

func NewImports() *Imports {
	return &Imports{
		imports: map[string]*Import{},
	}
}

// ------------------------------------------------------------------------------------------------
// ~ Public methods
// ------------------------------------------------------------------------------------------------

func (i *Imports) Import(location string) *Import {
	if _, ok := i.imports[location]; !ok {
		i.imports[location] = NewImport(location)
	}
	return i.imports[location]
}

func (i *Imports) Write(f *code.Section) {
	for key := range i.imports {
		f.Sprintf(i.imports[key].String())
	}
}
