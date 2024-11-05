package config

type Config struct {
	// Config version
	Version string `json:"version" yaml:"version" jsonschema:"required"`
	// Globally redact visitor ip
	RedactVisitorIP bool `json:"redactVisitorIp" yaml:"redactVisitorIp"`
	// Google Tag settings
	GoogleTag GoogleTag `json:"googleTag" yaml:"googleTag"`
	// Google API settings
	GoogleAPI GoogleAPI `json:"googleApi" yaml:"googleApi"`
	// Google Tag Manager settings
	GoogleTagManager GoogleTagManager `json:"googleTagManager" yaml:"googleTagManager"`
	// Google Ads provider settings
	GoogleAds GoogleAds `json:"googleAds" yaml:"googleAds"`
	// CookieBot provider settings
	Cookiebot Cookiebot `json:"cookiebot" yaml:"cookiebot"`
	// Google Analytics provider settings
	GoogleAnalytics GoogleAnalytics `json:"googleAnalytics" yaml:"googleAnalytics"`
	// Conversion Linker provider settings
	ConversionLinker ConversionLinker `json:"conversionLinker" yaml:"conversionLinker"`
	// Facebook provider settings
	Facebook Facebook `json:"facebook" yaml:"facebook"`
	// Emarsys provider settings
	Emarsys Emarsys `json:"emarsys" yaml:"emarsys"`
	// Tracify provider settings
	Tracify Tracify `json:"tracify" yaml:"tracify"`
	// Umami provider settings
	Umami Umami `json:"umami" yaml:"umami"`
}
