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
	// Criteo application id
	ApplicationID string `json:"applicationId" yaml:"applicationId"`
	// Custom tag template settings
	Templates CriteoTemplates `json:"templates" yaml:"templates"`
	// Google Consent settings
	GoogleConsent GoogleConsent `json:"googleConsent" yaml:"googleConsent"`
	// Google Tag Manager server container settings
	ServerContainer contemplate.Config `json:"serverContainer" yaml:"serverContainer"`
}

type CriteoTemplates struct {
	// Path to a custom events api tag template file, defaults to the manually installed one
	EventsAPITag string `json:"eventsApiTag" yaml:"eventsApiTag"`
	// Path to a custom user identification template file, defaults to the manually installed one
	UserIdentification string `json:"userIdentification" yaml:"userIdentification"`
}
