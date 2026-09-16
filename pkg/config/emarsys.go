package config

import (
	"github.com/foomo/gocontemplate/pkg/contemplate"
)

type Emarsys struct {
	// Enable provider
	Enabled bool `json:"enabled" yaml:"enabled"`
	// Emarsys merchant id
	MerchantID string `json:"merchantId" yaml:"merchantId"`
	// Enable test mode
	TestMode bool `json:"testMode" yaml:"testMode"`
	// Enable debug mode
	DebugMode bool `json:"debugMode" yaml:"debugMode"`
	// Custom tag template settings
	Templates EmarsysTemplates `json:"templates" yaml:"templates"`
	// Google Consent settings
	GoogleConsent GoogleConsent `json:"googleConsent" yaml:"googleConsent"`
	// Google Tag Manager web container settings
	WebContainer contemplate.Config `json:"webContainer" yaml:"webContainer"`
	// Google Tag Manager server container settings
	ServerContainer contemplate.Config `json:"serverContainer" yaml:"serverContainer"`
}

type EmarsysTemplates struct {
	// Path to a custom web extend tag template file, defaults to the embedded one
	WebExtendTag string `json:"webExtendTag" yaml:"webExtendTag"`
	// Path to a custom initialization client template file, defaults to the embedded one
	InitializationClient string `json:"initializationClient" yaml:"initializationClient"`
	// Path to a custom initialization tag template file, defaults to the embedded one
	InitializationTag string `json:"initializationTag" yaml:"initializationTag"`
}
