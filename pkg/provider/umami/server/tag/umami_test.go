package tag_test

import (
	"testing"

	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/provider/umami/server/tag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/api/tagmanager/v2"
)

func TestNewUmami(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		params map[string]string
		want   []string
	}{
		{
			name:   "no params",
			params: nil,
			want:   nil,
		},
		{
			name:   "empty params",
			params: map[string]string{},
			want:   nil,
		},
		{
			name: "sorted params",
			params: map[string]string{
				"value":    "eventModel.value",
				"currency": "eventModel.currency",
				"":         "eventModel.",
			},
			want: []string{"currency", "value"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := tag.NewUmami("add_to_cart", config.Umami{WebsiteID: "website"}, &tagmanager.CustomTemplate{}, tt.params)
			assert.Equal(t, "Umami - add_to_cart", got.Name)

			var list *tagmanager.Parameter

			for _, p := range got.Parameter {
				if p.Key == "eventParameters" {
					list = p
				}
			}

			if tt.want == nil {
				assert.Nil(t, list)
				return
			}

			require.NotNil(t, list)
			assert.Equal(t, "list", list.Type)

			names := make([]string, 0, len(list.List))
			for _, item := range list.List {
				require.Len(t, item.Map, 1)
				assert.Equal(t, "name", item.Map[0].Key)
				names = append(names, item.Map[0].Value)
			}

			assert.Equal(t, tt.want, names)
		})
	}
}

func TestNewUmami_DistinctIDParameter(t *testing.T) {
	t.Parallel()

	find := func(tag *tagmanager.Tag, key string) *tagmanager.Parameter {
		for _, p := range tag.Parameter {
			if p.Key == key {
				return p
			}
		}

		return nil
	}

	t.Run("set", func(t *testing.T) {
		t.Parallel()

		got := tag.NewUmami("page_view", config.Umami{DistinctIDParameter: "session_id"}, &tagmanager.CustomTemplate{}, nil)

		p := find(got, "distinctIdParameter")
		require.NotNil(t, p)
		assert.Equal(t, "template", p.Type)
		assert.Equal(t, "session_id", p.Value)
	})

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		got := tag.NewUmami("page_view", config.Umami{}, &tagmanager.CustomTemplate{}, nil)

		assert.Nil(t, find(got, "distinctIdParameter"))
	})
}
