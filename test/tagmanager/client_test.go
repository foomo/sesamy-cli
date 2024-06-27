package tagmanager_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"testing"

	testingx "github.com/foomo/go/testing"
	tagx "github.com/foomo/go/testing/tag"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	"github.com/foomo/sesamy-cli/pkg/tagmanager/common/variable"
	"github.com/foomo/sesamy-cli/pkg/tagmanager/server/template"
	webtrigger "github.com/foomo/sesamy-cli/pkg/tagmanager/web/trigger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/api/option"
	gtagmanager "google.golang.org/api/tagmanager/v2"
)

func TestNewClient_Server(t *testing.T) {
	testingx.Tags(t, tagx.Skip)

	c, err := tagmanager.NewClient(
		context.TODO(),
		slog.New(slog.NewTextHandler(os.Stdout, nil)),
		os.Getenv("TEST_ACCOUNT_ID"),
		os.Getenv("TEST_SERVER_CONTAINER_ID"),
		os.Getenv("TEST_SERVER_WORKSPACE_ID"),
		os.Getenv("TEST_MEASUREMENT_ID"),
		tagmanager.ClientWithClientOptions(
			option.WithCredentialsFile(os.Getenv("TEST_CREDENTIALS_FILE")),
		),
	)
	require.NoError(t, err)

	{ // --- Folders ---
		t.Run("upsert folder", func(t *testing.T) {
			obj, err := c.UpsertFolder("Sesamy")
			require.NoError(t, err)
			dump(t, obj)
		})
	}

	{ // --- Variables ---
		t.Run("upsert GTM client", func(t *testing.T) {
			client, err := c.UpsertVariable(variable.NewConstant("web-container-id", os.Getenv("TEST_WEB_CONTAINER_GID")))
			if assert.NoError(t, err) {
				dump(t, client)
			}
		})
	}

	{ // --- Built-In Variables ---
		t.Run("list built-in variables", func(t *testing.T) {
			cmd := c.Service().Accounts.Containers.Workspaces.BuiltInVariables.List(c.WorkspacePath())
			if r, err := cmd.Do(); assert.NoError(t, err) {
				dump(t, r)
			}
		})
		t.Run("list built-in variables", func(t *testing.T) {
			cmd := c.Service().Accounts.Containers.Workspaces.BuiltInVariables.Create(c.WorkspacePath()).Type()
			if r, err := cmd.Do(); assert.NoError(t, err) {
				dump(t, r)
			}
		})
	}

	{ // --- Clients ---
		t.Run("list clients", func(t *testing.T) {
			cmd := c.Service().Accounts.Containers.Workspaces.Clients.List(c.WorkspacePath())
			if r, err := cmd.Do(); assert.NoError(t, err) {
				dump(t, r)
			}
		})

		// t.Run("upsert GTM client", func(t *testing.T) {
		// 	client, err := c.UpsertClient(client2.NewGTM("Google Tag Manager Web Container", "Constant.web-container-id"))
		// 	if assert.NoError(t, err) {
		// 		dump(t, client)
		// 	}
		// })
	}

	{ // --- Triggers ---
		t.Run("list triggers", func(t *testing.T) {
			cmd := c.Service().Accounts.Containers.Workspaces.Triggers.List(c.WorkspacePath())
			if r, err := cmd.Do(); assert.NoError(t, err) {
				dump(t, r)
			}
		})
	}

	{ // --- Tags ---
		t.Run("list tags", func(t *testing.T) {
			cmd := c.Service().Accounts.Containers.Workspaces.Tags.List(c.WorkspacePath())
			if r, err := cmd.Do(); assert.NoError(t, err) {
				dump(t, r)
			}
		})
	}

	{ // --- Tags ---
		t.Run("upsert template", func(t *testing.T) {
			obj, err := c.UpsertCustomTemplate(template.NewUmami("Sesamy Umami"))
			require.NoError(t, err)
			dump(t, obj)
		})

		t.Run("list templates", func(t *testing.T) {
			cmd := c.Service().Accounts.Containers.Workspaces.Templates.List(c.WorkspacePath())
			if r, err := cmd.Do(); assert.NoError(t, err) {
				// dump(t, r)
				fmt.Println(r.Template[2].TemplateData)
			}
		})
	}
}

