package tagmanager

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"google.golang.org/api/option"
	"google.golang.org/api/tagmanager/v2"
)

var (
	ErrNotFound = errors.New("not found")
)

type (
	Client struct {
		l                *slog.Logger
		notes            string
		accountID        string
		containerID      string
		workspaceID      string
		measurementID    string
		folderName       string
		clientOptions    []option.ClientOption
		requestThrottler *time.Ticker
		// cache
		service          *tagmanager.Service
		clients          map[string]*tagmanager.Client
		folders          map[string]*tagmanager.Folder
		variables        map[string]*tagmanager.Variable
		builtInVariables map[string]*tagmanager.BuiltInVariable
		triggers         map[string]*tagmanager.Trigger
		tags             map[string]*tagmanager.Tag
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

func NewClient(ctx context.Context, l *slog.Logger, accountID, containerID, workspaceID, measurementID string, opts ...ClientOption) (*Client, error) {
	inst := &Client{
		l:             l,
		accountID:     accountID,
		containerID:   containerID,
		workspaceID:   workspaceID,
		measurementID: measurementID,
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

// ------------------------------------------------------------------------------------------------
// ~ Public methods
// ------------------------------------------------------------------------------------------------

func (c *Client) Client(name string) (*tagmanager.Client, error) {
	if c.clients == nil {
		c.clients = map[string]*tagmanager.Client{}

		c.l.Debug("listing", "type", "Client")
		r, err := c.Service().Accounts.Containers.Workspaces.Clients.List(c.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		for _, value := range r.Client {
			c.clients[name] = value
		}
	}

	if _, ok := c.clients[name]; !ok {
		return nil, ErrNotFound
	}

	return c.clients[name], nil
}

func (c *Client) Folder(name string) (*tagmanager.Folder, error) {
	if c.folders == nil {
		c.folders = map[string]*tagmanager.Folder{}

		c.l.Debug("listing", "type", "Folder")
		r, err := c.Service().Accounts.Containers.Workspaces.Folders.List(c.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		for _, value := range r.Folder {
			c.folders[name] = value
		}
	}

	if _, ok := c.folders[name]; !ok {
		return nil, ErrNotFound
	}

	return c.folders[name], nil
}

func (c *Client) Variable(name string) (*tagmanager.Variable, error) {
	if c.variables == nil {
		c.variables = map[string]*tagmanager.Variable{}

		c.l.Debug("listing", "type", "Variable")
		r, err := c.Service().Accounts.Containers.Workspaces.Variables.List(c.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		for _, value := range r.Variable {
			c.variables[name] = value
		}
	}

	if _, ok := c.variables[name]; !ok {
		return nil, ErrNotFound
	}
	return c.variables[name], nil
}

func (c *Client) BuiltInVariable(name string) (*tagmanager.Variable, error) {
	if c.builtInVariables == nil {
		c.builtInVariables = map[string]*tagmanager.BuiltInVariable{}

		c.l.Debug("listing", "type", "Built-In Variable")
		r, err := c.Service().Accounts.Containers.Workspaces.BuiltInVariables.List(c.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		for _, value := range r.BuiltInVariable {
			c.builtInVariables[value.Type] = value
		}
	}

	if _, ok := c.variables[name]; !ok {
		return nil, ErrNotFound
	}

	return c.variables[name], nil
}

func (c *Client) Trigger(name string) (*tagmanager.Trigger, error) {
	if c.triggers == nil {
		c.triggers = map[string]*tagmanager.Trigger{}

		c.l.Debug("listing", "type", "Trigger")
		r, err := c.Service().Accounts.Containers.Workspaces.Triggers.List(c.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		for _, value := range r.Trigger {
			c.triggers[name] = value
		}
	}

	if _, ok := c.triggers[name]; !ok {
		return nil, ErrNotFound
	}
	return c.triggers[name], nil
}

func (c *Client) Tag(name string) (*tagmanager.Tag, error) {
	if c.tags == nil {
		c.tags = map[string]*tagmanager.Tag{}

		c.l.Debug("listing", "type", "Tag")
		r, err := c.Service().Accounts.Containers.Workspaces.Tags.List(c.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		for _, value := range r.Tag {
			c.tags[name] = value
		}
	}

	if _, ok := c.tags[name]; !ok {
		return nil, ErrNotFound
	}

	return c.tags[name], nil
}

func (c *Client) UpsertClient(client *tagmanager.Client) (*tagmanager.Client, error) {
	l := c.l.With("type", "Client", "name", client.Name)

	client.AccountId = c.AccountID()
	client.ContainerId = c.ContainerID()
	client.WorkspaceId = c.WorkspaceID()
	client.Notes = c.Notes()

	cache, err := c.Client(client.Name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	folder, err := c.Folder(c.folderName)
	if err != nil {
		return nil, err
	}
	client.ParentFolderId = folder.FolderId

	if cache == nil {
		l.Info("creating")
		c.clients[client.Name], err = c.Service().Accounts.Containers.Workspaces.Clients.Create(c.WorkspacePath(), client).Do()
	} else {
		l.Info("updating")
		c.clients[client.Name], err = c.Service().Accounts.Containers.Workspaces.Clients.Update(c.WorkspacePath()+"/clients/"+cache.ClientId, client).Do()
	}
	if err != nil {
		return nil, err
	}

	return c.Client(client.Name)
}

func (c *Client) UpsertFolder(name string) (*tagmanager.Folder, error) {
	l := c.l.With("type", "Folder", "name", name)

	cache, err := c.Folder(name)
	if err != nil && !errors.Is(err, ErrNotFound) && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	folder := &tagmanager.Folder{
		AccountId:   c.accountID,
		ContainerId: c.containerID,
		WorkspaceId: c.workspaceID,
		Name:        name,
		Notes:       c.notes,
	}

	if cache == nil {
		l.Info("creating")
		c.folders[name], err = c.Service().Accounts.Containers.Workspaces.Folders.Create(c.WorkspacePath(), folder).Do()
	} else {
		l.Info("updating")
		c.folders[name], err = c.Service().Accounts.Containers.Workspaces.Folders.Update(c.WorkspacePath()+"/folders/"+cache.FolderId, folder).Do()
	}
	if err != nil {
		return nil, err
	}

	return c.Folder(name)
}

func (c *Client) UpsertVariable(variable *tagmanager.Variable) (*tagmanager.Variable, error) {
	l := c.l.With("type", "Variable", "name", variable.Name)

	variable.AccountId = c.AccountID()
	variable.ContainerId = c.ContainerID()
	variable.WorkspaceId = c.WorkspaceID()
	variable.Notes = c.Notes()

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
		l.Info("creating")
		c.variables[variable.Name], err = c.Service().Accounts.Containers.Workspaces.Variables.Create(c.WorkspacePath(), variable).Do()
	} else {
		l.Info("updating")
		c.variables[variable.Name], err = c.Service().Accounts.Containers.Workspaces.Variables.Update(c.WorkspacePath()+"/variables/"+cache.VariableId, variable).Do()
	}
	if err != nil {
		return nil, err
	}

	return c.Variable(variable.Name)
}

func (c *Client) EnableBuiltInVariable(name string) (*tagmanager.Variable, error) {
	l := c.l.With("type", "Built-In Variable", "name", name)

	cache, err := c.BuiltInVariable(name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	if cache != nil {
		return cache, nil
	}

	l.Info("creating")
	resp, err := c.Service().Accounts.Containers.Workspaces.BuiltInVariables.Create(c.WorkspacePath()).Type(name).Do()
	if err != nil {
		return nil, err
	}
	for _, builtInVariable := range resp.BuiltInVariable {
		c.builtInVariables[builtInVariable.Name] = builtInVariable
	}

	return c.BuiltInVariable(name)
}

func (c *Client) UpsertTrigger(trigger *tagmanager.Trigger) (*tagmanager.Trigger, error) {
	l := c.l.With("type", "Trigger", "name", trigger.Name)

	trigger.AccountId = c.AccountID()
	trigger.ContainerId = c.ContainerID()
	trigger.WorkspaceId = c.WorkspaceID()
	trigger.Notes = c.Notes()

	cache, err := c.Trigger(trigger.Name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	folder, err := c.Folder(c.folderName)
	if err != nil {
		return nil, err
	}
	trigger.ParentFolderId = folder.FolderId

	if cache == nil {
		l.Info("creating")
		c.triggers[trigger.Name], err = c.Service().Accounts.Containers.Workspaces.Triggers.Create(c.WorkspacePath(), trigger).Do()
	} else {
		l.Info("updating")
		c.triggers[trigger.Name], err = c.Service().Accounts.Containers.Workspaces.Triggers.Update(c.WorkspacePath()+"/triggers/"+cache.TriggerId, trigger).Do()
	}
	if err != nil {
		return nil, err
	}

	return c.Trigger(trigger.Name)
}

func (c *Client) UpsertTag(tag *tagmanager.Tag) (*tagmanager.Tag, error) {
	l := c.l.With("type", "Tag", "name", tag.Name)

	tag.AccountId = c.AccountID()
	tag.ContainerId = c.ContainerID()
	tag.WorkspaceId = c.WorkspaceID()
	tag.Notes = c.Notes()

	cache, err := c.Tag(tag.Name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	folder, err := c.Folder(c.folderName)
	if err != nil {
		return nil, err
	}
	tag.ParentFolderId = folder.FolderId

	if cache == nil {
		l.Info("creating")
		c.tags[tag.Name], err = c.Service().Accounts.Containers.Workspaces.Tags.Create(c.WorkspacePath(), tag).Do()
	} else {
		l.Info("updating")
		c.tags[tag.Name], err = c.Service().Accounts.Containers.Workspaces.Tags.Update(c.WorkspacePath()+"/tags/"+cache.TagId, tag).Do()
	}
	if err != nil {
		return nil, err
	}

	return c.Tag(tag.Name)
}
