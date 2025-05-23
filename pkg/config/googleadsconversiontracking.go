package config

type GoogleAdsConversionTracking struct {
	// Conversion label
	Label string `json:"label" yaml:"label"`
	// Optional conversion id overriding the default
	ConversionID string `json:"conversionId" yaml:"conversionId"`
}
