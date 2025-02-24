package config

type LookupTable struct {
	// Input source
	Input string `json:"input" yaml:"input"`
	// Key value data map
	KeyTable map[string]string `json:"keyTable" yaml:"keyTable"`
	// Vaule key data map
	ValueTable map[string]string `json:"valueTable" yaml:"valueTable"`
}
