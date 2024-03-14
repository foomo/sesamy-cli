package config

type Google struct {
	GA4                GA4    `yaml:"ga4"`
	GTM                GTM    `yaml:"gtm"`
	CredentialsFile    string `yaml:"credentials_file"`
	CredentialsJSON    string `yaml:"credentials_json"`
	ServerContainerURL string `yaml:"server_container_url"`
}
