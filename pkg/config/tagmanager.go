package config

type Tagmanager struct {
	Packages Packages           `yaml:"packages"`
	Prefixes TagmanagerPrefixes `yaml:"prefixes"`
}
