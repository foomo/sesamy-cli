package config

type OpenAIAds struct {
	// Enable provider
	Enabled bool `json:"enabled" yaml:"enabled"`
	// Google Consent settings
	GoogleConsent GoogleConsent `json:"googleConsent" yaml:"googleConsent"`
	// OpenAI Ads pixel id
	PixelID string `json:"pixelId" yaml:"pixelId"`
	// OpenAI Ads api key
	APIKey string `json:"apiKey" yaml:"apiKey"`
	// OpenAI Ads Conversion settings
	Conversion OpenAIAdsConversion `json:"conversion" yaml:"conversion"`
	// Custom tag template settings
	Templates OpenAIAdsTemplates `json:"templates" yaml:"templates"`
}

type OpenAIAdsTemplates struct {
	// Path to a custom conversions api tag template file, defaults to the manually installed one
	ConversionsAPITag string `json:"conversionsApiTag" yaml:"conversionsApiTag"`
}
