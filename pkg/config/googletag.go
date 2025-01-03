package config

type GoogleTag struct {
	// A tag ID is an identifier that you put on your page to load a given Google tag
	TagID string `json:"tagId" yaml:"tagId"`
	// Enable debug mode for all user devices
	DebugMode bool `json:"debugMode" yaml:"debugMode"`
	// Whether a page_view should be sent on initial load
	SendPageView bool `json:"sendPageView" yaml:"sendPageView"`
	// TypeScript settings
	TypeScript TypeScript `json:"typeScript" yaml:"typeScript"`
}
