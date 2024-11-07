package config

type Hotjar struct {
	Enabled bool   `json:"enabled" yaml:"enabled"`
	SiteID  string `json:"siteId" yaml:"siteId"`
}
