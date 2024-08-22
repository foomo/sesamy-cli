package config

type GoogleGTag struct {
	Enabled        bool  `json:"enabled" yaml:"enabled"`
	Priority       int64 `json:"priority" yaml:"priority"`
	EcommerceItems bool  `json:"ecommerceItems" yaml:"ecommerceItems"`
}
