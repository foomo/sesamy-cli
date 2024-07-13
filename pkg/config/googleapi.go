package config

import (
	googleapioption "google.golang.org/api/option"
)

type GoogleAPI struct {
	Credentials     string `json:"credentials" yaml:"credentials"`
	CredentialsFile string `json:"credentialsFile" yaml:"credentialsFile"`
	RequestQuota    int    `json:"requestQuota" yaml:"requestQuota"`
}

func (c GoogleAPI) GetClientOption() googleapioption.ClientOption {
	var ret googleapioption.ClientOption
	if c.CredentialsFile != "" {
		ret = googleapioption.WithCredentialsFile(c.CredentialsFile)
	} else {
		ret = googleapioption.WithCredentialsJSON([]byte(c.Credentials))
	}
	return ret
}
