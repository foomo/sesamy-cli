package config

import (
	"github.com/foomo/gocontemplate/pkg/contemplate"
)

type Criteo struct {
	// Enable provider
	Enabled bool `json:"enabled" yaml:"enabled"`
	// Criteo caller id
	CallerID string `json:"callerId" yaml:"callerId"`
	// Criteo partner id
	PartnerID string `json:"partnerId" yaml:"partnerId"`
	// Google Consent settings
	GoogleConsent GoogleConsent `json:"googleConsent" yaml:"googleConsent"`
	// Google Tag Manager server container settings
	ServerContainer contemplate.Config `json:"serverContainer" yaml:"serverContainer"`
}
