package config

import (
	"github.com/foomo/gocontemplate/pkg/contemplate"
)

type MicrosoftAdsConversion struct {
	// Enable provider
	Enabled bool `json:"enabled" yaml:"enabled"`
	// Google Tag Manager server container settings
	ServerContainer MicrosoftAdsServerContainer `json:"serverContainer" yaml:"serverContainer"`
}

type MicrosoftAdsServerContainer struct {
	contemplate.Config `json:",inline" yaml:",squash"`
	Settings           map[string]MicrosoftAdsConversionTag `json:"settings" yaml:"settings"`
}

type MicrosoftAdsConversionTag struct {
	PageType  string `json:"pageType" yaml:"pageType"`
	EventType string `json:"eventType" yaml:"eventType"`
}

func (s *MicrosoftAdsServerContainer) Setting(eventName string) MicrosoftAdsConversionTag {
	if value, ok := s.Settings[eventName]; ok {
		return value
	}
	return MicrosoftAdsConversionTag{}
}
