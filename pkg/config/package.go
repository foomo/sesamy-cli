package config

type Package struct {
	Path   string   `yaml:"path"`
	Events []string `yaml:"events"`
}

func (c Package) ExportEvent(event string) bool {
	for _, name := range c.Events {
		if name == event {
			return true
		}
	}
	return false
}
