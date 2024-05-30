package tagmanager

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/mitchellh/hashstructure/v2"
	"github.com/pkg/errors"
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
		if v > 0 {
			o.requestThrottler = time.NewTicker((100 * time.Second) / time.Duration(v))
		}
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
		l:                l,
		accountID:        accountID,
		containerID:      containerID,
		workspaceID:      workspaceID,
		measurementID:    measurementID,
		requestThrottler: time.NewTicker((100 * time.Second) / time.Duration(15)),
		notes:            "Managed by Sesamy. DO NOT EDIT.",
		folderName:       "Sesamy",
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

func (c *Client) Notes(v any) string {
	var hash string
	if v != nil {
		if value, err := hashstructure.Hash(v, hashstructure.FormatV2, nil); err != nil {
			c.l.Warn("failed to hash struct:", "error", err)
		} else {
			hash = fmt.Sprintf(" [%d]", value)
		}
	}
	return c.notes + hash
}

// ------------------------------------------------------------------------------------------------
// ~ Public methods
// ------------------------------------------------------------------------------------------------

func (c *Client) GetClient(name string) (*tagmanager.Client, error) {
	elems, err := c.LoadClients()
	if err != nil {
		return nil, err
	}

	if _, ok := elems[name]; !ok {
		return nil, ErrNotFound
	}

	return elems[name], nil
}

