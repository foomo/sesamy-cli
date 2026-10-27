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
