package config

type GTM struct {
	AccountID string    `yaml:"account_id"`
	Web       Container `yaml:"web"`
	Server    Container `yaml:"server"`
}
