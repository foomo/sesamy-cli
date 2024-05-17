package internal_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/foomo/sesamy-cli/internal"
	"github.com/foomo/sesamy-cli/pkg/config"
	_ "github.com/foomo/sesamy-go"              // force inclusion
	_ "github.com/foomo/sesamy-go/event/params" // force inclusion
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetEventParameters(t *testing.T) {
	params, err := internal.GetStructTypeParameters(context.TODO(), config.Packages{
		{
			Path: "github.com/foomo/sesamy-go/event",
			Events: []string{
				"PageView",
				"SelectItem",
				"AddToCart",
			},
		},
	})
	require.NoError(t, err)

	actual, err := json.Marshal(params)
	require.NoError(t, err)

	expected := `{"add_to_cart":["currency","value","items"],"page_view":["page_title","page_location"],"select_item":["item_list_id","item_list_name","items"]}`
	if !assert.JSONEq(t, expected, string(actual)) {
		t.Log(string(actual))
	}
}

func TestScanner(t *testing.T) {
	scanner := internal.NewScanner(&internal.Config{
		Packages: []*internal.ConfigPackage{
			{
				Path:  "github.com/foomo/sesamy-go/event",
				Names: []string{"PageView"},
			},
		},
	})
	err := scanner.Scan(context.TODO())
	require.NoError(t, err)

	// actual, err := json.Marshal(scanner)
	// require.NoError(t, err)
	//
	// expected := `{"add_to_cart":["currency","value","items"],"page_view":["page_title","page_location"],"select_item":["item_list_id","item_list_name","items"]}`
	// if !assert.JSONEq(t, expected, string(actual)) {
	// 	t.Log(string(actual))
	// }
}
