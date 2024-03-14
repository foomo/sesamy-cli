package config

type Config struct {
	Google Google `yaml:"google"`
	// https://github.com/gzuidhof/tygo
	Typescript Typescript `yaml:"typescript"`
	// https://github.com/gzuidhof/tygo
	Tagmanager Tagmanager `yaml:"tagmanager"`
}
