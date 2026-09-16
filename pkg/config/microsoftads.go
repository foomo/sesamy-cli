package config

type MicrosoftAds struct {
	// Enable provider
	Enabled bool `json:"enabled" yaml:"enabled"`
	// Google Consent settings
	GoogleConsent GoogleConsent `json:"googleConsent" yaml:"googleConsent"`
	// Microsoft Ads UET Tag ID
	TagID string `json:"tagId" yaml:"tagId"`
	// Microsoft Ads Conversion settings
	Conversion MicrosoftAdsConversion `json:"conversion" yaml:"conversion"`
	// Custom tag template settings
	Templates MicrosoftAdsTemplates `json:"templates" yaml:"templates"`
}

type MicrosoftAdsTemplates struct {
	// Path to a custom conversion tag template file, defaults to the embedded one
	ConversionTag string `json:"conversionTag" yaml:"conversionTag"`
}
