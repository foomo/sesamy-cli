package config

import (
	"github.com/foomo/gocontemplate/pkg/contemplate"
)

type TypeScript struct {
	contemplate.Config `json:",inline" yaml:",squash"`
	OutputPath         string `json:"outputPath" yaml:"outputPath"`
}
