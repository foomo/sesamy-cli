package config

type GoogleTag struct {
	// A tag ID is an identifier that you put on your page to load a given Google tag
	TagID string `json:"tagId" yaml:"tagId"`
	// Whether a page_view should be sent on initial load
	SendPageView bool `json:"sendPageView" yaml:"sendPageView"`
	// Data layer variables to be added to the event settings
	DataLayerVariables map[string]string `json:"dataLayerVariables" yaml:"dataLayerVariables"`
	// TypeScript settings
	TypeScript TypeScript `json:"typeScript" yaml:"typeScript"`
}