func (c *Client) LoadClients() (map[string]*tagmanager.Client, error) {
	if c.clients == nil {
		c.l.Debug("loading", "type", "Client")
		r, err := c.Service().Accounts.Containers.Workspaces.Clients.List(c.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		res := map[string]*tagmanager.Client{}
		for _, value := range r.Client {
			res[value.Name] = value
		}
		c.clients = res
	}

	return c.clients, nil
}

func (c *Client) GetFolder(name string) (*tagmanager.Folder, error) {
	elems, err := c.LoadFolders()
	if err != nil {
		return nil, err
	}

	if _, ok := elems[name]; !ok {
		return nil, ErrNotFound
	}

	return elems[name], nil
}

func (c *Client) LoadFolders() (map[string]*tagmanager.Folder, error) {
	if c.folders == nil {
		c.l.Debug("loading", "type", "Folder")
		r, err := c.Service().Accounts.Containers.Workspaces.Folders.List(c.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		res := map[string]*tagmanager.Folder{}
		for _, value := range r.Folder {
			res[value.Name] = value
		}
		c.folders = res
	}
	return c.folders, nil
}

func (c *Client) GetVariable(name string) (*tagmanager.Variable, error) {
	elems, err := c.LoadVariables()
	if err != nil {
		return nil, err
	}

	if _, ok := elems[name]; !ok {
		return nil, ErrNotFound
	}

	return elems[name], nil
}

func (c *Client) LoadVariables() (map[string]*tagmanager.Variable, error) {
	if c.variables == nil {
		c.l.Debug("loading", "type", "Variable")
		r, err := c.Service().Accounts.Containers.Workspaces.Variables.List(c.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		res := map[string]*tagmanager.Variable{}
		for _, value := range r.Variable {
			res[value.Name] = value
		}
		c.variables = res
	}

	return c.variables, nil
}

func (c *Client) GetBuiltInVariable(typeName string) (*tagmanager.BuiltInVariable, error) {
	elems, err := c.LoadBuiltInVariables()
	if err != nil {
		return nil, err
	}

	if _, ok := elems[typeName]; !ok {
		return nil, ErrNotFound
	}

	return elems[typeName], nil
}

func (c *Client) LoadBuiltInVariables() (map[string]*tagmanager.BuiltInVariable, error) {
	if c.builtInVariables == nil {
		c.l.Debug("loading", "type", "BuiltInVariable")
		r, err := c.Service().Accounts.Containers.Workspaces.BuiltInVariables.List(c.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		res := map[string]*tagmanager.BuiltInVariable{}
		for _, value := range r.BuiltInVariable {
			res[value.Type] = value
		}
		c.builtInVariables = res
	}

	return c.builtInVariables, nil
}

func (c *Client) Trigger(name string) (*tagmanager.Trigger, error) {
	elems, err := c.LoadTriggers()
	if err != nil {
		return nil, err
	}

	if _, ok := elems[name]; !ok {
		return nil, ErrNotFound
	}

	return elems[name], nil
}

func (c *Client) LoadTriggers() (map[string]*tagmanager.Trigger, error) {
	if c.triggers == nil {
		c.l.Debug("loading", "type", "Trigger")
		r, err := c.Service().Accounts.Containers.Workspaces.Triggers.List(c.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		res := map[string]*tagmanager.Trigger{}
		for _, value := range r.Trigger {
			res[value.Name] = value
		}
		c.triggers = res
	}

	return c.triggers, nil
}

func (c *Client) Tag(name string) (*tagmanager.Tag, error) {
	elems, err := c.LoadTags()
	if err != nil {
		return nil, err
	}

	if _, ok := elems[name]; !ok {
		return nil, ErrNotFound
	}

	return elems[name], nil
}

func (c *Client) LoadTags() (map[string]*tagmanager.Tag, error) {
	if c.tags == nil {
		c.l.Debug("loading", "type", "Tag")
		r, err := c.Service().Accounts.Containers.Workspaces.Tags.List(c.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		res := map[string]*tagmanager.Tag{}
		for _, value := range r.Tag {
			res[value.Name] = value
		}
		c.tags = res
	}

	return c.tags, nil
}

func (c *Client) UpsertClient(item *tagmanager.Client) (*tagmanager.Client, error) {
	l := c.l.With("type", "Client", "name", item.Name)

	item.Notes = c.Notes(item)
	item.AccountId = c.AccountID()
	item.ContainerId = c.ContainerID()
	item.WorkspaceId = c.WorkspaceID()

	cache, err := c.GetClient(item.Name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	folder, err := c.GetFolder(c.folderName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve folder")
	}
	item.ParentFolderId = folder.FolderId

	if cache == nil {
		l.Info("⚠️ creating")
		c.clients[item.Name], err = c.Service().Accounts.Containers.Workspaces.Clients.Create(c.WorkspacePath(), item).Do()
	} else if item.Notes == cache.Notes {
		l.Info("✅ unchanged", "id", cache.ClientId)
	} else {
		l.Info("⚠️ updating", "id", cache.ClientId)
		c.clients[item.Name], err = c.Service().Accounts.Containers.Workspaces.Clients.Update(c.WorkspacePath()+"/clients/"+cache.ClientId, item).Do()
	}
	if err != nil {
		return nil, err
	}

	return c.GetClient(item.Name)
}

func (c *Client) UpsertFolder(name string) (*tagmanager.Folder, error) {
	l := c.l.With("type", "Folder", "name", name)

	item := &tagmanager.Folder{
		Name: name,
	}

	item.Notes = c.Notes(item)
	item.AccountId = c.AccountID()
	item.ContainerId = c.ContainerID()
	item.WorkspaceId = c.WorkspaceID()

	cache, err := c.GetFolder(name)
	if err != nil && !errors.Is(err, ErrNotFound) && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	if cache == nil {
		l.Info("⚠️ creating")
		c.folders[name], err = c.Service().Accounts.Containers.Workspaces.Folders.Create(c.WorkspacePath(), item).Do()
	} else if item.Notes == cache.Notes {
		l.Info("✅ unchanged", "id", item.FolderId)
	} else {
		l.Info("⚠️ updating", "id", cache.FolderId)
		c.folders[name], err = c.Service().Accounts.Containers.Workspaces.Folders.Update(c.WorkspacePath()+"/folders/"+cache.FolderId, item).Do()
	}
	if err != nil {
		return nil, err
	}

	return c.GetFolder(name)
}

func (c *Client) UpsertVariable(item *tagmanager.Variable) (*tagmanager.Variable, error) {
	l := c.l.With("type", "Variable", "name", item.Name)

	item.Notes = c.Notes(item)
	item.AccountId = c.AccountID()
	item.ContainerId = c.ContainerID()
	item.WorkspaceId = c.WorkspaceID()

	cache, err := c.GetVariable(item.Name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	folder, err := c.GetFolder(c.folderName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve folder")
	}
	item.ParentFolderId = folder.FolderId

	if cache == nil {
		l.Info("⚠️ creating")
		c.variables[item.Name], err = c.Service().Accounts.Containers.Workspaces.Variables.Create(c.WorkspacePath(), item).Do()
	} else if item.Notes == cache.Notes {
		l.Info("✅ unchanged", "id", cache.VariableId)
	} else {
		l.Info("⚠️ updating", "id", cache.VariableId)
		c.variables[item.Name], err = c.Service().Accounts.Containers.Workspaces.Variables.Update(c.WorkspacePath()+"/variables/"+cache.VariableId, item).Do()
	}
	if err != nil {
		return nil, err
	}

	return c.GetVariable(item.Name)
}

func (c *Client) EnableBuiltInVariable(typeName string) (*tagmanager.BuiltInVariable, error) {
	l := c.l.With("type", "Built-In Variable", "typeName", typeName)

	cache, err := c.GetBuiltInVariable(typeName)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	if cache != nil {
		l.Info("✅ unchanged")
		return cache, nil
	}

	l.Info("⚠️ creating")
	resp, err := c.Service().Accounts.Containers.Workspaces.BuiltInVariables.Create(c.WorkspacePath()).Type(typeName).Do()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create built-in variable")
	}

	for _, builtInVariable := range resp.BuiltInVariable {
		c.builtInVariables[builtInVariable.Type] = builtInVariable
	}

	return c.GetBuiltInVariable(typeName)
}

func (c *Client) UpsertTrigger(item *tagmanager.Trigger) (*tagmanager.Trigger, error) {
	l := c.l.With("type", "Trigger", "name", item.Name)

	item.Notes = c.Notes(item)
	item.AccountId = c.AccountID()
	item.ContainerId = c.ContainerID()
	item.WorkspaceId = c.WorkspaceID()

	cache, err := c.Trigger(item.Name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	folder, err := c.GetFolder(c.folderName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve folder")
	}
	item.ParentFolderId = folder.FolderId

	if cache == nil {
		l.Info("⚠️ creating")
		c.triggers[item.Name], err = c.Service().Accounts.Containers.Workspaces.Triggers.Create(c.WorkspacePath(), item).Do()
	} else if item.Notes == cache.Notes {
		l.Info("✅ unchanged", "id", cache.TriggerId)
	} else {
		l.Info("⚠️ updating", "id", cache.TriggerId)
		c.triggers[item.Name], err = c.Service().Accounts.Containers.Workspaces.Triggers.Update(c.WorkspacePath()+"/triggers/"+cache.TriggerId, item).Do()
	}
	if err != nil {
		return nil, err
	}

	return c.Trigger(item.Name)
}

func (c *Client) UpsertTag(item *tagmanager.Tag) (*tagmanager.Tag, error) {
	l := c.l.With("type", "Tag", "name", item.Name)

	item.Notes = c.Notes(item)
	item.AccountId = c.AccountID()
	item.ContainerId = c.ContainerID()
	item.WorkspaceId = c.WorkspaceID()

	cache, err := c.Tag(item.Name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	folder, err := c.GetFolder(c.folderName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve folder")
	}
	item.ParentFolderId = folder.FolderId

	if cache == nil {
		l.Info("⚠️ creating")
		c.tags[item.Name], err = c.Service().Accounts.Containers.Workspaces.Tags.Create(c.WorkspacePath(), item).Do()
	} else if item.Notes == cache.Notes {
		l.Info("✅ unchanged", "id", cache.TagId)
	} else {
		l.Info("⚠️ updating", "id", cache.TagId)
		c.tags[item.Name], err = c.Service().Accounts.Containers.Workspaces.Tags.Update(c.WorkspacePath()+"/tags/"+cache.TagId, item).Do()
	}
	if err != nil {
		return nil, err
	}

	return c.Tag(item.Name)
}
