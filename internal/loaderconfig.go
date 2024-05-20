package internal

type LoaderConfig struct {
	Packages []*PackageConfig `json:"packages" yaml:"packages"`
}

func (c *LoaderConfig) Package(path string) *PackageConfig {
	for _, value := range c.Packages {
		if value.Path == path {
			return value
		}
	}
	return nil
}

func (c *LoaderConfig) PackagePaths() []string {
	ret := make([]string, len(c.Packages))
	for i, value := range c.Packages {
		ret[i] = value.Path
	}
	return ret
}
