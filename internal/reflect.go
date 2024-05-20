package internal

type PackageConfig struct {
	Path  string   `json:"path" yaml:"path"`
	Types []string `json:"types" yaml:"types"`
}
