package config

type GoogleAds struct {
	Enabled       bool                `json:"enabled" yaml:"enabled"`
	GoogleConsent GoogleConsent       `json:"googleConsent" yaml:"googleConsent"`
	Conversion    GoogleAdsConversion `json:"conversion" yaml:"conversion"`
}
