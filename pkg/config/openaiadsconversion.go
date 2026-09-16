package config

import (
	"github.com/foomo/gocontemplate/pkg/contemplate"
)

type OpenAIAdsConversion struct {
	// Enable provider
	Enabled bool `json:"enabled" yaml:"enabled"`
	// Only validate the events without persisting them
	ValidateOnly bool `json:"validateOnly" yaml:"validateOnly"`
	// Google Tag Manager server container settings
	ServerContainer OpenAIAdsServerContainer `json:"serverContainer" yaml:"serverContainer"`
}

type OpenAIAdsServerContainer struct {
	contemplate.Config `json:",inline" yaml:",squash"`
	Settings           map[string]OpenAIAdsConversionTag `json:"settings" yaml:"settings"`
}

type OpenAIAdsConversionTag struct {
	// Standard event name sent to OpenAI Ads, defaults to the event name
	EventName string `json:"eventName" yaml:"eventName"`
}

func (s *OpenAIAdsServerContainer) Setting(eventName string) OpenAIAdsConversionTag {
	if value, ok := s.Settings[eventName]; ok {
		return value
	}

	return OpenAIAdsConversionTag{}
}
