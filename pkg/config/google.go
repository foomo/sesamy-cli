package config

type Google struct {
	GT              GT     `yaml:"gt"`
	GA4             GA4    `yaml:"ga4"`
	GTM             GTM    `yaml:"gtm"`
	RequestQuota    int    `yaml:"request_quota"`
	CredentialsFile string `yaml:"credentials_file"`
	CredentialsJSON string `yaml:"credentials_json"`
}
