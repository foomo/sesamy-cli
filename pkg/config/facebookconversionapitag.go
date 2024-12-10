package config

type FacebookConversionAPITag struct {
	// Extend Meta Pixel cookies (fbp/fbc)
	ExtendCookies bool `json:"extendCookies" yaml:"extendCookies"`
	// Enable Use of HTTP Only Secure Cookie (gtmeec) to Enhance Event Data
	EnableEventEnhancement bool `json:"enableEventEnhancement" yaml:"enableEventEnhancement"`
}
