package config

type Utm struct {
	// Enable provider
	Enabled bool `json:"enabled" yaml:"enabled"`
	// Custom tag template settings
	Templates UtmTemplates `json:"templates" yaml:"templates"`
	// Google Consent settings
	GoogleConsent GoogleConsent `json:"googleConsent" yaml:"googleConsent"`
}

type UtmTemplates struct {
	// Path to a custom attribution variable template file, defaults to the embedded one
	AttributionVariable string `json:"attributionVariable" yaml:"attributionVariable"`
	// Path to a custom cookie writer tag template file, defaults to the embedded one
	CookieWriterTag string `json:"cookieWriterTag" yaml:"cookieWriterTag"`
}
