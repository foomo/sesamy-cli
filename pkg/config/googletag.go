package config

type GoogleTag struct {
	TagID        string     `json:"tagId" yaml:"tagId"`
	DebugMode    bool       `json:"debugMode" yaml:"debugMode"`
	SendPageView bool       `json:"sendPageView" yaml:"sendPageView"`
	TypeScript   TypeScript `json:"typeScript" yaml:"typeScript"`
}
