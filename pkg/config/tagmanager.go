package config

import (
	"github.com/foomo/sesamy-cli/internal"
)

type Tagmanager struct {
	internal.LoaderConfig `yaml:",squash"`
	Tags                  TagmanagerTags     `yaml:"tags"`
	Prefixes              TagmanagerPrefixes `yaml:"prefixes"`
}

type TagmanagerTags struct {
	GA4Enabled bool `yaml:"gA4_enabled"`
}
