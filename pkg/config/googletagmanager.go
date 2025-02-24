package config

type GoogleTagManager struct {
	// Google Tag Manager account id
	AccountID string `json:"accountId" yaml:"accountId"`
	// Google Tag Manager web container settings
	WebContainer GoogleTagManagerContainer `json:"webContainer" yaml:"webContainer"`
	// Google Tag Manager server container settings
	ServerContainer GoogleTagManagerContainer `json:"serverContainer" yaml:"serverContainer"`
	// Google Tag Manager server container variables
	ServerContaienrVariables GoogleTagManagerServerContainerVariables `json:"serverContainerVariables" yaml:"serverContainerVariables"`
}
