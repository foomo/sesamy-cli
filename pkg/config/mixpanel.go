package config

import (
	"github.com/foomo/gocontemplate/pkg/contemplate"
)

type (
	Mixpanel struct {
		// Enable provider
		Enabled bool `json:"enabled" yaml:"enabled"`
		// Mixpanel project token
		ProjectToken string `json:"projectToken" yaml:"projectToken"`
		// Google Consent settings
		GoogleConsent GoogleConsent `json:"googleConsent" yaml:"googleConsent"`
		// Google Tag Manager server container settings
		ServerContainer MixpanelServerContainer `json:"serverContainer" yaml:"serverContainer"`
	}
	MixpanelServerContainer struct {
		// Track events
		Track contemplate.Config `json:"track" yaml:"track"`
	}
)
