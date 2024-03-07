package tagmanager

import (
	"context"
	"errors"

	"google.golang.org/api/option"
	"google.golang.org/api/tagmanager/v2"
)

var (
	ErrNotFound = errors.New("not found")
)

type (
	Client struct {
		notes         string
		accountID     string
		containerID   string
		workspaceID   string
		measurementID string
		folderName    string
		clientOptions []option.ClientOption
		// cache
		service   *tagmanager.Service
		clients   map[string]*tagmanager.Client
		folders   map[string]*tagmanager.Folder
		variables map[string]*tagmanager.Variable
		triggers  map[string]*tagmanager.Trigger
		tags      map[string]*tagmanager.Tag
	}
	ClientOption func(*Client)
)

// ------------------------------------------------------------------------------------------------
// ~ Options
// ------------------------------------------------------------------------------------------------

func ClientWithNotes(v string) ClientOption {
	return func(o *Client) {
		o.notes = v
	}
}

func ClientWithFolderName(v string) ClientOption {
	return func(o *Client) {
		o.folderName = v
	}
}

func ClientWithClientOptions(v ...option.ClientOption) ClientOption {
	return func(o *Client) {
		o.clientOptions = append(o.clientOptions, v...)
	}
}

// ------------------------------------------------------------------------------------------------
// ~ Constructor
// ------------------------------------------------------------------------------------------------

func NewClient(accountID, containerID, workspaceID, measurementID string, opts ...ClientOption) (*Client, error) {
	inst := &Client{
		accountID:     accountID,
		containerID:   containerID,
		workspaceID:   workspaceID,
		measurementID: measurementID,
		clients:       map[string]*tagmanager.Client{},
		folders:       map[string]*tagmanager.Folder{},
		variables:     map[string]*tagmanager.Variable{},
		triggers:      map[string]*tagmanager.Trigger{},
		tags:          map[string]*tagmanager.Tag{},
		notes:         "Managed by Sesamy. DO NOT EDIT.",
		folderName:    "Sesamy",
		clientOptions: []option.ClientOption{
			//// https://developers.google.com/tag-platform/tag-manager/api/v2/authorization#AboutAuthorization
			//option.WithScopes(
			//	"https://www.googleapis.com/auth/tagmanager.readonly",
			//	"https://www.googleapis.com/auth/tagmanager.edit.containers",
			//	// https://www.googleapis.com/auth/tagmanager.delete.containers
			//	//"https://www.googleapis.com/auth/tagmanager.edit.containerversions",
			//	//"https://www.googleapis.com/auth/tagmanager.publish",
			//	// https://www.googleapis.com/auth/tagmanager.manage.users
			//	// https://www.googleapis.com/auth/tagmanager.manage.accounts
			//),
		},
	}

	for _, opt := range opts {
		opt(inst)
	}

	return inst, nil
}

// ------------------------------------------------------------------------------------------------
// ~ Getter
// ------------------------------------------------------------------------------------------------

func (c *Client) AccountID() string {
	return c.accountID
}

func (c *Client) ContainerID() string {
	return c.containerID
}

func (c *Client) WorkspaceID() string {
	return c.workspaceID
}

func (c *Client) Notes() string {
	return c.notes
}

func (c *Client) MeasurementID() string {
	return c.measurementID
}

func (c *Client) FolderName() string {
	return c.folderName
}

func (c *Client) Service(ctx context.Context) (*tagmanager.Service, error) {
	if c.service == nil {
		v, err := tagmanager.NewService(ctx, c.clientOptions...)
		if err != nil {
			return nil, err
		}
		c.service = v
	}
	return c.service, nil
}

func (c *Client) AccountPath() string {
	return "accounts/" + c.accountID
}

func (c *Client) ContainerPath() string {
	return c.AccountPath() + "/containers/" + c.containerID
}

func (c *Client) WorkspacePath() string {
	return c.ContainerPath() + "/workspaces/" + c.workspaceID
}

