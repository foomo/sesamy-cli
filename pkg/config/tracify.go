package config

import (
	"github.com/foomo/gocontemplate/pkg/contemplate"
)

type Tracify struct {
	Enabled         bool               `json:"enabled" yaml:"enabled"`
	Token           string             `json:"token" yaml:"token"`
	CustomerSiteID  string             `json:"customerSiteId" yaml:"customerSiteId"`
	GoogleConsent   GoogleConsent      `json:"googleConsent" yaml:"googleConsent"`
	ServerContainer contemplate.Config `json:"serverContainer" yaml:"serverContainer"`
}
