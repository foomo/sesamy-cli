package config

import (
	"github.com/foomo/gocontemplate/pkg/contemplate"
)

type GoogleAnalytics struct {
	// Enable provider
	Enabled    bool       `json:"enabled" yaml:"enabled"`
	GoogleGTag GoogleGTag `json:"googleGTag" yaml:"googleGTag"`
	// Google Consent settings
	GoogleConsent GoogleConsent      `json:"googleConsent" yaml:"googleConsent"`
	WebContainer  contemplate.Config `json:"webContainer" yaml:"webContainer"`
	// Google Tag Manager server container settings
	ServerContainer contemplate.Config `json:"serverContainer" yaml:"serverContainer"`
}
