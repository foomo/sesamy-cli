package config

type Hotjar struct {
	// Enable provider
	Enabled bool   `json:"enabled" yaml:"enabled"`
	SiteID  string `json:"siteId" yaml:"siteId"`
}
