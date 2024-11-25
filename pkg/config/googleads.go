package config

type GoogleAds struct {
	// Enable provider
	Enabled bool `json:"enabled" yaml:"enabled"`
	// Google Consent settings
	GoogleConsent GoogleConsent `json:"googleConsent" yaml:"googleConsent"`
	// Conversion id
	ConversionID string `json:"conversionId" yaml:"conversionId"`
	// Google Ads Conversion settings
	Conversion GoogleAdsConversion `json:"conversion" yaml:"conversion"`
	// Google Ads Remarketing settings
	Remarketing GoogleAdsRemarketing `json:"remarketing" yaml:"remarketing"`
}
