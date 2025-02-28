package tagmanager

import (
	"context"
	"crypto/md5" //nolint: gosec //just a checksum
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
	"google.golang.org/api/tagmanager/v2"
)

type (
	TagManager struct {
		l         *slog.Logger
		notes     string
		accountID string
		container config.GoogleTagManagerContainer
		// folderName       string
		clientOptions    []option.ClientOption
		requestThrottler *time.Ticker
		// cache
		service          *tagmanager.Service
		clients          *AccessedMap[*tagmanager.Client]
		folders          *AccessedMap[*tagmanager.Folder]
		variables        *AccessedMap[*tagmanager.Variable]
		builtInVariables *AccessedMap[*tagmanager.BuiltInVariable]
		triggers         *AccessedMap[*tagmanager.Trigger]
		tags             *AccessedMap[*tagmanager.Tag]
		customTemplates  *AccessedMap[*tagmanager.CustomTemplate]
		transformations  *AccessedMap[*tagmanager.Transformation]
	}
	Option func(*TagManager)
)

type AccessedMap[T any] struct {
	data map[string]T
	keys map[string]bool
}

func NewAccessedMap[T any](data map[string]T) *AccessedMap[T] {
	return &AccessedMap[T]{
		data: data,
		keys: make(map[string]bool, len(data)),
	}
}

func (l AccessedMap[T]) Has(key string) bool {
	_, ok := l.data[key]
	return ok
}

func (l AccessedMap[T]) Get(key string) T {
	if l.Has(key) {
		l.keys[key] = true
	}
	return l.data[key]
}

func (l AccessedMap[T]) Set(key string, value T) {
	l.keys[key] = true
	l.data[key] = value
}

func (l AccessedMap[T]) Misssed() map[string]T {
	ret := map[string]T{}
	for k := range l.data {
		if !l.keys[k] {
			ret[k] = l.data[k]
		}
	}
	return ret
}

// ------------------------------------------------------------------------------------------------
// ~ Options
// ------------------------------------------------------------------------------------------------

func WithNotes(v string) Option {
	return func(o *TagManager) {
		o.notes = v
	}
}

// func WithFolderName(v string) Option {
// 	return func(o *TagManager) {
// 		o.folderName = v
// 	}
// }

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
		notes:            "DO NOT EDIT!\n\nManaged by Sesamy",
		// folderName:       "Sesamy",
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

func (t *TagManager) Missed() map[string][]string {
	ret := map[string][]string{}
	if t.clients != nil {
		for _, i2 := range t.clients.Misssed() {
			ret["Clients"] = append(ret["Clients"], i2.Name)
		}
	}
	if t.folders != nil {
		for _, i2 := range t.folders.Misssed() {
			ret["Folders"] = append(ret["Folders"], i2.Name)
		}
	}
	if t.variables != nil {
		for _, i2 := range t.variables.Misssed() {
			ret["Variables"] = append(ret["Variables"], i2.Name)
		}
	}
	if t.builtInVariables != nil {
		for _, i2 := range t.builtInVariables.Misssed() {
			ret["Built In Variables"] = append(ret["Built In Variables"], i2.Name)
		}
	}
	if t.triggers != nil {
		for _, i2 := range t.triggers.Misssed() {
			ret["Triggers"] = append(ret["Triggers"], i2.Name)
		}
	}
	if t.tags != nil {
		for _, i2 := range t.tags.Misssed() {
			ret["Tags"] = append(ret["Tags"], i2.Name)
		}
	}
	if t.customTemplates != nil {
		for _, i2 := range t.customTemplates.Misssed() {
			ret["Custom Templates"] = append(ret["Custom Templates"], i2.Name)
		}
	}
	if t.transformations != nil {
		for _, i2 := range t.transformations.Misssed() {
			ret["Transformations"] = append(ret["Transformations"], i2.Name)
		}
	}
	return ret
}

func (t *TagManager) AccountID() string {
	return t.accountID
}

func (t *TagManager) ContainerID() string {
	return t.container.ContainerID
}

