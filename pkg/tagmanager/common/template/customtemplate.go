package template

import (
	"context"
	"os"

	"github.com/pkg/errors"
	"google.golang.org/api/tagmanager/v2"
)

// NewFromFile loads a custom template from the given path. The file contents are
// used verbatim, so it must be a complete GTM template.
func NewFromFile(name, filename string) (*tagmanager.CustomTemplate, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read tag template: "+filename)
	}

	return &tagmanager.CustomTemplate{
		Name:         name,
		TemplateData: string(data),
	}, nil
}

// Resolve returns the template loaded from filename, or embedded if filename is empty.
func Resolve(filename, name string, embedded *tagmanager.CustomTemplate) (*tagmanager.CustomTemplate, error) {
	if filename == "" {
		return embedded, nil
	}

	return NewFromFile(name, filename)
}

// Upserter provisions a custom template, e.g. tagmanager.TagManager.
type Upserter interface {
	UpsertCustomTemplate(ctx context.Context, item *tagmanager.CustomTemplate) (*tagmanager.CustomTemplate, error)
}

// ResolveOrLookup provisions the template from filename when set, otherwise it falls
// back to lookup, which is used by providers relying on a manually installed template.
func ResolveOrLookup(
	ctx context.Context,
	tm Upserter,
	filename, name string,
	lookup func() (*tagmanager.CustomTemplate, error),
) (*tagmanager.CustomTemplate, error) {
	if filename == "" {
		return lookup()
	}

	item, err := NewFromFile(name, filename)
	if err != nil {
		return nil, err
	}

	return tm.UpsertCustomTemplate(ctx, item)
}
