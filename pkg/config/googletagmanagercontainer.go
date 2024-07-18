package config

type GoogleTagManagerContainer struct {
	TagID       string `json:"tagId" yaml:"tagId"`
	ContainerID string `json:"containerId" yaml:"containerId"`
	WorkspaceID string `json:"workspaceId" yaml:"workspaceId"`
}
