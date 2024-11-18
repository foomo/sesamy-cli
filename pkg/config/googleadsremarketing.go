package config

type GoogleAdsRemarketing struct {
	// Enable provider
	Enabled                bool `json:"enabled" yaml:"enabled"`
	EnableConversionLinker bool `json:"enableConversionLinker" yaml:"enableConversionLinker"`
}
