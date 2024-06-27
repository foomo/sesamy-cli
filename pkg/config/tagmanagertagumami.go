package config

type TagmanagerTagUmami struct {
	Enabled     bool   `yaml:"enabled"`
	Domain      string `yaml:"domain"`
	WebsiteID   string `yaml:"website_id"`
	EndpointURL string `yaml:"endpoint_url"`
}
