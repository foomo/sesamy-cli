package config

import (
	"github.com/foomo/gocontemplate/pkg/contemplate"
)

type (
	GoogleAdsConversion struct {
		// Enable provider
		Enabled bool `json:"enabled" yaml:"enabled"`
		// Google Tag Manager server container settings
		ServerContainer GoogleAdsConversionServerContainer `json:"serverContainer" yaml:"serverContainer"`
	}
	GoogleAdsConversionServerContainer struct {
		contemplate.Config `json:",inline" yaml:",squash"`
		Settings           map[string]GoogleAdsConversionTracking `json:"settings" yaml:"settings"`
	}
)

func (s *GoogleAdsConversionServerContainer) Setting(eventName string) GoogleAdsConversionTracking {
	if value, ok := s.Settings[eventName]; ok {
		return value
	}
	return GoogleAdsConversionTracking{}
}
