package config

import (
	"github.com/foomo/gocontemplate/pkg/contemplate"
)

type Facebook struct {
	// Enable provider
	Enabled        bool   `json:"enabled" yaml:"enabled"`
	PixelID        string `json:"pixelId" yaml:"pixelId"`
	APIAccessToken string `json:"apiAccessToken" yaml:"apiAccessToken"`
	TestEventToken string `json:"testEventToken" yaml:"testEventToken"`
	// Google Consent settings
	GoogleConsent GoogleConsent `json:"googleConsent" yaml:"googleConsent"`
	// Google Tag Manager server container settings
	ServerContainer contemplate.Config `json:"serverContainer" yaml:"serverContainer"`
}
