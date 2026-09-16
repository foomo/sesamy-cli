package config

import (
	"github.com/foomo/gocontemplate/pkg/contemplate"
)

type GoogleAnalytics struct {
	// Enable provider
	Enabled bool `json:"enabled" yaml:"enabled"`
	// Google Analytics account id
	AccountID string `json:"accountId" yaml:"accountId"`
	// Google Analytics property id
	PropertyID string `json:"propertyId" yaml:"propertyId"`
	// Google Consent settings
	GoogleConsent GoogleConsent `json:"googleConsent" yaml:"googleConsent"`
	// GTag.js override configuration
	GoogleGTagJSOverride GoogleAnalyticsGTagJSOverride `json:"googleGTagJSOverride" yaml:"googleGTagJSOverride"`
	// Enable mpv2 user data transformation (experimental)
	EnableMPv2UserDataTransformation bool `json:"enableMPv2UserDataTransformation" yaml:"enableMPv2UserDataTransformation"`
	// Custom tag template settings
	Templates GoogleAnalyticsTemplates `json:"templates" yaml:"templates"`
	// Google Tag Manager web container settings
	WebContainer contemplate.Config `json:"webContainer" yaml:"webContainer"`
	// Google Tag Manager server container settings
	ServerContainer contemplate.Config `json:"serverContainer" yaml:"serverContainer"`
}

type GoogleAnalyticsTemplates struct {
	// Path to a custom json request value variable template file, defaults to the embedded one
	JSONRequestValueVariable string `json:"jsonRequestValueVariable" yaml:"jsonRequestValueVariable"`
	// Path to a custom gtag client template file, defaults to the embedded one
	GTagClient string `json:"gtagClient" yaml:"gtagClient"`
}
