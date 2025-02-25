package config

type GoogleTagManagerWebContainerVariables struct {
	// List of event data variables
	DataLayer []string `json:"dataLayer" yaml:"dataLayer"`
	// Map of lookup table variables
	LookupTables map[string]LookupTable `json:"lookupTables" yaml:"lookupTables"`
}
