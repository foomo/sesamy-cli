package config

import (
	"github.com/foomo/gocontemplate/pkg/contemplate"
)

type Tracify struct {
	// Enable provider
	Enabled bool `json:"enabled" yaml:"enabled"`
	// Tracify token
	Token string `json:"token" yaml:"token"`
	// Tracify customer site id
	CustomerSiteID string `json:"customerSiteId" yaml:"customerSiteId"`
	// Google Consent settings
	GoogleConsent GoogleConsent `json:"googleConsent" yaml:"googleConsent"`
	// Google Tag Manager server container settings
	ServerContainer contemplate.Config `json:"serverContainer" yaml:"serverContainer"`
}
