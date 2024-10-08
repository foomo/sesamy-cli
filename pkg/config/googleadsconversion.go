package config

import (
	"github.com/foomo/gocontemplate/pkg/contemplate"
)

type GoogleAdsConversion struct {
	Enabled         bool               `json:"enabled" yaml:"enabled"`
	ConversionLabel string             `json:"conversionLabel" yaml:"conversionLabel"`
	ServerContainer contemplate.Config `json:"serverContainer" yaml:"serverContainer"`
}
