package config

import (
	"github.com/foomo/gocontemplate/pkg/contemplate"
)

type Facebook struct {
	Enabled         bool               `json:"enabled" yaml:"enabled"`
	PixelID         string             `json:"pixelId" yaml:"pixelId"`
	APIAccessToken  string             `json:"apiAccessToken" yaml:"apiAccessToken"`
	TestEventToken  string             `json:"testEventToken" yaml:"testEventToken"`
	GoogleConsent   GoogleConsent      `json:"googleConsent" yaml:"googleConsent"`
	ServerContainer contemplate.Config `json:"serverContainer" yaml:"serverContainer"`
}
