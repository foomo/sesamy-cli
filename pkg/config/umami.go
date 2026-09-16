package config

import (
	"github.com/foomo/gocontemplate/pkg/contemplate"
)

type Umami struct {
	// Enable provider
	Enabled     bool   `json:"enabled" yaml:"enabled"`
	Domain      string `json:"domain" yaml:"domain"`
	WebsiteID   string `json:"websiteId" yaml:"websiteId"`
	EndpointURL string `json:"endpointUrl" yaml:"endpointUrl"`
	// Custom tag template settings
	Templates UmamiTemplates `json:"templates" yaml:"templates"`
	// Google Consent settings
	GoogleConsent GoogleConsent `json:"googleConsent" yaml:"googleConsent"`
	// Google Tag Manager server container settings
	ServerContainer contemplate.Config `json:"serverContainer" yaml:"serverContainer"`
}

type UmamiTemplates struct {
	// Path to a custom tag template file, defaults to the embedded one
	Tag string `json:"tag" yaml:"tag"`
}
