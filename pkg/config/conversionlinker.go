package config

type ConversionLinker struct {
	// Enable provider
	Enabled bool `json:"enabled" yaml:"enabled"`
	// Google Consent settings
	GoogleConsent GoogleConsent `json:"googleConsent" yaml:"googleConsent"`
}
