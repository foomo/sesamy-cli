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
	// Enable stating mode
	StagingModeEnabled bool `json:"stagingModeEnabled" yaml:"stagingModeEnabled"`
	// Custom tag template settings
	Templates TracifyTemplates `json:"templates" yaml:"templates"`
	// Google Consent settings
	GoogleConsent GoogleConsent `json:"googleConsent" yaml:"googleConsent"`
	// Google Tag Manager server container settings
	ServerContainer contemplate.Config `json:"serverContainer" yaml:"serverContainer"`
}

type TracifyTemplates struct {
	// Path to a custom tag template file, defaults to the embedded one
	Tag string `json:"tag" yaml:"tag"`
}
