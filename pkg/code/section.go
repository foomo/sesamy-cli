package code

import (
	"fmt"
	"slices"

	"github.com/wissance/stringFormatter"
)

type Section struct {
	Parts []string
}

func (s *Section) Sprint(a ...any) {
	value := fmt.Sprint(a...)
	if !slices.Contains(s.Parts, value) {
		s.Parts = append(s.Parts, value)
	}
}

func (s *Section) Sprintf(format string, a ...any) {
	value := fmt.Sprintf(format, a...)
	if !slices.Contains(s.Parts, value) {
		s.Parts = append(s.Parts, value)
	}
}

// Tprintn {n} , n here is a number to notes order of argument list to use i.e. {0}, {1}
func (s *Section) Tprintn(template string, a ...any) {
	value := stringFormatter.Format(template, a...)
	if !slices.Contains(s.Parts, value) {
		s.Parts = append(s.Parts, value)
	}
}

// Tprintm {name} to notes arguments by name i.e. {name}, {last_name}, {address} and so on ...
func (s *Section) Tprintm(template string, a map[string]any) {
	value := stringFormatter.FormatComplex(template, a)
	if !slices.Contains(s.Parts, value) {
		s.Parts = append(s.Parts, value)
	}
}
