package config

import (
	"github.com/foomo/gocontemplate/pkg/contemplate"
)

type GoogleAdsConversion struct {
	// Enable provider
	Enabled         bool   `json:"enabled" yaml:"enabled"`
	ConversionLabel string `json:"conversionLabel" yaml:"conversionLabel"`
	// Google Tag Manager server container settings
	ServerContainer contemplate.Config `json:"serverContainer" yaml:"serverContainer"`
}