func (t *TagManager) WorkspaceID() string {
	return t.container.WorkspaceID
}

// func (t *TagManager) FolderName() string {
// 	return t.folderName
// }

// func (t *TagManager) SetFolderName(v string) {
// 	t.folderName = v
// }

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
		if out, err := json.Marshal(v); err != nil {
			t.l.Warn("failed to marshal tag manager", "error", err)
		} else {
			hash = fmt.Sprintf(" - %x", md5.Sum(out)) //nolint: gosec //just a checksum
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

	if !elems.Has(name) {
		return nil, ErrNotFound
	}

	return elems.Get(name), nil
}

func (t *TagManager) LoadClients() (*AccessedMap[*tagmanager.Client], error) {
	if t.clients == nil {
		t.l.Info("└  ⬇︎ Loading list", "type", "Client")
		r, err := t.Service().Accounts.Containers.Workspaces.Clients.List(t.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		res := map[string]*tagmanager.Client{}
		for _, value := range r.Client {
			res[value.Name] = value
		}
		t.clients = NewAccessedMap(res)
	}

	return t.clients, nil
}

func (t *TagManager) LookupFolder(name string) (*tagmanager.Folder, error) {
	elems, err := t.LoadFolders()
	if err != nil {
		return nil, err
	}

	if !elems.Has(name) {
		return nil, ErrNotFound
	}

	return elems.Get(name), nil
}

func (t *TagManager) LoadFolders() (*AccessedMap[*tagmanager.Folder], error) {
	if t.folders == nil {
		t.l.Info("└  ⬇︎ Loading list", "type", "Folder")
		r, err := t.Service().Accounts.Containers.Workspaces.Folders.List(t.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		res := map[string]*tagmanager.Folder{}
		for _, value := range r.Folder {
			res[value.Name] = value
		}
		t.folders = NewAccessedMap(res)
	}
	return t.folders, nil
}

func (t *TagManager) LookupVariable(name string) (*tagmanager.Variable, error) {
	elems, err := t.LoadVariables()
	if err != nil {
		return nil, err
	}

	if !elems.Has(name) {
		return nil, ErrNotFound
	}

	return elems.Get(name), nil
}

func (t *TagManager) LoadVariables() (*AccessedMap[*tagmanager.Variable], error) {
	if t.variables == nil {
		t.l.Info("└  ⬇︎ Loading list", "type", "Variable")
		r, err := t.Service().Accounts.Containers.Workspaces.Variables.List(t.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		res := map[string]*tagmanager.Variable{}
		for _, value := range r.Variable {
			res[value.Name] = value
		}
		t.variables = NewAccessedMap(res)
	}

	return t.variables, nil
}

func (t *TagManager) GetBuiltInVariable(typeName string) (*tagmanager.BuiltInVariable, error) {
	elems, err := t.LoadBuiltInVariables()
	if err != nil {
		return nil, err
	}

	if !elems.Has(typeName) {
		return nil, ErrNotFound
	}

	return elems.Get(typeName), nil
}

func (t *TagManager) LoadBuiltInVariables() (*AccessedMap[*tagmanager.BuiltInVariable], error) {
	if t.builtInVariables == nil {
		t.l.Info("└  ⬇︎ Loading list", "type", "BuiltInVariable")
		r, err := t.Service().Accounts.Containers.Workspaces.BuiltInVariables.List(t.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		res := map[string]*tagmanager.BuiltInVariable{}
		for _, value := range r.BuiltInVariable {
			res[value.Type] = value
		}
		t.builtInVariables = NewAccessedMap(res)
	}

	return t.builtInVariables, nil
}

func (t *TagManager) LookupTrigger(name string) (*tagmanager.Trigger, error) {
	elems, err := t.LoadTriggers()
	if err != nil {
		return nil, err
	}

	if !elems.Has(name) {
		return nil, ErrNotFound
	}

	return elems.Get(name), nil
}

func (t *TagManager) LookupTemplate(name string) (*tagmanager.CustomTemplate, error) {
	elems, err := t.LoadCustomTemplates()
	if err != nil {
		return nil, err
	}

	if !elems.Has(name) {
		return nil, ErrNotFound
	}

	return elems.Get(name), nil
}

func (t *TagManager) LookupTransformation(name string) (*tagmanager.Transformation, error) {
	elems, err := t.LoadTransformations()
	if err != nil {
		return nil, err
	}

	if !elems.Has(name) {
		return nil, ErrNotFound
	}

	return elems.Get(name), nil
}

func (t *TagManager) LoadTriggers() (*AccessedMap[*tagmanager.Trigger], error) {
	if t.triggers == nil {
		t.l.Info("└  ⬇︎ Loading list", "type", "Trigger")
		r, err := t.Service().Accounts.Containers.Workspaces.Triggers.List(t.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		res := map[string]*tagmanager.Trigger{}
		for _, value := range r.Trigger {
			res[value.Name] = value
		}
		t.triggers = NewAccessedMap(res)
	}

	return t.triggers, nil
}

func (t *TagManager) LookupTag(name string) (*tagmanager.Tag, error) {
	elems, err := t.LoadTags()
	if err != nil {
		return nil, err
	}

	if !elems.Has(name) {
		return nil, ErrNotFound
	}

	return elems.Get(name), nil
}

func (t *TagManager) LoadTags() (*AccessedMap[*tagmanager.Tag], error) {
	if t.tags == nil {
		t.l.Info("└  ⬇︎ Loading list", "type", "Tag")
		r, err := t.Service().Accounts.Containers.Workspaces.Tags.List(t.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		res := map[string]*tagmanager.Tag{}
		for _, value := range r.Tag {
			res[value.Name] = value
		}
		t.tags = NewAccessedMap(res)
	}

	return t.tags, nil
}

func (t *TagManager) CustomTemplate(name string) (*tagmanager.CustomTemplate, error) {
	elems, err := t.LoadCustomTemplates()
	if err != nil {
		return nil, err
	}

	if !elems.Has(name) {
		return nil, ErrNotFound
	}

	return elems.Get(name), nil
}

func (t *TagManager) LoadCustomTemplates() (*AccessedMap[*tagmanager.CustomTemplate], error) {
	if t.customTemplates == nil {
		t.l.Info("└  ⬇︎ Loading list", "type", "CustomTemplate")
		r, err := t.Service().Accounts.Containers.Workspaces.Templates.List(t.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		res := map[string]*tagmanager.CustomTemplate{}
		for _, value := range r.Template {
			res[value.Name] = value
		}
		t.customTemplates = NewAccessedMap(res)
	}

	return t.customTemplates, nil
}

func (t *TagManager) LoadTransformations() (*AccessedMap[*tagmanager.Transformation], error) {
	if t.transformations == nil {
		t.l.Info("└  ⬇︎ Loading list", "type", "Transformation")
		r, err := t.Service().Accounts.Containers.Workspaces.Transformations.List(t.WorkspacePath()).Do()
		if err != nil {
			return nil, err
		}

		res := map[string]*tagmanager.Transformation{}
		for _, value := range r.Transformation {
			res[value.Name] = value
		}
		t.transformations = NewAccessedMap(res)
	}

	return t.transformations, nil
}

func (t *TagManager) UpsertClient(folder *tagmanager.Folder, item *tagmanager.Client) (*tagmanager.Client, error) {
	l := t.l.With("type", "Client", "name", item.Name)

	item.ParentFolderId = folder.FolderId

	item.Notes = t.Notes(item)
	item.AccountId = t.AccountID()
	item.ContainerId = t.ContainerID()
	item.WorkspaceId = t.WorkspaceID()

	cache, err := t.LookupClient(item.Name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	var value *tagmanager.Client
	if cache == nil {
		l.Info("└  🚀 New")
		value, err = t.Service().Accounts.Containers.Workspaces.Clients.Create(t.WorkspacePath(), item).Do()
		t.clients.Set(item.Name, value)
	} else if item.Notes == cache.Notes {
		l.Info("└  ✔︎ OK", "id", cache.ClientId)
	} else {
		l.Info("└  🔄 Update", "id", cache.ClientId)
		value, err = t.Service().Accounts.Containers.Workspaces.Clients.Update(t.WorkspacePath()+"/clients/"+cache.ClientId, item).Do()
		t.clients.Set(item.Name, value)
	}
	if err != nil {
		if out, err := json.MarshalIndent(item, "", "  "); err == nil {
			l.Debug(string(out))
		}
		return nil, err
	}

	return t.LookupClient(item.Name)
}

func (t *TagManager) UpsertTransformation(folder *tagmanager.Folder, item *tagmanager.Transformation) (*tagmanager.Transformation, error) {
	l := t.l.With("type", "Transformation", "name", item.Name)

	item.ParentFolderId = folder.FolderId

	item.Notes = t.Notes(item)
	item.AccountId = t.AccountID()
	item.ContainerId = t.ContainerID()
	item.WorkspaceId = t.WorkspaceID()

	cache, err := t.LookupTransformation(item.Name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	var value *tagmanager.Transformation
	if cache == nil {
		l.Info("└  🚀 New")
		value, err = t.Service().Accounts.Containers.Workspaces.Transformations.Create(t.WorkspacePath(), item).Do()
		t.transformations.Set(item.Name, value)
	} else if item.Notes == cache.Notes {
		l.Info("└  ✔︎ OK", "id", cache.TransformationId)
	} else {
		l.Info("└  🔄 Update", "id", cache.TransformationId)
		value, err = t.Service().Accounts.Containers.Workspaces.Transformations.Update(t.WorkspacePath()+"/transformations/"+cache.TransformationId, item).Do()
		t.transformations.Set(item.Name, value)
	}
	if err != nil {
		if out, err := json.MarshalIndent(item, "", "  "); err == nil {
			l.Debug(string(out))
		}
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

	var value *tagmanager.Folder
	if cache == nil {
		l.Info("└  🚀 New")
		value, err = t.Service().Accounts.Containers.Workspaces.Folders.Create(t.WorkspacePath(), item).Do()
		t.folders.Set(item.Name, value)
	} else if item.Notes == cache.Notes {
		l.Info("└  ✔︎ OK", "id", item.FolderId)
	} else {
		l.Info("└  🔄 Update", "id", cache.FolderId)
		value, err = t.Service().Accounts.Containers.Workspaces.Folders.Update(t.WorkspacePath()+"/folders/"+cache.FolderId, item).Do()
		t.folders.Set(item.Name, value)
	}
	if err != nil {
		if out, err := json.MarshalIndent(item, "", "  "); err == nil {
			l.Debug(string(out))
		}
		return nil, err
	}

	return t.LookupFolder(name)
}

func (t *TagManager) UpsertVariable(folder *tagmanager.Folder, item *tagmanager.Variable) (*tagmanager.Variable, error) {
	l := t.l.With("type", "Variable", "name", item.Name)

	item.ParentFolderId = folder.FolderId

	item.Notes = t.Notes(item)
	item.AccountId = t.AccountID()
	item.ContainerId = t.ContainerID()
	item.WorkspaceId = t.WorkspaceID()

	cache, err := t.LookupVariable(item.Name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	var value *tagmanager.Variable
	if cache == nil {
		l.Info("└  🚀 New")
		value, err = t.Service().Accounts.Containers.Workspaces.Variables.Create(t.WorkspacePath(), item).Do()
		t.variables.Set(item.Name, value)
	} else if item.Notes == cache.Notes {
		l.Info("└  ✔︎ OK", "id", cache.VariableId)
	} else {
		l.Info("└  🔄 Update", "id", cache.VariableId)
		value, err = t.Service().Accounts.Containers.Workspaces.Variables.Update(t.WorkspacePath()+"/variables/"+cache.VariableId, item).Do()
		t.variables.Set(item.Name, value)
	}
	if err != nil {
		if out, err := json.MarshalIndent(item, "", "  "); err == nil {
			l.Debug(string(out))
		}
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
		l.Info("└  ✔︎ OK")
		return cache, nil
	}

	l.Info("└  🚀 New")
	resp, err := t.Service().Accounts.Containers.Workspaces.BuiltInVariables.Create(t.WorkspacePath()).Type(typeName).Do()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create built-in variable")
	}

	for _, builtInVariable := range resp.BuiltInVariable {
		t.builtInVariables.Set(builtInVariable.Type, builtInVariable)
	}

	return t.GetBuiltInVariable(typeName)
}

func (t *TagManager) UpsertTrigger(folder *tagmanager.Folder, item *tagmanager.Trigger) (*tagmanager.Trigger, error) {
	l := t.l.With("type", "Trigger", "name", item.Name)

	item.ParentFolderId = folder.FolderId

	item.Notes = t.Notes(item)
	item.AccountId = t.AccountID()
	item.ContainerId = t.ContainerID()
	item.WorkspaceId = t.WorkspaceID()

	cache, err := t.LookupTrigger(item.Name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	var value *tagmanager.Trigger
	if cache == nil {
		l.Info("└  🚀 New")
		value, err = t.Service().Accounts.Containers.Workspaces.Triggers.Create(t.WorkspacePath(), item).Do()
		t.triggers.Set(item.Name, value)
	} else if item.Notes == cache.Notes {
		l.Info("└  ✔︎ OK", "id", cache.TriggerId)
	} else {
		l.Info("└  🔄 Update", "id", cache.TriggerId)
		value, err = t.Service().Accounts.Containers.Workspaces.Triggers.Update(t.WorkspacePath()+"/triggers/"+cache.TriggerId, item).Do()
		t.triggers.Set(item.Name, value)
	}
	if err != nil {
		if out, err := json.MarshalIndent(item, "", "  "); err == nil {
			l.Debug(string(out))
		}
		return nil, err
	}

	return t.LookupTrigger(item.Name)
}

func (t *TagManager) UpsertTag(folder *tagmanager.Folder, item *tagmanager.Tag) (*tagmanager.Tag, error) {
	l := t.l.With("type", "Tag", "name", item.Name)

	item.ParentFolderId = folder.FolderId

	item.Notes = t.Notes(item)
	item.AccountId = t.AccountID()
	item.ContainerId = t.ContainerID()
	item.WorkspaceId = t.WorkspaceID()

	cache, err := t.LookupTag(item.Name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	var value *tagmanager.Tag
	if cache == nil {
		l.Info("└  🚀 New")
		value, err = t.Service().Accounts.Containers.Workspaces.Tags.Create(t.WorkspacePath(), item).Do()
		t.tags.Set(item.Name, value)
	} else if item.Notes == cache.Notes {
		l.Info("└  ✔︎ OK", "id", cache.TagId)
	} else {
		l.Info("└  🔄 Update", "id", cache.TagId)
		value, err = t.Service().Accounts.Containers.Workspaces.Tags.Update(t.WorkspacePath()+"/tags/"+cache.TagId, item).Do()
		t.tags.Set(item.Name, value)
	}
	if err != nil {
		if out, err := json.MarshalIndent(item, "", "  "); err == nil {
			l.Debug(string(out))
		}
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

	var value *tagmanager.CustomTemplate
	if cache == nil {
		l.Info("└  🚀 New")
		value, err = t.Service().Accounts.Containers.Workspaces.Templates.Create(t.WorkspacePath(), item).Do()
		t.customTemplates.Set(item.Name, value)
	} else if item.TemplateData == cache.TemplateData {
		l.Info("└  ✔︎ OK", "id", cache.TemplateId)
	} else {
		l.Info("└  🔄 Update", "id", cache.TemplateId)
		value, err = t.Service().Accounts.Containers.Workspaces.Templates.Update(t.WorkspacePath()+"/templates/"+cache.TemplateId, item).Do()
		t.customTemplates.Set(item.Name, value)
	}
	if err != nil {
		if out, err := json.MarshalIndent(item, "", "  "); err == nil {
			l.Debug(string(out))
		}
		return nil, err
	}
	return t.CustomTemplate(item.Name)
}
