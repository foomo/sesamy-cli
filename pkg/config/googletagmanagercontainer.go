package config

type GoogleTagManagerContainer struct {
	// The container tag id
	TagID string `json:"tagId" yaml:"tagId"`
	// The container id
	ContainerID string `json:"containerId" yaml:"containerId"`
	// The workspace id that should be used by the api
	WorkspaceID string `json:"workspaceId" yaml:"workspaceId"`
}
