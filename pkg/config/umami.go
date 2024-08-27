package config

import (
	"github.com/foomo/gocontemplate/pkg/contemplate"
)

type Umami struct {
	Enabled         bool               `json:"enabled" yaml:"enabled"`
	Domain          string             `json:"domain" yaml:"domain"`
	WebsiteID       string             `json:"websiteId" yaml:"websiteId"`
	EndpointURL     string             `json:"endpointUrl" yaml:"endpointUrl"`
	GoogleConsent   GoogleConsent      `json:"googleConsent" yaml:"googleConsent"`
	ServerContainer contemplate.Config `json:"serverContainer" yaml:"serverContainer"`
}
