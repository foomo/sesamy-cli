package config

import (
	"github.com/gzuidhof/tygo/tygo"
)

type Typescript struct {
	Packages     []*tygo.PackageConfig `yaml:"packages"`
	TypeMappings map[string]string     `yaml:"type_mappings"`
}

func (e Typescript) PackageNames() []string {
	ret := make([]string, len(e.Packages))
	for i, value := range e.Packages {
		ret[i] = value.Path
	}
	return ret
}

func (e Typescript) PackageConfig(path string) *tygo.PackageConfig {
	for _, value := range e.Packages {
		if value.Path == path {
			return value
		}
	}
	return nil
}
