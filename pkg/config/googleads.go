package config

type GoogleAds struct {
	Enabled    bool                `json:"enabled" yaml:"enabled"`
	Conversion GoogleAdsConversion `json:"conversion" yaml:"conversion"`
}
