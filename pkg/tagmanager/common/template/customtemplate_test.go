package template_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/foomo/sesamy-cli/pkg/tagmanager/common/template"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/api/tagmanager/v2"
)

func TestNewFromFile(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		data     string
		filename string
		wantErr  bool
	}{
		{
			name:     "reads template verbatim",
			data:     "___INFO___\n\n{\n  \"type\": \"TAG\"\n}\n",
			filename: "mixpanel.tpl",
		},
		{
			name:     "keeps percent verbs untouched",
			data:     "___INFO___ 100% \"displayName\": \"%s\"\n",
			filename: "verbs.tpl",
		},
		{
			name:     "missing file",
			filename: "does-not-exist.tpl",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			path := filepath.Join(t.TempDir(), tt.filename)
			if !tt.wantErr {
				require.NoError(t, os.WriteFile(path, []byte(tt.data), 0o600))
			}

			actual, err := template.NewFromFile("Mixpanel", path)
			if tt.wantErr {
				require.Error(t, err)
				assert.Nil(t, actual)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, "Mixpanel", actual.Name)
			assert.Equal(t, tt.data, actual.TemplateData)
		})
	}
}

func TestResolve(t *testing.T) {
	t.Parallel()

	embedded := &tagmanager.CustomTemplate{Name: "Embedded", TemplateData: "___INFO___ embedded"}

	t.Run("empty path returns embedded", func(t *testing.T) {
		t.Parallel()

		actual, err := template.Resolve("", "Custom", embedded)
		require.NoError(t, err)
		assert.Same(t, embedded, actual)
	})

	t.Run("path overrides embedded", func(t *testing.T) {
		t.Parallel()

		path := filepath.Join(t.TempDir(), "custom.tpl")
		require.NoError(t, os.WriteFile(path, []byte("___INFO___ custom"), 0o600))

		actual, err := template.Resolve(path, "Custom", embedded)
		require.NoError(t, err)
		assert.Equal(t, "Custom", actual.Name)
		assert.Equal(t, "___INFO___ custom", actual.TemplateData)
	})

	t.Run("missing path errors", func(t *testing.T) {
		t.Parallel()

		actual, err := template.Resolve(filepath.Join(t.TempDir(), "nope.tpl"), "Custom", embedded)
		require.Error(t, err)
		assert.Nil(t, actual)
	})
}

type upserterStub struct {
	called bool
	item   *tagmanager.CustomTemplate
}

func (u *upserterStub) UpsertCustomTemplate(_ context.Context, item *tagmanager.CustomTemplate) (*tagmanager.CustomTemplate, error) {
	u.called = true
	u.item = item

	return item, nil
}

func TestResolveOrLookup(t *testing.T) {
	t.Parallel()

	looked := &tagmanager.CustomTemplate{Name: "Gallery", TemplateData: "___INFO___ gallery"}
	lookup := func() (*tagmanager.CustomTemplate, error) { return looked, nil }

	t.Run("empty path looks up without upserting", func(t *testing.T) {
		t.Parallel()

		tm := &upserterStub{}

		actual, err := template.ResolveOrLookup(t.Context(), tm, "", "Custom", lookup)
		require.NoError(t, err)
		assert.Same(t, looked, actual)
		assert.False(t, tm.called, "must not provision when no override is set")
	})

	t.Run("path upserts instead of looking up", func(t *testing.T) {
		t.Parallel()

		path := filepath.Join(t.TempDir(), "custom.tpl")
		require.NoError(t, os.WriteFile(path, []byte("___INFO___ custom"), 0o600))

		tm := &upserterStub{}

		actual, err := template.ResolveOrLookup(t.Context(), tm, path, "Custom", lookup)
		require.NoError(t, err)
		require.True(t, tm.called, "must provision the custom template")
		assert.Equal(t, "Custom", actual.Name)
		assert.Equal(t, "___INFO___ custom", actual.TemplateData)
	})

	t.Run("missing path errors without upserting", func(t *testing.T) {
		t.Parallel()

		tm := &upserterStub{}

		actual, err := template.ResolveOrLookup(t.Context(), tm, filepath.Join(t.TempDir(), "nope.tpl"), "Custom", lookup)
		require.Error(t, err)
		assert.Nil(t, actual)
		assert.False(t, tm.called)
	})
}
