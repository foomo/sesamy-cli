package config

import (
	"github.com/foomo/gocontemplate/pkg/contemplate"
)

type Typescript struct {
	contemplate.Config `yaml:",squash"`
	OutputPath         string `yaml:"output_path"`
}
