package config

import (
	"github.com/foomo/gocontemplate/pkg/contemplate"
)

type Tagmanager struct {
	contemplate.Config `yaml:",squash"`
	Tags               TagmanagerTags     `yaml:"tags"`
	Prefixes           TagmanagerPrefixes `yaml:"prefixes"`
}

type TagmanagerTags struct {
	GA4Enabled bool `yaml:"gA4_enabled"`
}
