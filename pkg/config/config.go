package config

import (
	"github.com/gzuidhof/tygo/tygo"
)

type Config struct {
	Google Google `yaml:"google"`
	// https://github.com/gzuidhof/tygo
	Typescript Source `yaml:"typescript"`
	// https://github.com/gzuidhof/tygo
	Tagmanager Source `yaml:"tagmanager"`
}

type Google struct {
	GA4                GA4    `yaml:"ga4"`
	GTM                GTM    `yaml:"gtm"`
	CredentialsFile    string `yaml:"credentials_file"`
	CredentialsJSON    string `yaml:"credentials_json"`
	ServerContainerURL string `yaml:"server_container_url"`
}

type GA4 struct {
	MeasurementID string `yaml:"measurement_id"`
}

type GTM struct {
	AccountID string    `yaml:"account_id"`
	Web       Container `yaml:"web"`
	Server    Container `yaml:"server"`
}

type Source struct {
	Packages     []*tygo.PackageConfig `yaml:"packages"`
	TypeMappings map[string]string     `yaml:"type_mappings"`
}

func (e Source) PackageNames() []string {
	ret := make([]string, len(e.Packages))
	for i, value := range e.Packages {
		ret[i] = value.Path
	}
	return ret
}

func (e Source) PackageConfig(path string) *tygo.PackageConfig {
	for _, value := range e.Packages {
		if value.Path == path {
			return value
		}
	}
	return nil
}

type Container struct {
	ContainerID   string `yaml:"container_id"`
	WorkspaceID   string `yaml:"workspace_id"`
	MeasurementID string `yaml:"measurement_id"`
}
