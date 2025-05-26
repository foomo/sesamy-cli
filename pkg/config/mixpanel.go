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
		// Set events
		Set contemplate.Config `json:"set" yaml:"set"`
		// SetOnce events
		SetOnce contemplate.Config `json:"setOnce" yaml:"setOnce"`
		// Reset events
		Reset contemplate.Config `json:"reset" yaml:"reset"`
		// Track events
		Track contemplate.Config `json:"track" yaml:"track"`
		// Identify events
		Identify contemplate.Config `json:"identify" yaml:"identify"`
	}
)
