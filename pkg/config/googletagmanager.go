package config

type GoogleTagManager struct {
	AccountID       string                    `json:"accountId" yaml:"accountId"`
	WebContainer    GoogleTagManagerContainer `json:"webContainer" yaml:"webContainer"`
	ServerContainer GoogleTagManagerContainer `json:"serverContainer" yaml:"serverContainer"`
}
