package code

import (
	"maps"
	"slices"
	"strings"
)

const (
	SectionAnnotations = 0
	SectionCopyright   = 100
	SectionHead        = 200
	SectionBody        = 300
	SectionFoot        = 400
)

type File struct {
	Sections map[int]*Section
}

func NewFile() *File {
	return &File{
		Sections: map[int]*Section{
			SectionAnnotations: {},
			SectionCopyright:   {},
			SectionHead:        {},
			SectionBody:        {},
			SectionFoot:        {},
		},
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

	sections := slices.AppendSeq(make([]int, 0, len(f.Sections)), maps.Keys(f.Sections))
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
