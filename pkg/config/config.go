package config

type Config struct {
	// Config version
	Version string `json:"version" yaml:"version" jsonschema:"required"`
	// Globally redact visitor ip
	RedactVisitorIP bool `json:"redactVisitorIp" yaml:"redactVisitorIp"`
	// Enable region specific settings
	// https://developers.google.com/tag-platform/tag-manager/server-side/enable-region-specific-settings
	EnableGeoResolution bool `json:"enableGeoResolution" yaml:"enableGeoResolution"`
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
	// MicrosoftAds provider settings
	MicrosoftAds MicrosoftAds `json:"microsoftAds" yaml:"microsoftAds"`
	// Mixpanel provider settings
	Mixpanel Mixpanel `json:"mixpanel" yaml:"mixpanel"`
	// Emarsys provider settings
	Emarsys Emarsys `json:"emarsys" yaml:"emarsys"`
	// Hotjar provider settings
	Hotjar Hotjar `json:"hotjar" yaml:"hotjar"`
	// Criteo provider settings
	Criteo Criteo `json:"criteo" yaml:"criteo"`
	// Tracify provider settings
	Tracify Tracify `json:"tracify" yaml:"tracify"`
	// Umami provider settings
	Umami Umami `json:"umami" yaml:"umami"`
}
