package internal_test

import (
	"context"
	"fmt"
	"go/types"
	"testing"

	"github.com/foomo/sesamy-cli/internal"
	_ "github.com/foomo/sesamy-go"              // force inclusion
	_ "github.com/foomo/sesamy-go/event/params" // force inclusion
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewLoader(t *testing.T) {
	scanner := internal.NewLoader(&internal.LoaderConfig{
		Packages: []*internal.PackageConfig{
			{
				Path:  "github.com/foomo/sesamy-go/event",
				Types: []string{"PageView"},
			},
		},
	})
	err := scanner.Load(context.TODO())
	require.NoError(t, err)

	{
		pkg := scanner.Package("github.com/foomo/sesamy-go/event")
		pkgType := pkg.LookupScopeType("PageView")
		if eventStruct := internal.TC[*types.Struct](pkgType.Type().Underlying()); eventStruct != nil {
			for i := range eventStruct.NumFields() {
				if eventField := eventStruct.Field(i); eventField.Name() == "Params" {
					if paramsStruct := internal.TC[*types.Struct](eventField.Type().Underlying()); paramsStruct != nil {
						for j := range paramsStruct.NumFields() {
							fmt.Println(paramsStruct.Field(j).Name())
						}
					}
				}
			}
		}
	}

	// actual, err := json.Marshal(scanner)
	// require.NoError(t, err)
	//
	// expected := `{"add_to_cart":["currency","value","items"],"page_view":["page_title","page_location"],"select_item":["item_list_id","item_list_name","items"]}`
	// if !assert.JSONEq(t, expected, string(actual)) {
	// 	t.Log(string(actual))
	// }
}

func TestLoader_LookupTypesByType(t *testing.T) {
	scanner := internal.NewLoader(&internal.LoaderConfig{
		Packages: []*internal.PackageConfig{
			{
				Path:  "github.com/foomo/sesamy-go/event",
				Types: []string{"PageView"},
			},
		},
	})
	err := scanner.Load(context.TODO())
	require.NoError(t, err)

	pkg := scanner.Package("github.com/foomo/sesamy-go")
	require.NotNil(t, pkg)
	pkgType := pkg.LookupType("Event")
	require.NotNil(t, pkgType)

	pkgTypes := scanner.LookupTypesByType(pkgType)
	assert.NotEmpty(t, pkgTypes)
}
