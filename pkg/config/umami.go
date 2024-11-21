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
	// Google Consent settings
	GoogleConsent GoogleConsent `json:"googleConsent" yaml:"googleConsent"`
	// Google Tag Manager server container settings
	ServerContainer contemplate.Config `json:"serverContainer" yaml:"serverContainer"`
}
