package config

import (
	"github.com/gzuidhof/tygo/tygo"
)

type Tagmanager struct {
	Packages     []*tygo.PackageConfig `yaml:"packages"`
	TypeMappings map[string]string     `yaml:"type_mappings"`
	Prefixes     TagmanagerPrefixes    `yaml:"prefixes"`
}

func (e Tagmanager) PackageNames() []string {
	ret := make([]string, len(e.Packages))
	for i, value := range e.Packages {
		ret[i] = value.Path
	}
	return ret
}

func (e Tagmanager) PackageConfig(path string) *tygo.PackageConfig {
	for _, value := range e.Packages {
		if value.Path == path {
			return value
		}
	}
	return nil
}
