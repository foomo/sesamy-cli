package config

import (
	"github.com/foomo/gocontemplate/pkg/contemplate"
)

type Emarsys struct {
	Enabled          bool               `json:"enabled" yaml:"enabled"`
	MerchantID       string             `json:"merchantId" yaml:"merchantId"`
	NewPageViewEvent string             `json:"newPageViewEvent" yaml:"newPageViewEvent"`
	WebContainer     contemplate.Config `json:"webContainer" yaml:"webContainer"`
	ServerContainer  contemplate.Config `json:"serverContainer" yaml:"serverContainer"`
}
