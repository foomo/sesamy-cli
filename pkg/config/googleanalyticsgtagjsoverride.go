package config

type GoogleAnalyticsGTagJSOverride struct {
	// Enable override
	Enabled bool `json:"enabled" yaml:"enabled"`
	// Client priority
	Priority int64 `json:"priority" yaml:"priority"`
	// Allow sending items for non ecommerce events
	EcommerceItems bool `json:"ecommerceItems" yaml:"ecommerceItems"`
}
