package config

type Cookiebot struct {
	// Enable provider
	Enabled                      bool   `json:"enabled" yaml:"enabled"`
	TemplateName                 string `json:"templateName" yaml:"templateName"`
	CookiebotID                  string `json:"cookiebotId" yaml:"cookiebotId"`
	CDNRegion                    string `json:"cdnRegion" yaml:"cdnRegion"`
	URLPassthrough               bool   `json:"urlPassthrough" yaml:"urlPassthrough"`
	AdvertiserConsentModeEnabled bool   `json:"advertiserConsentModeEnabled" yaml:"advertiserConsentModeEnabled"`
}
