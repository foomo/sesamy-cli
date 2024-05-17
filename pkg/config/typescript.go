package config

type Typescript struct {
	Packages   Packages `yaml:"packages"`
	OutputPath string   `yaml:"output_path"`
}
