package config

type GoogleTagManager struct {
	AccountID    string                    `json:"accountId" yaml:"accountId"`
	WebContainer GoogleTagManagerContainer `json:"webContainer" yaml:"webContainer"`
	// Google Tag Manager server container settings
	ServerContainer GoogleTagManagerContainer `json:"serverContainer" yaml:"serverContainer"`
}
