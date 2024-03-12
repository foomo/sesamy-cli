package tagmanager

import (
	"context"
	"errors"
	"sort"
	"time"

	"google.golang.org/api/option"
	"google.golang.org/api/tagmanager/v2"
)

var (
	ErrNotFound = errors.New("not found")
)

type (
	Client struct {
		notes            string
		accountID        string
		containerID      string
		workspaceID      string
		measurementID    string
		folderName       string
		clientOptions    []option.ClientOption
		requestThrottler *time.Ticker
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

func ClientWithRequestQuota(v int) ClientOption {
	return func(o *Client) {
		o.requestThrottler = time.NewTicker((100 * time.Second) / time.Duration(v))
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

func NewClient(ctx context.Context, accountID, containerID, workspaceID, measurementID string, opts ...ClientOption) (*Client, error) {
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
			option.WithRequestReason("Sesamy container provisioning"),
		},
	}

	for _, opt := range opts {
		opt(inst)
	}

	if value, err := tagmanager.NewService(ctx, inst.clientOptions...); err != nil {
		return nil, err
	} else {
		inst.service = value
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

func (c *Client) Service() *tagmanager.Service {
	if c.requestThrottler != nil {
		<-c.requestThrottler.C
	}
	return c.service
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

func (c *Client) Client(name string) (*tagmanager.Client, error) {
	if _, ok := c.clients[name]; !ok {
		r, err := c.Service().Accounts.Containers.Workspaces.Clients.List(c.WorkspacePath()).Do()
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

func (c *Client) Folder(name string) (*tagmanager.Folder, error) {
	if _, ok := c.folders[name]; !ok {
		r, err := c.Service().Accounts.Containers.Workspaces.Folders.List(c.WorkspacePath()).Do()
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

func (c *Client) Variable(name string) (*tagmanager.Variable, error) {
	if _, ok := c.variables[name]; !ok {
		r, err := c.Service().Accounts.Containers.Workspaces.Variables.List(c.WorkspacePath()).Do()
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

func (c *Client) Trigger(name string) (*tagmanager.Trigger, error) {
	if _, ok := c.triggers[name]; !ok {
		r, err := c.Service().Accounts.Containers.Workspaces.Triggers.List(c.WorkspacePath()).Do()
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

func (c *Client) Tag(name string) (*tagmanager.Tag, error) {
	if _, ok := c.tags[name]; !ok {
		r, err := c.Service().Accounts.Containers.Workspaces.Tags.List(c.WorkspacePath()).Do()
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

func (c *Client) UpsertGA4Client(name string) (*tagmanager.Client, error) {
	cache, err := c.Client(name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	folder, err := c.Folder(c.folderName)
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
		c.clients[name], err = c.Service().Accounts.Containers.Workspaces.Clients.Create(c.WorkspacePath(), obj).Do()
	} else {
		c.clients[name], err = c.Service().Accounts.Containers.Workspaces.Clients.Update(c.WorkspacePath()+"/clients/"+cache.ClientId, obj).Do()
	}
	if err != nil {
		return nil, err
	}

	return c.Client(name)
}

func (c *Client) UpsertGTMClient(name, webContainerMeasurementID string) (*tagmanager.Client, error) {
	cache, err := c.Client(name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	folder, err := c.Folder(c.folderName)
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
		c.clients[name], err = c.Service().Accounts.Containers.Workspaces.Clients.Create(c.WorkspacePath(), obj).Do()
	} else {
		c.clients[name], err = c.Service().Accounts.Containers.Workspaces.Clients.Update(c.WorkspacePath()+"/clients/"+cache.ClientId, obj).Do()
	}
	if err != nil {
		return nil, err
	}

	return c.Client(name)
}

func (c *Client) UpsertFolder(name string) (*tagmanager.Folder, error) {
	cache, err := c.Folder(name)
	if err != nil && !errors.Is(err, ErrNotFound) && !errors.Is(err, ErrNotFound) {
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
		c.folders[name], err = c.Service().Accounts.Containers.Workspaces.Folders.Create(c.WorkspacePath(), obj).Do()
	} else {
		c.folders[name], err = c.Service().Accounts.Containers.Workspaces.Folders.Update(c.WorkspacePath()+"/folders/"+cache.FolderId, obj).Do()
	}
	if err != nil {
		return nil, err
	}

	return c.Folder(name)
}

func (c *Client) UpsertVariable(variable *tagmanager.Variable) (*tagmanager.Variable, error) {
	cache, err := c.Variable(variable.Name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	folder, err := c.Folder(c.folderName)
	if err != nil {
		return nil, err
	}
	variable.ParentFolderId = folder.FolderId

	if cache == nil {
		c.variables[variable.Name], err = c.Service().Accounts.Containers.Workspaces.Variables.Create(c.WorkspacePath(), variable).Do()
	} else {
		c.variables[variable.Name], err = c.Service().Accounts.Containers.Workspaces.Variables.Update(c.WorkspacePath()+"/variables/"+cache.VariableId, variable).Do()
	}
	if err != nil {
		return nil, err
	}

	return c.Variable(variable.Name)
}

func (c *Client) UpsertConstantVariable(name, value string) (*tagmanager.Variable, error) {
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
	return c.UpsertVariable(obj)
}

func (c *Client) UpsertEventModelVariable(name string) (*tagmanager.Variable, error) {
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
	return c.UpsertVariable(obj)
}

func (c *Client) UpsertGTEventSettingsVariable(name string, variables map[string]*tagmanager.Variable) (*tagmanager.Variable, error) {
	fullname := "GTEventSettings." + name

	parameters := make([]string, 0, len(variables))
	for k := range variables {
		parameters = append(parameters, k)
	}
	sort.Strings(parameters)

	list := make([]*tagmanager.Parameter, len(parameters))
	for i, parameter := range parameters {
		list[i] = &tagmanager.Parameter{
			Type: "map",
			Map: []*tagmanager.Parameter{
				{
					Key:   "parameter",
					Type:  "template",
					Value: parameter,
				},
				{
					Key:   "parameterValue",
					Type:  "template",
					Value: "{{" + variables[parameter].Name + "}}",
				},
			},
		}
	}

	obj := &tagmanager.Variable{
		AccountId:   c.accountID,
		ContainerId: c.containerID,
		WorkspaceId: c.workspaceID,
		Name:        fullname,
		Notes:       c.notes,
		Parameter: []*tagmanager.Parameter{
			{
				Key:  "eventSettingsTable",
				Type: "list",
				List: list,
			},
		},
		Type: "gtes",
	}

	return c.UpsertVariable(obj)
}

func (c *Client) UpsertGoogleTagSettingsVariable(name string, variables map[string]*tagmanager.Variable) (*tagmanager.Variable, error) {
	parameters := make([]string, 0, len(variables))
	for k := range variables {
		parameters = append(parameters, k)
	}
	sort.Strings(parameters)

	list := make([]*tagmanager.Parameter, len(parameters))
	for i, parameter := range parameters {
		list[i] = &tagmanager.Parameter{
			Type: "map",
			Map: []*tagmanager.Parameter{
				{
					Key:   "parameter",
					Type:  "template",
					Value: parameter,
				},
				{
					Key:   "parameterValue",
					Type:  "template",
					Value: "{{" + variables[parameter].Name + "}}",
				},
			},
		}
	}

	obj := &tagmanager.Variable{
		AccountId:   c.accountID,
		ContainerId: c.containerID,
		WorkspaceId: c.workspaceID,
		Name:        name,
		Notes:       c.notes,
		Parameter: []*tagmanager.Parameter{
			{
				Key:  "configSettingsTable",
				Type: "list",
				List: list,
			},
		},
		Type: "gtes",
	}

	return c.UpsertVariable(obj)
}

func (c *Client) UpsertCustomEventTrigger(name string) (*tagmanager.Trigger, error) {
	fullname := "Event." + name
	cache, err := c.Trigger(fullname)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	folder, err := c.Folder(c.folderName)
	if err != nil {
		return nil, err
	}

	obj := &tagmanager.Trigger{
		AccountId:      c.accountID,
		ContainerId:    c.containerID,
		WorkspaceId:    c.workspaceID,
		ParentFolderId: folder.FolderId,
		Type:           "customEvent",
		Name:           fullname,
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
		c.triggers[fullname], err = c.Service().Accounts.Containers.Workspaces.Triggers.Create(c.WorkspacePath(), obj).Do()
	} else {
		c.triggers[fullname], err = c.Service().Accounts.Containers.Workspaces.Triggers.Update(c.WorkspacePath()+"/triggers/"+cache.TriggerId, obj).Do()
	}
	if err != nil {
		return nil, err
	}

	return c.Trigger(fullname)
}

func (c *Client) UpsertClientTrigger(name string, client *tagmanager.Client) (*tagmanager.Trigger, error) {
	fullname := "Client." + name
	cache, err := c.Trigger(fullname)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	folder, err := c.Folder(c.folderName)
	if err != nil {
		return nil, err
	}

	obj := &tagmanager.Trigger{
		AccountId:      c.accountID,
		ContainerId:    c.containerID,
		WorkspaceId:    c.workspaceID,
		ParentFolderId: folder.FolderId,
		Type:           "always",
		Name:           fullname,
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
						Value: client.Name,
					},
				},
				Type: "equals",
			},
		},
	}

	if cache == nil {
		c.triggers[fullname], err = c.Service().Accounts.Containers.Workspaces.Triggers.Create(c.WorkspacePath(), obj).Do()
	} else {
		c.triggers[fullname], err = c.Service().Accounts.Containers.Workspaces.Triggers.Update(c.WorkspacePath()+"/triggers/"+cache.TriggerId, obj).Do()
	}
	if err != nil {
		return nil, err
	}

	return c.Trigger(fullname)
}

func (c *Client) UpsertGA4WebTag(name string, eventSettings *tagmanager.Variable, measurementID *tagmanager.Variable, trigger *tagmanager.Trigger) (*tagmanager.Tag, error) {
	fullname := "GA4." + name

	cache, err := c.Tag(fullname)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	folder, err := c.Folder(c.folderName)
	if err != nil {
		return nil, err
	}

	obj := &tagmanager.Tag{
		AccountId:       c.accountID,
		ContainerId:     c.containerID,
		WorkspaceId:     c.workspaceID,
		FiringTriggerId: []string{trigger.TriggerId},
		ParentFolderId:  folder.FolderId,
		Name:            fullname,
		Notes:           c.notes,
		Parameter: []*tagmanager.Parameter{
			{
				Type:  "boolean",
				Key:   "sendEcommerceData",
				Value: "false",
			},
			{
				Type:  "boolean",
				Key:   "enhancedUserId",
				Value: "false",
			},
			{
				Type:  "template",
				Key:   "eventName",
				Value: name,
			},
			{
				Type:  "template",
				Key:   "measurementIdOverride",
				Value: "{{" + measurementID.Name + "}}",
			},
			{
				Type:  "template",
				Key:   "eventSettingsVariable",
				Value: "{{" + eventSettings.Name + "}}",
			},
		},
		Type: "gaawe",
	}

	// if len(parameters) > 0 {
	// 	list := &tagmanager.Parameter{
	// 		Type: "list",
	// 		Key:  "eventSettingsTable",
	// 		List: []*tagmanager.Parameter{},
	// 	}
	//
	// 	for _, parameterName := range parameters {
	// 		param, err := c.UpsertEventModelVariable(parameterName)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		list.List = append(list.List, &tagmanager.Parameter{
	// 			Type: "map",
	// 			Map: []*tagmanager.Parameter{
	// 				{
	// 					Type:  "template",
	// 					Key:   "parameter",
	// 					Value: parameterName,
	// 				},
	// 				{
	// 					Type:  "template",
	// 					Key:   "parameterValue",
	// 					Value: "{{" + param.Name + "}}",
	// 				},
	// 			},
	// 		})
	// 	}
	// 	obj.Parameter = append(obj.Parameter, list)
	// }

	if cache == nil {
		c.tags[fullname], err = c.Service().Accounts.Containers.Workspaces.Tags.Create(c.WorkspacePath(), obj).Do()
	} else {
		c.tags[fullname], err = c.Service().Accounts.Containers.Workspaces.Tags.Update(c.WorkspacePath()+"/tags/"+cache.TagId, obj).Do()
	}
	if err != nil {
		return nil, err
	}

	return c.Tag(fullname)
}

func (c *Client) UpsertGoogleTagWebTag(name string, measurementID *tagmanager.Variable, configSettings *tagmanager.Variable) (*tagmanager.Tag, error) {
	cache, err := c.Tag(name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	folder, err := c.Folder(c.folderName)
	if err != nil {
		return nil, err
	}

	obj := &tagmanager.Tag{
		AccountId:       c.accountID,
		ContainerId:     c.containerID,
		WorkspaceId:     c.workspaceID,
		FiringTriggerId: []string{"2147479573"},
		ParentFolderId:  folder.FolderId,
		Name:            name,
		Notes:           c.notes,
		Parameter: []*tagmanager.Parameter{
			{
				Key:   "tagId",
				Type:  "template",
				Value: "{{" + measurementID.Name + "}}",
			},
			{
				Key:   "configSettingsVariable",
				Type:  "template",
				Value: "{{" + configSettings.Name + "}}",
			},
		},
		Type: "googtag",
	}

	if cache == nil {
		c.tags[name], err = c.Service().Accounts.Containers.Workspaces.Tags.Create(c.WorkspacePath(), obj).Do()
	} else {
		c.tags[name], err = c.Service().Accounts.Containers.Workspaces.Tags.Update(c.WorkspacePath()+"/tags/"+cache.TagId, obj).Do()
	}
	if err != nil {
		return nil, err
	}

	return c.Tag(name)
}

func (c *Client) UpsertGA4ServerTag(name string, measurementID *tagmanager.Variable, trigger *tagmanager.Trigger) (*tagmanager.Tag, error) {
	cache, err := c.Tag(name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	folder, err := c.Folder(c.folderName)
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
		c.tags[name], err = c.Service().Accounts.Containers.Workspaces.Tags.Create(c.WorkspacePath(), obj).Do()
	} else {
		c.tags[name], err = c.Service().Accounts.Containers.Workspaces.Tags.Update(c.WorkspacePath()+"/tags/"+cache.TagId, obj).Do()
	}
	if err != nil {
		return nil, err
	}

	return c.Tag(name)
}
