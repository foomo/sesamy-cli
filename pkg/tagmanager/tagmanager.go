package tagmanager

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
	"google.golang.org/api/tagmanager/v2"
)

type (
	TagManager struct {
		l                *slog.Logger
		notes            string
		accountID        string
		container        config.GoogleTagManagerContainer
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
		customTemplates  map[string]*tagmanager.CustomTemplate
		transformations  map[string]*tagmanager.Transformation
	}
	Option func(*TagManager)
)

// ------------------------------------------------------------------------------------------------
// ~ Options
// ------------------------------------------------------------------------------------------------

func WithNotes(v string) Option {
	return func(o *TagManager) {
		o.notes = v
	}
}

func WithFolderName(v string) Option {
	return func(o *TagManager) {
		o.folderName = v
	}
}

func WithRequestQuota(v int) Option {
	return func(o *TagManager) {
		if v > 0 {
			o.requestThrottler = time.NewTicker((100 * time.Second) / time.Duration(v))
		}
	}
}

func WithClientOptions(v ...option.ClientOption) Option {
	return func(o *TagManager) {
		o.clientOptions = append(o.clientOptions, v...)
	}
}

// ------------------------------------------------------------------------------------------------
// ~ Constructor
// ------------------------------------------------------------------------------------------------

func New(ctx context.Context, l *slog.Logger, accountID string, container config.GoogleTagManagerContainer, opts ...Option) (*TagManager, error) {
	inst := &TagManager{
		l:                l,
		accountID:        accountID,
		container:        container,
		requestThrottler: time.NewTicker((100 * time.Second) / time.Duration(15)),
		notes:            "Managed by Sesamy. DO NOT EDIT.",
		folderName:       "Sesamy",
		clientOptions: []option.ClientOption{
			option.WithLogger(l),
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

func (t *TagManager) AccountID() string {
	return t.accountID
}

func (t *TagManager) ContainerID() string {
	return t.container.ContainerID
}

func (t *TagManager) WorkspaceID() string {
	return t.container.WorkspaceID
}

func (t *TagManager) FolderName() string {
	return t.folderName
}

func (t *TagManager) SetFolderName(v string) {
	t.folderName = v
}

func (t *TagManager) Service() *tagmanager.Service {
	if t.requestThrottler != nil {
		<-t.requestThrottler.C
	}
	return t.service
}

func (t *TagManager) AccountPath() string {
	return "accounts/" + t.accountID
}

func (t *TagManager) ContainerPath() string {
	return t.AccountPath() + "/containers/" + t.ContainerID()
}

func (t *TagManager) WorkspacePath() string {
	return t.ContainerPath() + "/workspaces/" + t.WorkspaceID()
}

func (t *TagManager) Notes(v any) string {
	var hash string
	if v != nil {
		if value, err := hashstructure.Hash(v, hashstructure.FormatV2, nil); err != nil {
			t.l.Warn("failed to hash struct:", "error", err)
		} else {
			hash = fmt.Sprintf(" [%d]", value)
		}
	}
	return t.notes + hash
}

// ------------------------------------------------------------------------------------------------
// ~ Public methods
// ------------------------------------------------------------------------------------------------

func (t *TagManager) LookupClient(name string) (*tagmanager.Client, error) {
	elems, err := t.LoadClients()
	if err != nil {
		return nil, err
	}

	if _, ok := elems[name]; !ok {
		return nil, ErrNotFound
	}

	return elems[name], nil
}

func (t *TagManager) LoadClients() (map[string]*tagmanager.Client, error) {
	if t.clients == nil {
		t.l.Info("🛄 Loading list", "type", "Client")
		r, err := t.Service().Accounts.Containers.Workspaces.Clients.List(t.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		res := map[string]*tagmanager.Client{}
		for _, value := range r.Client {
			res[value.Name] = value
		}
		t.clients = res
	}

	return t.clients, nil
}

func (t *TagManager) LookupFolder(name string) (*tagmanager.Folder, error) {
	elems, err := t.LoadFolders()
	if err != nil {
		return nil, err
	}

	if _, ok := elems[name]; !ok {
		return nil, ErrNotFound
	}

	return elems[name], nil
}

func (t *TagManager) LoadFolders() (map[string]*tagmanager.Folder, error) {
	if t.folders == nil {
		t.l.Info("🛄 Loading list", "type", "Folder")
		r, err := t.Service().Accounts.Containers.Workspaces.Folders.List(t.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		res := map[string]*tagmanager.Folder{}
		for _, value := range r.Folder {
			res[value.Name] = value
		}
		t.folders = res
	}
	return t.folders, nil
}

func (t *TagManager) LookupVariable(name string) (*tagmanager.Variable, error) {
	elems, err := t.LoadVariables()
	if err != nil {
		return nil, err
	}

	if _, ok := elems[name]; !ok {
		return nil, ErrNotFound
	}

	return elems[name], nil
}

func (t *TagManager) LoadVariables() (map[string]*tagmanager.Variable, error) {
	if t.variables == nil {
		t.l.Info("🛄 Loading list", "type", "Variable")
		r, err := t.Service().Accounts.Containers.Workspaces.Variables.List(t.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		res := map[string]*tagmanager.Variable{}
		for _, value := range r.Variable {
			res[value.Name] = value
		}
		t.variables = res
	}

	return t.variables, nil
}

func (t *TagManager) GetBuiltInVariable(typeName string) (*tagmanager.BuiltInVariable, error) {
	elems, err := t.LoadBuiltInVariables()
	if err != nil {
		return nil, err
	}

	if _, ok := elems[typeName]; !ok {
		return nil, ErrNotFound
	}

	return elems[typeName], nil
}

func (t *TagManager) LoadBuiltInVariables() (map[string]*tagmanager.BuiltInVariable, error) {
	if t.builtInVariables == nil {
		t.l.Info("🛄 Loading list", "type", "BuiltInVariable")
		r, err := t.Service().Accounts.Containers.Workspaces.BuiltInVariables.List(t.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		res := map[string]*tagmanager.BuiltInVariable{}
		for _, value := range r.BuiltInVariable {
			res[value.Type] = value
		}
		t.builtInVariables = res
	}

	return t.builtInVariables, nil
}

func (t *TagManager) LookupTrigger(name string) (*tagmanager.Trigger, error) {
	elems, err := t.LoadTriggers()
	if err != nil {
		return nil, err
	}

	if _, ok := elems[name]; !ok {
		return nil, ErrNotFound
	}

	return elems[name], nil
}

func (t *TagManager) LookupTemplate(name string) (*tagmanager.CustomTemplate, error) {
	elems, err := t.LoadCustomTemplates()
	if err != nil {
		return nil, err
	}

	if _, ok := elems[name]; !ok {
		return nil, ErrNotFound
	}

	return elems[name], nil
}

func (t *TagManager) LookupTransformation(name string) (*tagmanager.Transformation, error) {
	elems, err := t.LoadTransformations()
	if err != nil {
		return nil, err
	}

	if _, ok := elems[name]; !ok {
		return nil, ErrNotFound
	}

	return elems[name], nil
}

func (t *TagManager) LoadTriggers() (map[string]*tagmanager.Trigger, error) {
	if t.triggers == nil {
		t.l.Info("🛄 Loading list", "type", "Trigger")
		r, err := t.Service().Accounts.Containers.Workspaces.Triggers.List(t.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		res := map[string]*tagmanager.Trigger{}
		for _, value := range r.Trigger {
			res[value.Name] = value
		}
		t.triggers = res
	}

	return t.triggers, nil
}

func (t *TagManager) LookupTag(name string) (*tagmanager.Tag, error) {
	elems, err := t.LoadTags()
	if err != nil {
		return nil, err
	}

	if _, ok := elems[name]; !ok {
		return nil, ErrNotFound
	}

	return elems[name], nil
}

func (t *TagManager) LoadTags() (map[string]*tagmanager.Tag, error) {
	if t.tags == nil {
		t.l.Info("🛄 Loading list", "type", "Tag")
		r, err := t.Service().Accounts.Containers.Workspaces.Tags.List(t.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		res := map[string]*tagmanager.Tag{}
		for _, value := range r.Tag {
			res[value.Name] = value
		}
		t.tags = res
	}

	return t.tags, nil
}

func (t *TagManager) CustomTemplate(name string) (*tagmanager.CustomTemplate, error) {
	elems, err := t.LoadCustomTemplates()
	if err != nil {
		return nil, err
	}

	if _, ok := elems[name]; !ok {
		return nil, ErrNotFound
	}

	return elems[name], nil
}

func (t *TagManager) LoadCustomTemplates() (map[string]*tagmanager.CustomTemplate, error) {
	if t.customTemplates == nil {
		t.l.Info("🛄 Loading list", "type", "CustomTemplate")
		r, err := t.Service().Accounts.Containers.Workspaces.Templates.List(t.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		res := map[string]*tagmanager.CustomTemplate{}
		for _, value := range r.Template {
			res[value.Name] = value
		}
		t.customTemplates = res
	}

	return t.customTemplates, nil
}

func (t *TagManager) LoadTransformations() (map[string]*tagmanager.Transformation, error) {
	if t.transformations == nil {
		t.l.Info("🛄 Loading list", "type", "Transformation")
		r, err := t.Service().Accounts.Containers.Workspaces.Transformations.List(t.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		res := map[string]*tagmanager.Transformation{}
		for _, value := range r.Transformation {
			res[value.Name] = value
		}
		t.transformations = res
	}

	return t.transformations, nil
}

func (t *TagManager) UpsertClient(item *tagmanager.Client) (*tagmanager.Client, error) {
	l := t.l.With("type", "Client", "name", item.Name)

	folder, err := t.LookupFolder(t.folderName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve folder")
	}
	item.ParentFolderId = folder.FolderId

	item.Notes = t.Notes(item)
	item.AccountId = t.AccountID()
	item.ContainerId = t.ContainerID()
	item.WorkspaceId = t.WorkspaceID()

	cache, err := t.LookupClient(item.Name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	if cache == nil {
		l.Info("🚀 New")
		t.clients[item.Name], err = t.Service().Accounts.Containers.Workspaces.Clients.Create(t.WorkspacePath(), item).Do()
	} else if item.Notes == cache.Notes {
		l.Info("✅ OK", "id", cache.ClientId)
	} else {
		l.Info("🔄 Update", "id", cache.ClientId)
		t.clients[item.Name], err = t.Service().Accounts.Containers.Workspaces.Clients.Update(t.WorkspacePath()+"/clients/"+cache.ClientId, item).Do()
	}
	if err != nil {
		return nil, err
	}

	return t.LookupClient(item.Name)
}

func (t *TagManager) UpsertTransformation(item *tagmanager.Transformation) (*tagmanager.Transformation, error) {
	l := t.l.With("type", "Transformation", "name", item.Name)

	folder, err := t.LookupFolder(t.folderName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve folder")
	}
	item.ParentFolderId = folder.FolderId

	item.Notes = t.Notes(item)
	item.AccountId = t.AccountID()
	item.ContainerId = t.ContainerID()
	item.WorkspaceId = t.WorkspaceID()

	cache, err := t.LookupTransformation(item.Name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	if cache == nil {
		l.Info("🚀 New")
		t.transformations[item.Name], err = t.Service().Accounts.Containers.Workspaces.Transformations.Create(t.WorkspacePath(), item).Do()
	} else if item.Notes == cache.Notes {
		l.Info("✅ OK", "id", cache.TransformationId)
	} else {
		l.Info("🔄 Update", "id", cache.TransformationId)
		t.transformations[item.Name], err = t.Service().Accounts.Containers.Workspaces.Transformations.Update(t.WorkspacePath()+"/transformations/"+cache.TransformationId, item).Do()
	}
	if err != nil {
		return nil, err
	}

	return t.LookupTransformation(item.Name)
}

func (t *TagManager) UpsertFolder(name string) (*tagmanager.Folder, error) {
	l := t.l.With("type", "Folder", "name", name)

	item := &tagmanager.Folder{
		Name: name,
	}

	item.Notes = t.Notes(item)
	item.AccountId = t.AccountID()
	item.ContainerId = t.ContainerID()
	item.WorkspaceId = t.WorkspaceID()

	cache, err := t.LookupFolder(name)
	if err != nil && !errors.Is(err, ErrNotFound) && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	if cache == nil {
		l.Info("🚀 New")
		t.folders[name], err = t.Service().Accounts.Containers.Workspaces.Folders.Create(t.WorkspacePath(), item).Do()
	} else if item.Notes == cache.Notes {
		l.Info("✅ OK", "id", item.FolderId)
	} else {
		l.Info("🔄 Update", "id", cache.FolderId)
		t.folders[name], err = t.Service().Accounts.Containers.Workspaces.Folders.Update(t.WorkspacePath()+"/folders/"+cache.FolderId, item).Do()
	}
	if err != nil {
		return nil, err
	}

	return t.LookupFolder(name)
}

func (t *TagManager) UpsertVariable(item *tagmanager.Variable) (*tagmanager.Variable, error) {
	l := t.l.With("type", "Variable", "name", item.Name)

	folder, err := t.LookupFolder(t.folderName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve folder")
	}
	item.ParentFolderId = folder.FolderId

	item.Notes = t.Notes(item)
	item.AccountId = t.AccountID()
	item.ContainerId = t.ContainerID()
	item.WorkspaceId = t.WorkspaceID()

	cache, err := t.LookupVariable(item.Name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	if cache == nil {
		l.Info("🚀 New")
		t.variables[item.Name], err = t.Service().Accounts.Containers.Workspaces.Variables.Create(t.WorkspacePath(), item).Do()
	} else if item.Notes == cache.Notes {
		l.Info("✅ OK", "id", cache.VariableId)
	} else {
		l.Info("🔄 Update", "id", cache.VariableId)
		t.variables[item.Name], err = t.Service().Accounts.Containers.Workspaces.Variables.Update(t.WorkspacePath()+"/variables/"+cache.VariableId, item).Do()
	}
	if err != nil {
		return nil, err
	}

	return t.LookupVariable(item.Name)
}

func (t *TagManager) EnableBuiltInVariable(typeName string) (*tagmanager.BuiltInVariable, error) {
	l := t.l.With("type", "Built-In Variable", "typeName", typeName)

	cache, err := t.GetBuiltInVariable(typeName)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	if cache != nil {
		l.Info("✅ OK")
		return cache, nil
	}

	l.Info("🚀 New")
	resp, err := t.Service().Accounts.Containers.Workspaces.BuiltInVariables.Create(t.WorkspacePath()).Type(typeName).Do()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create built-in variable")
	}

	for _, builtInVariable := range resp.BuiltInVariable {
		t.builtInVariables[builtInVariable.Type] = builtInVariable
	}

	return t.GetBuiltInVariable(typeName)
}

func (t *TagManager) UpsertTrigger(item *tagmanager.Trigger) (*tagmanager.Trigger, error) {
	l := t.l.With("type", "Trigger", "name", item.Name)

	folder, err := t.LookupFolder(t.folderName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve folder")
	}
	item.ParentFolderId = folder.FolderId

	item.Notes = t.Notes(item)
	item.AccountId = t.AccountID()
	item.ContainerId = t.ContainerID()
	item.WorkspaceId = t.WorkspaceID()

	cache, err := t.LookupTrigger(item.Name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	if cache == nil {
		l.Info("🚀 New")
		t.triggers[item.Name], err = t.Service().Accounts.Containers.Workspaces.Triggers.Create(t.WorkspacePath(), item).Do()
	} else if item.Notes == cache.Notes {
		l.Info("✅ OK", "id", cache.TriggerId)
	} else {
		l.Info("🔄 Update", "id", cache.TriggerId)
		t.triggers[item.Name], err = t.Service().Accounts.Containers.Workspaces.Triggers.Update(t.WorkspacePath()+"/triggers/"+cache.TriggerId, item).Do()
	}
	if err != nil {
		return nil, err
	}

	return t.LookupTrigger(item.Name)
}

func (t *TagManager) UpsertTag(item *tagmanager.Tag) (*tagmanager.Tag, error) {
	l := t.l.With("type", "Tag", "name", item.Name)

	folder, err := t.LookupFolder(t.folderName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve folder")
	}
	item.ParentFolderId = folder.FolderId

	item.Notes = t.Notes(item)
	item.AccountId = t.AccountID()
	item.ContainerId = t.ContainerID()
	item.WorkspaceId = t.WorkspaceID()

	cache, err := t.LookupTag(item.Name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	if cache == nil {
		l.Info("🚀 New")
		t.tags[item.Name], err = t.Service().Accounts.Containers.Workspaces.Tags.Create(t.WorkspacePath(), item).Do()
	} else if item.Notes == cache.Notes {
		l.Info("✅ OK", "id", cache.TagId)
	} else {
		l.Info("🔄 Update", "id", cache.TagId)
		t.tags[item.Name], err = t.Service().Accounts.Containers.Workspaces.Tags.Update(t.WorkspacePath()+"/tags/"+cache.TagId, item).Do()
	}
	if err != nil {
		return nil, err
	}

	return t.LookupTag(item.Name)
}

func (t *TagManager) UpsertCustomTemplate(item *tagmanager.CustomTemplate) (*tagmanager.CustomTemplate, error) {
	l := t.l.With("type", "CustomTemplate", "name", item.Name)

	item.AccountId = t.AccountID()
	item.ContainerId = t.ContainerID()
	item.WorkspaceId = t.WorkspaceID()

	cache, err := t.CustomTemplate(item.Name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	if cache == nil {
		l.Info("🚀 New")
		t.customTemplates[item.Name], err = t.Service().Accounts.Containers.Workspaces.Templates.Create(t.WorkspacePath(), item).Do()
	} else if item.TemplateData == cache.TemplateData {
		l.Info("✅ OK", "id", cache.TemplateId)
	} else {
		l.Info("🔄 Update", "id", cache.TemplateId)
		t.customTemplates[item.Name], err = t.Service().Accounts.Containers.Workspaces.Templates.Update(t.WorkspacePath()+"/templates/"+cache.TemplateId, item).Do()
	}
	if err != nil {
		return nil, err
	}

	return t.CustomTemplate(item.Name)
}