func TestNewClient_Web(t *testing.T) {
	testingx.Tags(t, tagx.Skip)

	c, err := tagmanager.NewClient(
		context.TODO(),
		slog.New(slog.NewTextHandler(os.Stdout, nil)),
		os.Getenv("TEST_ACCOUNT_ID"),
		os.Getenv("TEST_WEB_CONTAINER_ID"),
		os.Getenv("TEST_WEB_WORKSPACE_ID"),
		os.Getenv("TEST_MEASUREMENT_ID"),
		tagmanager.ClientWithClientOptions(
			option.WithCredentialsFile(os.Getenv("TEST_CREDENTIALS_FILE")),
		),
	)
	require.NoError(t, err)

	{ // --- Containers ---
		t.Run("list containers", func(t *testing.T) {
			cmd := c.Service().Accounts.Containers.List(c.AccountPath())
			if r, err := cmd.Do(); assert.NoError(t, err) {
				dump(t, r)
			}
		})
	}

	{ // --- Workspaces ---
		t.Run("list workspaces", func(t *testing.T) {
			cmd := c.Service().Accounts.Containers.Workspaces.List(c.ContainerPath())
			if r, err := cmd.Do(); assert.NoError(t, err) {
				dump(t, r)
			}
		})
	}

	folderID := "25"
	{ // --- Folders ---
		name := "Sesamy"
		t.Run("list folders", func(t *testing.T) {
			cmd := c.Service().Accounts.Containers.Workspaces.Folders.List(c.WorkspacePath())
			if r, err := cmd.Do(); assert.NoError(t, err) {
				dump(t, r)
			}
		})

		t.Run("create folder", func(t *testing.T) {
			cmd := c.Service().Accounts.Containers.Workspaces.Folders.Create(c.WorkspacePath(), &gtagmanager.Folder{
				AccountId:   c.AccountID(),
				ContainerId: c.ContainerID(),
				WorkspaceId: c.WorkspaceID(),
				Name:        name,
				Notes:       c.Notes(nil),
			})
			if r, err := cmd.Do(); assert.NoError(t, err) {
				dump(t, r)
			}
		})

		t.Run("get folder", func(t *testing.T) {
			cmd := c.Service().Accounts.Containers.Workspaces.Folders.List(c.WorkspacePath())
			if r, err := cmd.Do(); assert.NoError(t, err) {
				for _, folder := range r.Folder {
					if folder.Name == name {
						t.Log("ID: " + folder.FolderId)
						return
					}
				}
				t.Error("not found")
			}
		})

		t.Run("upsert folder", func(t *testing.T) {
			obj, err := c.UpsertFolder(name)
			require.NoError(t, err)
			t.Log("ID: " + obj.FolderId)
		})
	}

	{ // --- Built-In Variables ---
		t.Run("list built-in variables", func(t *testing.T) {
			cmd := c.Service().Accounts.Containers.Workspaces.BuiltInVariables.List(c.WorkspacePath())
			if r, err := cmd.Do(); assert.NoError(t, err) {
				dump(t, r)
			}
		})
	}

	{ // --- Built-In Variables ---
		t.Run("list gtag config", func(t *testing.T) {
			cmd := c.Service().Accounts.Containers.Workspaces.GtagConfig.List(c.WorkspacePath())
			if r, err := cmd.Do(); assert.NoError(t, err) {
				dump(t, r)
			}
		})
	}

	{ // --- Variables ---
		name := "Constant.ga4-measurement-id"
		t.Run("list variables", func(t *testing.T) {
			cmd := c.Service().Accounts.Containers.Workspaces.Variables.List(c.WorkspacePath())
			if r, err := cmd.Do(); assert.NoError(t, err) {
				dump(t, r)
			}
		})

		t.Run("create variable", func(t *testing.T) {
			cmd := c.Service().Accounts.Containers.Workspaces.Variables.Create(c.WorkspacePath(), &gtagmanager.Variable{
				AccountId:      c.AccountID(),
				ContainerId:    c.ContainerID(),
				WorkspaceId:    c.WorkspaceID(),
				ParentFolderId: folderID,
				Name:           "Constant." + name,
				Notes:          c.Notes(nil),
				Parameter: []*gtagmanager.Parameter{
					{
						Key:   "value",
						Type:  "template",
						Value: c.MeasurementID(),
					},
				},
				Type: "c",
			})
			if r, err := cmd.Do(); assert.NoError(t, err) {
				dump(t, r)
			}
		})

		t.Run("get variable", func(t *testing.T) {
			cmd := c.Service().Accounts.Containers.Workspaces.Variables.List(c.WorkspacePath())
			if r, err := cmd.Do(); assert.NoError(t, err) {
				for _, value := range r.Variable {
					if value.Name == "Constant."+name {
						t.Log("ID: " + value.VariableId)
						return
					}
				}
				t.Error("not found")
			}
		})

		t.Run("upsert variable", func(t *testing.T) {
			obj, err := c.UpsertVariable(variable.NewConstant(name, c.MeasurementID()))
			require.NoError(t, err)
			t.Log("ID: " + obj.VariableId)
		})
	}

	{ // --- Triggers ---
		name := "login"
		t.Run("list triggers", func(t *testing.T) {
			cmd := c.Service().Accounts.Containers.Workspaces.Triggers.List(c.WorkspacePath())
			if r, err := cmd.Do(); assert.NoError(t, err) {
				dump(t, r)
			}
		})

		t.Run("create trigger", func(t *testing.T) {
			cmd := c.Service().Accounts.Containers.Workspaces.Triggers.Create(c.WorkspacePath(), &gtagmanager.Trigger{
				AccountId:      c.AccountID(),
				ContainerId:    c.ContainerID(),
				WorkspaceId:    c.WorkspaceID(),
				ParentFolderId: folderID,
				Type:           "customEvent",
				Name:           "Event." + name,
				Notes:          c.Notes(nil),
				CustomEventFilter: []*gtagmanager.Condition{
					{
						Parameter: []*gtagmanager.Parameter{
							{
								Key:   "arg0",
								Type:  "template",
								Value: "{{_event}}",
							},
							{
								Key:   "arg1",
								Type:  "template",
								Value: name,
							},
						},
						Type: "equals",
					},
				},
			})
			if r, err := cmd.Do(); assert.NoError(t, err) {
				dump(t, r)
			}
		})

		t.Run("get trigger", func(t *testing.T) {
			cmd := c.Service().Accounts.Containers.Workspaces.Triggers.List(c.WorkspacePath())
			if r, err := cmd.Do(); assert.NoError(t, err) {
				for _, value := range r.Trigger {
					if value.Name == "Event."+name {
						t.Log("ID: " + value.TriggerId)
						return
					}
				}
				t.Error("not found")
			}
		})

		t.Run("upsert trigger", func(t *testing.T) {
			obj, err := c.UpsertTrigger(webtrigger.NewCustomEvent("Event."+name, name))
			require.NoError(t, err)
			t.Log("ID: " + obj.TriggerId)
		})
	}

	{ // --- Tags ---
		name := "login"
		t.Run("list tags", func(t *testing.T) {
			cmd := c.Service().Accounts.Containers.Workspaces.Tags.List(c.WorkspacePath())
			if r, err := cmd.Do(); assert.NoError(t, err) {
				dump(t, r)
			}
		})

		t.Run("create tag", func(t *testing.T) {
			cmd := c.Service().Accounts.Containers.Workspaces.Tags.Create(c.WorkspacePath(), &gtagmanager.Tag{
				AccountId:      c.AccountID(),
				ContainerId:    c.ContainerID(),
				WorkspaceId:    c.WorkspaceID(),
				ParentFolderId: folderID,
				Name:           "GA4." + name,
				Notes:          c.Notes(nil),
				Parameter: []*gtagmanager.Parameter{
					{
						Type:  "boolean",
						Key:   "sendEcommerceData",
						Value: "true",
					},
					{
						Type:  "template",
						Key:   "getEcommerceDataFrom",
						Value: "dataLayer",
					},
					{
						Type:  "boolean",
						Key:   "enhancedUserId",
						Value: "false",
					},
					{
						Type:  "template",
						Key:   "eventName",
						Value: "Event." + name,
					},
					{
						Type:  "template",
						Key:   "measurementIdOverride",
						Value: "{{GA Measurement ID}}",
					},
					{
						Type: "list",
						Key:  "eventSettingsTable",
						List: []*gtagmanager.Parameter{
							{
								Type: "map",
								Map: []*gtagmanager.Parameter{
									{
										Type:  "parameter",
										Key:   "template",
										Value: "method",
									},
									{
										Type:  "parameterValue",
										Key:   "template",
										Value: "{{EventModel.method}}",
									},
								},
							},
						},
					},
				},
				Type: "gaawe",
			})
			if r, err := cmd.Do(); assert.NoError(t, err) {
				dump(t, r)
			}
		})

		// type Login struct {
		// 	Method string `json:"method"`
		// }

		// t.Run("upsert tag", func(t *testing.T) {
		//	obj, err := c.UpsertGA4WebTag(ctx, "login", eventParameters(Login{}))
		//	require.NoError(t, err)
		//	t.Log("ID: " + obj.TagId)
		// })
	}
}

// ------------------------------------------------------------------------------------------------
// ~ Private methods
// ------------------------------------------------------------------------------------------------

func dump(t *testing.T, i interface{ MarshalJSON() ([]byte, error) }) {
	t.Helper()
	var ret bytes.Buffer
	out, err := i.MarshalJSON()
	require.NoError(t, err)
	require.NoError(t, json.Indent(&ret, out, "", "  "))
	t.Log(ret.String())
}
