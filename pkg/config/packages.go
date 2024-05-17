package config

import (
	"github.com/pkg/errors"
)

type Packages []Package

func (c Packages) PackageNames() []string {
	ret := make([]string, len(c))
	for i, value := range c {
		ret[i] = value.Path
	}
	return ret
}

func (c Packages) PackageConfig(path string) (Package, error) {
	for _, value := range c {
		if value.Path == path {
			return value, nil
		}
	}
	return Package{}, errors.Errorf("package for path '%s' not found", path)
}
