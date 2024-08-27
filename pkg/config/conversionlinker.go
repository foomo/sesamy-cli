package config

type ConversionLinker struct {
	Enabled       bool          `json:"enabled" yaml:"enabled"`
	GoogleConsent GoogleConsent `json:"googleConsent" yaml:"googleConsent"`
}
