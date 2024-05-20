package config

import (
	"github.com/foomo/sesamy-cli/internal"
)

type Typescript struct {
	internal.LoaderConfig `yaml:",squash"`
	OutputPath            string `yaml:"output_path"`
}
