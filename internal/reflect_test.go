package internal_test

import (
	"testing"

	_ "github.com/foomo/sesamy-go"              // force inclusion
	_ "github.com/foomo/sesamy-go/event/params" // force inclusion
)

func TestGetEventParameters(t *testing.T) {
	// params, err := internal.GetStructTypeParameters(context.TODO(), config.Packages{
	// 	{
	// 		Path: "github.com/foomo/sesamy-go/event",
	// 		Events: []string{
	// 			"PageView",
	// 			"SelectItem",
	// 			"AddToCart",
	// 		},
	// 	},
	// })
	// require.NoError(t, err)
	//
	// actual, err := json.Marshal(params)
	// require.NoError(t, err)
	//
	// expected := `{"add_to_cart":["currency","value","items"],"page_view":["page_title","page_location"],"select_item":["item_list_id","item_list_name","items"]}`
	// if !assert.JSONEq(t, expected, string(actual)) {
	// 	t.Log(string(actual))
	// }
}
