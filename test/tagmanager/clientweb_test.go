package tagmanager_test

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"testing"

	testingx "github.com/foomo/go/testing"
	tagx "github.com/foomo/go/testing/tag"
	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	"github.com/foomo/sesamy-cli/pkg/tagmanager/common/trigger"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/api/option"
	gtagmanager "google.golang.org/api/tagmanager/v2"
)

func TestNewClient_Web(t *testing.T) {
	testingx.Tags(t, tagx.Skip)
	require.NoError(t, godotenv.Load("../../.env"))

	c, err := tagmanager.New(
		context.TODO(),
		slog.New(slog.NewTextHandler(os.Stdout, nil)),
		os.Getenv("TEST_ACCOUNT_ID"),
		config.GoogleTagManagerContainer{
			TagID:       os.Getenv("TEST_WEB_TAG_ID"),
			ContainerID: os.Getenv("TEST_WEB_CONTAINER_ID"),
			WorkspaceID: os.Getenv("TEST_WEB_WORKSPACE_ID"),
		},
		tagmanager.WithClientOptions(
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
		t.Run("list variables", func(t *testing.T) {
			cmd := c.Service().Accounts.Containers.Workspaces.Variables.List(c.WorkspacePath())
			if r, err := cmd.Do(); assert.NoError(t, err) {
				dump(t, r)
			}
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
			obj, err := c.UpsertTrigger(trigger.NewEvent(name))
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

	{ // --- Templates ---
		t.Run("list templates", func(t *testing.T) {
			cmd := c.Service().Accounts.Containers.Workspaces.Templates.List(c.WorkspacePath())
			if r, err := cmd.Do(); assert.NoError(t, err) {
				// dump(t, r)
				fmt.Println(r.Template[0].TemplateData)
			}
		})
	}
}
