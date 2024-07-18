package config

type GoogleTag struct {
	TagID        string `json:"tagId" yaml:"tagId"`
	DebugMode    bool   `json:"debugMode" yaml:"debugMode"`
	SendPageView bool   `json:"sendPageView" yaml:"sendPageView"`
	// WebContainer    contemplate.Config `json:"webContainer" yaml:"webContainer"`
	// ServerContainer contemplate.Config `json:"serverContainer" yaml:"serverContainer"`
	TypeScript TypeScript `json:"typeScript" yaml:"typeScript"`
}
