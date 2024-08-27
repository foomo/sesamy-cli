package config

import (
	"github.com/foomo/gocontemplate/pkg/contemplate"
)

type Emarsys struct {
	Enabled         bool               `json:"enabled" yaml:"enabled"`
	MerchantID      string             `json:"merchantId" yaml:"merchantId"`
	GoogleConsent   GoogleConsent      `json:"googleConsent" yaml:"googleConsent"`
	WebContainer    contemplate.Config `json:"webContainer" yaml:"webContainer"`
	ServerContainer contemplate.Config `json:"serverContainer" yaml:"serverContainer"`
}
