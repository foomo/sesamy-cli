package config

import (
	"github.com/foomo/gocontemplate/pkg/contemplate"
)

type GoogleAnalytics struct {
	// Enable provider
	Enabled bool `json:"enabled" yaml:"enabled"`
	// Google Consent settings
	GoogleConsent GoogleConsent `json:"googleConsent" yaml:"googleConsent"`
	// GTag.js override configuration
	GoogleGTagJSOverride GoogleAnalyticsGTagJSOverride `json:"googleGTagJSOverride" yaml:"googleGTagJSOverride"`
	// Enable mpv2 user data transformation (experimental)
	EnableMPv2UserDataTransformation bool `json:"enableMPv2UserDataTransformation" yaml:"enableMPv2UserDataTransformation"`
	// Google Tag Manager web container settings
	WebContainer contemplate.Config `json:"webContainer" yaml:"webContainer"`
	// Google Tag Manager server container settings
	ServerContainer contemplate.Config `json:"serverContainer" yaml:"serverContainer"`
}
