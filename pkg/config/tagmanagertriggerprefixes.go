package config

type TagmangerTriggerPrefixes struct {
	Client      string `yaml:"client"`
	CustomEvent string `yaml:"custom_event"`
}

func (p TagmangerTriggerPrefixes) ClientName(name string) string {
	return p.Client + name
}

func (p TagmangerTriggerPrefixes) CustomEventName(name string) string {
	return p.CustomEvent + name
}
