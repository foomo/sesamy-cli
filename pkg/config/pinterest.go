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
	// Custom tag template settings
	Templates PinterestTemplates `json:"templates" yaml:"templates"`
	// Google Consent settings
	GoogleConsent GoogleConsent `json:"googleConsent" yaml:"googleConsent"`
	// Google Tag Manager server container settings
	ServerContainer contemplate.Config `json:"serverContainer" yaml:"serverContainer"`
}

type PinterestTemplates struct {
	// Path to a custom conversions tag template file, defaults to the manually installed one
	Tag string `json:"tag" yaml:"tag"`
}
