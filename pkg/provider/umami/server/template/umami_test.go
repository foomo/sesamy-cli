package template_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/foomo/sesamy-cli/pkg/provider/umami/server/template"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUmami(t *testing.T) {
	t.Parallel()

	got := template.NewUmami("Umami")
	assert.Equal(t, "Umami", got.Name)

	section := func(name, next string) string {
		_, after, found := strings.Cut(got.TemplateData, name)
		require.True(t, found, name)
		before, _, found := strings.Cut(after, next)
		require.True(t, found, next)

		return before
	}

	var params []map[string]any
	require.NoError(t, json.Unmarshal([]byte(section("___TEMPLATE_PARAMETERS___", "___SANDBOXED_JS_FOR_SERVER___")), &params))
	assert.Contains(t, section("___TEMPLATE_PARAMETERS___", "___SANDBOXED_JS_FOR_SERVER___"), `"name": "eventParameters"`)

	var permissions []map[string]any
	require.NoError(t, json.Unmarshal([]byte(section("___SERVER_PERMISSIONS___", "___TESTS___")), &permissions))

	js := section("___SANDBOXED_JS_FOR_SERVER___", "___SERVER_PERMISSIONS___")
	assert.Contains(t, js, "umamiEvent.payload.ip = ip;")
	assert.Contains(t, js, "umamiEvent.payload.data = umamiData;")
}
