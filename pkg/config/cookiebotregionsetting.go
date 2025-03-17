package config

type CookiebotRegionSetting struct {
	// Region (leave blank to apply globally)
	Region string `json:"region" yaml:"region"`
	// Default consent for functionality_storage and personalization_storage
	Preferences string `json:"preferences" yaml:"preferences"`
	// Default consent for analytics_storage
	Statistics string `json:"statistics" yaml:"statistics"`
	// Default consent for ad_storage
	Marketing string `json:"marketing" yaml:"marketing"`
	// Default consent ad_user_data
	AdUserData string `json:"adUserData" yaml:"adUserData"`
	// Default consent ad_personalization
	AdPersonalization string `json:"adPersonalization" yaml:"adPersonalization"`
}
