package config

import (
	"github.com/foomo/gocontemplate/pkg/contemplate"
)

type GoogleAnalytics struct {
	Enabled         bool               `json:"enabled" yaml:"enabled"`
	WebContainer    contemplate.Config `json:"webContainer" yaml:"webContainer"`
	ServerContainer contemplate.Config `json:"serverContainer" yaml:"serverContainer"`
}
