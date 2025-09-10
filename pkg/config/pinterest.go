package config

import (
	"github.com/foomo/gocontemplate/pkg/contemplate"
)

type Pinterest struct {
	// Enable provider
	Enabled bool `json:"enabled" yaml:"enabled"`
	// Pinterest advertiser id
	AdvertiserID string `json:"advertiserId" yaml:"advertiserId"`
	// Pinterest API access token
	APIAccessToken string `json:"apiAccessToken" yaml:"apiAccessToken"`
	// Enable test mode
	TestModeEnabled bool `json:"testModeEnabled" yaml:"testModeEnabled"`
	// Google Consent settings
	GoogleConsent GoogleConsent `json:"googleConsent" yaml:"googleConsent"`
	// Google Tag Manager server container settings
	ServerContainer contemplate.Config `json:"serverContainer" yaml:"serverContainer"`
}