func (c *Client) Client(ctx context.Context, name string) (*tagmanager.Client, error) {
	if _, ok := c.clients[name]; !ok {
		client, err := c.Service(ctx)
		if err != nil {
			return nil, err
		}

		r, err := client.Accounts.Containers.Workspaces.Clients.List(c.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		var found bool
		for _, value := range r.Client {
			if value.Name == name {
				c.clients[name] = value
				found = true
				break
			}
		}

		if !found {
			return nil, ErrNotFound
		}
	}
	return c.clients[name], nil
}

func (c *Client) Folder(ctx context.Context, name string) (*tagmanager.Folder, error) {
	if _, ok := c.folders[name]; !ok {
		client, err := c.Service(ctx)
		if err != nil {
			return nil, err
		}

		r, err := client.Accounts.Containers.Workspaces.Folders.List(c.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		var found bool
		for _, value := range r.Folder {
			if value.Name == name {
				c.folders[name] = value
				found = true
				break
			}
		}

		if !found {
			return nil, ErrNotFound
		}
	}
	return c.folders[name], nil
}

func (c *Client) Variable(ctx context.Context, name string) (*tagmanager.Variable, error) {
	if _, ok := c.variables[name]; !ok {
		client, err := c.Service(ctx)
		if err != nil {
			return nil, err
		}

		r, err := client.Accounts.Containers.Workspaces.Variables.List(c.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		var found bool
		for _, value := range r.Variable {
			if value.Name == name {
				c.variables[name] = value
				found = true
				break
			}
		}

		if !found {
			return nil, ErrNotFound
		}
	}
	return c.variables[name], nil
}

func (c *Client) Trigger(ctx context.Context, name string) (*tagmanager.Trigger, error) {
	if _, ok := c.triggers[name]; !ok {
		client, err := c.Service(ctx)
		if err != nil {
			return nil, err
		}

		r, err := client.Accounts.Containers.Workspaces.Triggers.List(c.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		var found bool
		for _, value := range r.Trigger {
			if value.Name == name {
				c.triggers[name] = value
				found = true
				break
			}
		}

		if !found {
			return nil, ErrNotFound
		}
	}
	return c.triggers[name], nil
}

func (c *Client) Tag(ctx context.Context, name string) (*tagmanager.Tag, error) {
	if _, ok := c.tags[name]; !ok {
		client, err := c.Service(ctx)
		if err != nil {
			return nil, err
		}

		r, err := client.Accounts.Containers.Workspaces.Tags.List(c.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		var found bool
		for _, value := range r.Tag {
			if value.Name == name {
				c.tags[name] = value
				found = true
				break
			}
		}

		if !found {
			return nil, ErrNotFound
		}
	}
	return c.tags[name], nil
}

// ------------------------------------------------------------------------------------------------
// ~ Public methods
// ------------------------------------------------------------------------------------------------

func (c *Client) UpsertGA4Client(ctx context.Context, name string) (*tagmanager.Client, error) {
	cache, err := c.Client(ctx, name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	s, err := c.Service(ctx)
	if err != nil {
		return nil, err
	}

	folder, err := c.Folder(ctx, c.folderName)
	if err != nil {
		return nil, err
	}

	obj := &tagmanager.Client{
		AccountId:      c.accountID,
		ContainerId:    c.containerID,
		WorkspaceId:    c.workspaceID,
		ParentFolderId: folder.FolderId,
		Name:           name,
		Notes:          c.notes,
		Parameter: []*tagmanager.Parameter{
			{
				Type:  "boolean",
				Key:   "activateGtagSupport",
				Value: "false",
			},
			{
				Type:  "boolean",
				Key:   "activateDefaultPaths",
				Value: "true",
			},
			{
				Type:  "template",
				Key:   "cookieManagement",
				Value: "js",
			},
		},
		Type: "gaaw_client",
	}

	if cache == nil {
		c.clients[name], err = s.Accounts.Containers.Workspaces.Clients.Create(c.WorkspacePath(), obj).Do()
	} else {
		c.clients[name], err = s.Accounts.Containers.Workspaces.Clients.Update(c.WorkspacePath()+"/clients/"+cache.ClientId, obj).Do()
	}
	if err != nil {
		return nil, err
	}

	return c.Client(ctx, name)
}

func (c *Client) UpsertGTMClient(ctx context.Context, name, webContainerMeasurementID string) (*tagmanager.Client, error) {
	cache, err := c.Client(ctx, name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	s, err := c.Service(ctx)
	if err != nil {
		return nil, err
	}

	folder, err := c.Folder(ctx, c.folderName)
	if err != nil {
		return nil, err
	}

	obj := &tagmanager.Client{
		AccountId:      c.accountID,
		ContainerId:    c.containerID,
		WorkspaceId:    c.workspaceID,
		ParentFolderId: folder.FolderId,
		Name:           name,
		Notes:          c.notes,
		Parameter: []*tagmanager.Parameter{
			{
				Type:  "boolean",
				Key:   "activateResponseCompression",
				Value: "true",
			},
			{
				Type:  "boolean",
				Key:   "activateGeoResolution",
				Value: "false",
			},
			{
				Type:  "boolean",
				Key:   "activateDependencyServing",
				Value: "true",
			},
			{
				Type: "list",
				Key:  "allowedContainerIds",
				List: []*tagmanager.Parameter{
					{
						Type: "map",
						Map: []*tagmanager.Parameter{
							{
								Type:  "template",
								Key:   "containerId",
								Value: webContainerMeasurementID,
							},
						},
					},
				},
			},
		},
		Type: "gtm_client",
	}

	if cache == nil {
		c.clients[name], err = s.Accounts.Containers.Workspaces.Clients.Create(c.WorkspacePath(), obj).Do()
	} else {
		c.clients[name], err = s.Accounts.Containers.Workspaces.Clients.Update(c.WorkspacePath()+"/clients/"+cache.ClientId, obj).Do()
	}
	if err != nil {
		return nil, err
	}

	return c.Client(ctx, name)
}

func (c *Client) UpsertFolder(ctx context.Context, name string) (*tagmanager.Folder, error) {
	cache, err := c.Folder(ctx, name)
	if err != nil && !errors.Is(err, ErrNotFound) && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	client, err := c.Service(ctx)
	if err != nil {
		return nil, err
	}

	obj := &tagmanager.Folder{
		AccountId:   c.accountID,
		ContainerId: c.containerID,
		WorkspaceId: c.workspaceID,
		Name:        name,
		Notes:       c.notes,
	}

	if cache == nil {
		c.folders[name], err = client.Accounts.Containers.Workspaces.Folders.Create(c.WorkspacePath(), obj).Do()
	} else {
		c.folders[name], err = client.Accounts.Containers.Workspaces.Folders.Update(c.WorkspacePath()+"/folders/"+cache.FolderId, obj).Do()
	}
	if err != nil {
		return nil, err
	}

	return c.Folder(ctx, name)
}

func (c *Client) UpsertVariable(ctx context.Context, obj *tagmanager.Variable) (*tagmanager.Variable, error) {
	cache, err := c.Variable(ctx, obj.Name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	client, err := c.Service(ctx)
	if err != nil {
		return nil, err
	}

	folder, err := c.Folder(ctx, c.folderName)
	if err != nil {
		return nil, err
	}
	obj.ParentFolderId = folder.FolderId

	if cache == nil {
		c.variables[obj.Name], err = client.Accounts.Containers.Workspaces.Variables.Create(c.WorkspacePath(), obj).Do()
	} else {
		c.variables[obj.Name], err = client.Accounts.Containers.Workspaces.Variables.Update(c.WorkspacePath()+"/variables/"+cache.VariableId, obj).Do()
	}
	if err != nil {
		return nil, err
	}

	return c.Variable(ctx, obj.Name)
}

func (c *Client) UpsertConstantVariable(ctx context.Context, name, value string) (*tagmanager.Variable, error) {
	obj := &tagmanager.Variable{
		AccountId:   c.accountID,
		ContainerId: c.containerID,
		WorkspaceId: c.workspaceID,
		Name:        "Constant." + name,
		Notes:       c.notes,
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "value",
				Type:  "template",
				Value: value,
			},
		},
		Type: "c",
	}
	return c.UpsertVariable(ctx, obj)
}

func (c *Client) UpsertEventModelVariable(ctx context.Context, name string) (*tagmanager.Variable, error) {
	obj := &tagmanager.Variable{
		AccountId:   c.accountID,
		ContainerId: c.containerID,
		WorkspaceId: c.workspaceID,
		Name:        "EventModel." + name,
		Notes:       c.notes,
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "dataLayerVersion",
				Type:  "integer",
				Value: "2",
			},
			{
				Key:   "setDefaultValue",
				Type:  "boolean",
				Value: "false",
			},
			{
				Key:   "name",
				Type:  "template",
				Value: "eventModel." + name,
			},
		},
		Type: "v",
	}
	return c.UpsertVariable(ctx, obj)
}

func (c *Client) UpsertCustomEventTrigger(ctx context.Context, name string) (*tagmanager.Trigger, error) {
	name = "Event." + name
	cache, err := c.Trigger(ctx, name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	client, err := c.Service(ctx)
	if err != nil {
		return nil, err
	}

	folder, err := c.Folder(ctx, c.folderName)
	if err != nil {
		return nil, err
	}

	obj := &tagmanager.Trigger{
		AccountId:      c.accountID,
		ContainerId:    c.containerID,
		WorkspaceId:    c.workspaceID,
		ParentFolderId: folder.FolderId,
		Type:           "customEvent",
		Name:           name,
		Notes:          c.notes,
		CustomEventFilter: []*tagmanager.Condition{
			{
				Parameter: []*tagmanager.Parameter{
					{
						Key:   "arg0",
						Type:  "template",
						Value: "{{_event}}",
					},
					{
						Key:   "arg1",
						Type:  "template",
						Value: name,
					},
				},
				Type: "equals",
			},
		},
	}

	if cache == nil {
		c.triggers[name], err = client.Accounts.Containers.Workspaces.Triggers.Create(c.WorkspacePath(), obj).Do()
	} else {
		c.triggers[name], err = client.Accounts.Containers.Workspaces.Triggers.Update(c.WorkspacePath()+"/triggers/"+cache.TriggerId, obj).Do()
	}
	if err != nil {
		return nil, err
	}

	return c.Trigger(ctx, name)
}

func (c *Client) UpsertClientTrigger(ctx context.Context, name, clientName string) (*tagmanager.Trigger, error) {
	name = "Client." + name
	cache, err := c.Trigger(ctx, name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	client, err := c.Service(ctx)
	if err != nil {
		return nil, err
	}

	folder, err := c.Folder(ctx, c.folderName)
	if err != nil {
		return nil, err
	}

	obj := &tagmanager.Trigger{
		AccountId:      c.accountID,
		ContainerId:    c.containerID,
		WorkspaceId:    c.workspaceID,
		ParentFolderId: folder.FolderId,
		Type:           "always",
		Name:           name,
		Notes:          c.notes,
		Filter: []*tagmanager.Condition{
			{
				Parameter: []*tagmanager.Parameter{
					{
						Key:   "arg0",
						Type:  "template",
						Value: "{{Client Name}}",
					},
					{
						Key:   "arg1",
						Type:  "template",
						Value: clientName,
					},
				},
				Type: "equals",
			},
		},
	}

	if cache == nil {
		c.triggers[name], err = client.Accounts.Containers.Workspaces.Triggers.Create(c.WorkspacePath(), obj).Do()
	} else {
		c.triggers[name], err = client.Accounts.Containers.Workspaces.Triggers.Update(c.WorkspacePath()+"/triggers/"+cache.TriggerId, obj).Do()
	}
	if err != nil {
		return nil, err
	}

	return c.Trigger(ctx, name)
}

func (c *Client) UpsertGA4WebTag(ctx context.Context, name string, parameters []string, measurementID *tagmanager.Variable, trigger *tagmanager.Trigger) (*tagmanager.Tag, error) {
	name = "GA4." + name

	cache, err := c.Tag(ctx, name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	client, err := c.Service(ctx)
	if err != nil {
		return nil, err
	}

	folder, err := c.Folder(ctx, c.folderName)
	if err != nil {
		return nil, err
	}

	obj := &tagmanager.Tag{
		AccountId:       c.accountID,
		ContainerId:     c.containerID,
		WorkspaceId:     c.workspaceID,
		FiringTriggerId: []string{trigger.TriggerId},
		ParentFolderId:  folder.FolderId,
		Name:            name,
		Notes:           c.notes,
		Parameter: []*tagmanager.Parameter{
			{
				Type:  "boolean",
				Key:   "sendEcommerceData",
				Value: "false",
			},
			{
				Type:  "template",
				Key:   "getEcommerceDataFrom",
				Value: "dataLayer",
			},
			{
				Type:  "boolean",
				Key:   "enhancedUserId",
				Value: "false",
			},
			{
				Type:  "template",
				Key:   "eventName",
				Value: trigger.Name,
			},
			{
				Type:  "template",
				Key:   "measurementIdOverride",
				Value: "{{" + measurementID.Name + "}}",
			},
		},
		Type: "gaawe",
	}

	for _, parameterName := range parameters {
		if _, err := c.UpsertEventModelVariable(ctx, parameterName); err != nil {
			return nil, err
		}
		obj.Parameter = append(obj.Parameter, &tagmanager.Parameter{
			Type: "list",
			Key:  "eventSettingsTable",
			List: []*tagmanager.Parameter{
				{
					Type: "map",
					Map: []*tagmanager.Parameter{
						{
							Type:  "template",
							Key:   "parameter",
							Value: parameterName,
						},
						{
							Type:  "template",
							Key:   "parameterValue",
							Value: "{{EventModel." + parameterName + "}}",
						},
					},
				},
			},
		})
	}

	if cache == nil {
		c.tags[name], err = client.Accounts.Containers.Workspaces.Tags.Create(c.WorkspacePath(), obj).Do()
	} else {
		c.tags[name], err = client.Accounts.Containers.Workspaces.Tags.Update(c.WorkspacePath()+"/tags/"+cache.TagId, obj).Do()
	}
	if err != nil {
		return nil, err
	}

	return c.Tag(ctx, name)
}

func (c *Client) UpsertGA4ServerTag(ctx context.Context, name string, measurementID *tagmanager.Variable, trigger *tagmanager.Trigger) (*tagmanager.Tag, error) {
	cache, err := c.Tag(ctx, name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	client, err := c.Service(ctx)
	if err != nil {
		return nil, err
	}

	folder, err := c.Folder(ctx, c.folderName)
	if err != nil {
		return nil, err
	}

	obj := &tagmanager.Tag{
		AccountId:       c.accountID,
		ContainerId:     c.containerID,
		WorkspaceId:     c.workspaceID,
		FiringTriggerId: []string{trigger.TriggerId},
		ParentFolderId:  folder.FolderId,
		Name:            name,
		Notes:           c.notes,
		Parameter: []*tagmanager.Parameter{
			{
				Type:  "boolean",
				Key:   "redactVisitorIp",
				Value: "false",
			},
			{
				Type:  "template",
				Key:   "epToIncludeDropdown",
				Value: "all",
			},
			{
				Type:  "boolean",
				Key:   "enableGoogleSignals",
				Value: "false",
			},
			{
				Type:  "template",
				Key:   "upToIncludeDropdown",
				Value: "all",
			},
			{
				Type:  "template",
				Key:   "measurementId",
				Value: "{{" + measurementID.Name + "}}",
			},
			{
				Type:  "boolean",
				Key:   "enableEuid",
				Value: "false",
			},
		},
		Type: "sgtmgaaw",
	}

	if cache == nil {
		c.tags[name], err = client.Accounts.Containers.Workspaces.Tags.Create(c.WorkspacePath(), obj).Do()
	} else {
		c.tags[name], err = client.Accounts.Containers.Workspaces.Tags.Update(c.WorkspacePath()+"/tags/"+cache.TagId, obj).Do()
	}
	if err != nil {
		return nil, err
	}

	return c.Tag(ctx, name)
}
