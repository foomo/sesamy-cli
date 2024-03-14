package config

type TagmangerTagPrefixes struct {
	GA4Event       string `yaml:"ga4_event"`
	GoogleTag      string `yaml:"google_tag"`
	ServerGA4Event string `yaml:"server_ga4_event"`
}

func (p TagmangerTagPrefixes) GA4EventName(name string) string {
	return p.GA4Event + name
}

func (p TagmangerTagPrefixes) GoogleTagName(name string) string {
	return p.GoogleTag + name
}

func (p TagmangerTagPrefixes) ServerGA4EventName(name string) string {
	return p.ServerGA4Event + name
}
