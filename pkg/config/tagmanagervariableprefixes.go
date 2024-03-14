package config

type TagmangerVariablePrefixes struct {
	Constant        string `yaml:"constant"`
	EventModel      string `yaml:"event_model"`
	GTEventSettings string `yaml:"gt_event_settings"`
	GTSettings      string `yaml:"gt_settings"`
}

func (p TagmangerVariablePrefixes) ConstantName(name string) string {
	return p.Constant + name
}

func (p TagmangerVariablePrefixes) EventModelName(name string) string {
	return p.EventModel + name
}

func (p TagmangerVariablePrefixes) GTEventSettingsName(name string) string {
	return p.GTEventSettings + name
}

func (p TagmangerVariablePrefixes) GTSettingsName(name string) string {
	return p.GTSettings + name
}
