package config

type GoogleTagManagerServerContainerVariables struct {
	// List of event data variables
	EventData []string `json:"eventData" yaml:"eventData"`
	// Map of lookup table variables
	LookupTables map[string]LookupTable `json:"lookupTables" yaml:"lookupTables"`
}
