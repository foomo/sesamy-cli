package config

import (
	"github.com/foomo/gocontemplate/pkg/contemplate"
)

type (
	Facebook struct {
		// Enable provider
		Enabled        bool   `json:"enabled" yaml:"enabled"`
		PixelID        string `json:"pixelId" yaml:"pixelId"`
		APIAccessToken string `json:"apiAccessToken" yaml:"apiAccessToken"`
		TestEventToken string `json:"testEventToken" yaml:"testEventToken"`
		// Custom tag template settings
		Templates FacebookTemplates `json:"templates" yaml:"templates"`
		// Google Consent settings
		GoogleConsent GoogleConsent `json:"googleConsent" yaml:"googleConsent"`
		// Google Tag Manager server container settings
		ServerContainer FacebookServerContainer `json:"serverContainer" yaml:"serverContainer"`
	}
	FacebookServerContainer struct {
		contemplate.Config `json:",inline" yaml:",squash"`
		Settings           map[string]FacebookConversionAPITag `json:"settings" yaml:"settings"`
	}
	FacebookTemplates struct {
		// Path to a custom conversions api tag template file, defaults to the manually installed one
		ConversionsAPITag string `json:"conversionsApiTag" yaml:"conversionsApiTag"`
	}
)

func (s *FacebookServerContainer) Setting(eventName string) FacebookConversionAPITag {
	if value, ok := s.Settings[eventName]; ok {
		return value
	}

	return FacebookConversionAPITag{}
}
