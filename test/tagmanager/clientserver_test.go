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
	"github.com/foomo/sesamy-cli/pkg/provider/googleconsent/server/template"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/api/option"
)

func TestNewClient_Server(t *testing.T) {
	testingx.Tags(t, tagx.Skip)
	require.NoError(t, godotenv.Load("../../.env"))

	c, err := tagmanager.New(
		context.TODO(),
		slog.New(slog.NewTextHandler(os.Stdout, nil)),
		os.Getenv("TEST_ACCOUNT_ID"),
		config.GoogleTagManagerContainer{
			TagID:       os.Getenv("TEST_SERVER_TAG_ID"),
			ContainerID: os.Getenv("TEST_SERVER_CONTAINER_ID"),
			WorkspaceID: os.Getenv("TEST_SERVER_WORKSPACE_ID"),
		},
		tagmanager.WithClientOptions(
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

	{ // ---  Variables ---
		t.Run("list variables", func(t *testing.T) {
			cmd := c.Service().Accounts.Containers.Workspaces.Variables.List(c.WorkspacePath())
			if r, err := cmd.Do(); assert.NoError(t, err) {
				dump(t, r)
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
		t.Run("create built-in variables", func(t *testing.T) {
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

	{ // --- Transformations ---
		t.Run("list transformations", func(t *testing.T) {
			cmd := c.Service().Accounts.Containers.Workspaces.Transformations.List(c.WorkspacePath())
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

	{ // --- Templates ---
		t.Run("upsert template", func(t *testing.T) {
			obj, err := c.UpsertCustomTemplate(template.NewGoogleConsentModeCheck("TESTOMAT"))
			require.NoError(t, err)
			dump(t, obj)
		})
		t.Run("list templates", func(t *testing.T) {
			cmd := c.Service().Accounts.Containers.Workspaces.Templates.List(c.WorkspacePath())
			if r, err := cmd.Do(); assert.NoError(t, err) {
				dump(t, r)
				fmt.Println(r.Template[5].TemplateData)
			}
		})
	}
}
