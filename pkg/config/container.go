package config

type Container struct {
	ContainerID   string `yaml:"container_id"`
	WorkspaceID   string `yaml:"workspace_id"`
	MeasurementID string `yaml:"measurement_id"`
}
