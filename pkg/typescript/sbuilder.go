package typescript

import (
	"fmt"
	"slices"
	"strings"

	"github.com/wissance/stringFormatter"
	"golang.org/x/exp/maps"
)

const (
	SectionAnnotations = 0
	SectionCopyright   = 100
	SectionHead        = 200
	SectionBody        = 300
	SectionFoot        = 400
)

type (
	File struct {
		*strings.Builder
		Indent       int
		IndentString string
		Sections     map[int]*Section
	}
	Section struct {
		Parts []string
	}
)

func NewFile() *File {
	return &File{
		Indent: 0,
		Sections: map[int]*Section{
			SectionAnnotations: {},
			SectionCopyright:   {},
			SectionHead:        {},
			SectionBody:        {},
			SectionFoot:        {},
		},
		IndentString: "\t",
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

func (f *File) Annotations() *Section {
	return f.Section(SectionAnnotations)
}

func (f *File) Copyright() *Section {
	return f.Section(SectionCopyright)
}

func (f *File) Head() *Section {
	return f.Section(SectionHead)
}

func (f *File) Body() *Section {
	return f.Section(SectionBody)
}

func (f *File) Foot() *Section {
	return f.Section(SectionFoot)
}

func (f *File) Section(id int) *Section {
	return f.Sections[id]
}

func (f *File) AddSection(id int) {
	f.Sections[id] = &Section{}
}

func (f *File) String() string {
	b := &strings.Builder{}

	sections := maps.Keys(f.Sections)
	slices.Sort(sections)

	for _, id := range sections {
		section := f.Sections[id]
		sectionParts := section.Parts
		slices.Sort(sectionParts)
		for _, part := range sectionParts {
			b.WriteString(part + "\n")
		}
	}

	return b.String()
}
