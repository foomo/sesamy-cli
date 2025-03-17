package config

type Cookiebot struct {
	// Enable provider
	Enabled bool `json:"enabled" yaml:"enabled"`
	// Name of the manually installed Cookiebot CMP tag template
	TemplateName string `json:"templateName" yaml:"templateName"`
	// Create an account on Cookiebot.com and copy 'Domain Group ID' from the tab 'Your Scripts' in Cookiebot
	CookiebotID string `json:"cookiebotId" yaml:"cookiebotId"`
	// Select which CDN region Cookiebot uses
	CDNRegion string `json:"cdnRegion" yaml:"cdnRegion"`
	// When using URL passthrough, a few query parameters may be appended to links as users navigate through pages on your website
	URLPassthrough bool `json:"urlPassthrough" yaml:"urlPassthrough"`
	// If enabled, Google will deduce ad_storage, ad_user_data and ad_personalization data from the TC string.
	AdvertiserConsentModeEnabled bool `json:"advertiserConsentModeEnabled" yaml:"advertiserConsentModeEnabled"`
	// Default Consent state
	RegionSettings []CookiebotRegionSetting `json:"regionSettings" yaml:"regionSettings"`
}
