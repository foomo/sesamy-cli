package config

type Config struct {
	Version          string           `json:"version" yaml:"version"`
	RedactVisitorIP  bool             `json:"redactVisitorIp" yaml:"redactVisitorIp"`
	GoogleTag        GoogleTag        `json:"googleTag" yaml:"googleTag"`
	GoogleAPI        GoogleAPI        `json:"googleAPI" yaml:"googleAPI"`
	GoogleTagManager GoogleTagManager `json:"googleTagManager" yaml:"googleTagManager"`
	// Providers
	GoogleAnalytics  GoogleAnalytics  `json:"googleAnalytics" yaml:"googleAnalytics"`
	ConversionLinker ConversionLinker `json:"conversionLinker" yaml:"conversionLinker"`
	Umami            Umami            `json:"umami" yaml:"umami"`
}
