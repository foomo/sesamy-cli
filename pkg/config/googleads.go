package config

type GoogleAds struct {
	Enabled       bool                 `json:"enabled" yaml:"enabled"`
	GoogleConsent GoogleConsent        `json:"googleConsent" yaml:"googleConsent"`
	ConversionID  string               `json:"conversionId" yaml:"conversionId"`
	Conversion    GoogleAdsConversion  `json:"conversion" yaml:"conversion"`
	Remarketing   GoogleAdsRemarketing `json:"remarketing" yaml:"remarketing"`
}
