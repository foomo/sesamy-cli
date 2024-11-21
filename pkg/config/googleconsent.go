package config

type GoogleConsent struct {
	// Enable provider
	Enabled bool   `json:"enabled" yaml:"enabled"`
	Mode    string `json:"mode" yaml:"mode"`
}
