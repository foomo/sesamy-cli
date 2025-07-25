package list

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	"github.com/itchyny/json2yaml"
)

func dump(i interface{ MarshalJSON() ([]byte, error) }, err error) (string, error) {
	if err != nil {
		return "", err
	}
	out, err := i.MarshalJSON()
	if err != nil {
		return "", err
	}

	var output strings.Builder
	if err := json2yaml.Convert(&output, bytes.NewBuffer(out)); err != nil {
		return "", err
	}
	return output.String(), nil
}

func list(ctx context.Context, l *slog.Logger, tm *tagmanager.TagManager, resource string) (string, error) {
	switch resource {
	case "environments":
		return dump(tm.Service().Accounts.Containers.Environments.List(tm.ContainerPath()).Context(ctx).Do())
	case "workspaces":
		return dump(tm.Service().Accounts.Containers.Workspaces.List(tm.ContainerPath()).Context(ctx).Do())
	case "clients":
		return dump(tm.Service().Accounts.Containers.Workspaces.Clients.List(tm.WorkspacePath()).Context(ctx).Do())
	case "tags":
		return dump(tm.Service().Accounts.Containers.Workspaces.Tags.List(tm.WorkspacePath()).Context(ctx).Do())
	case "built-in-variables":
		return dump(tm.Service().Accounts.Containers.Workspaces.BuiltInVariables.List(tm.WorkspacePath()).Context(ctx).Do())
	case "folders":
		return dump(tm.Service().Accounts.Containers.Workspaces.Folders.List(tm.WorkspacePath()).Context(ctx).Do())
	case "variables":
		return dump(tm.Service().Accounts.Containers.Workspaces.Variables.List(tm.WorkspacePath()).Context(ctx).Do())
	case "templates":
		return dump(tm.Service().Accounts.Containers.Workspaces.Templates.List(tm.WorkspacePath()).Context(ctx).Do())
	case "templates-data":
		r, err := tm.Service().Accounts.Containers.Workspaces.Templates.List(tm.WorkspacePath()).Context(ctx).Do()
		if err != nil {
			return "", err
		}
		var ret strings.Builder
		for _, template := range r.Template {
			ret.WriteString("---- Template data: " + template.Name + " ----------------------\n")
			ret.WriteString(template.TemplateData + "\n")
		}
		return ret.String(), nil
	case "gtag-config":
		return dump(tm.Service().Accounts.Containers.Workspaces.GtagConfig.List(tm.WorkspacePath()).Context(ctx).Do())
	case "triggers":
		return dump(tm.Service().Accounts.Containers.Workspaces.Triggers.List(tm.WorkspacePath()).Context(ctx).Do())
	case "transformations":
		return dump(tm.Service().Accounts.Containers.Workspaces.Transformations.List(tm.WorkspacePath()).Context(ctx).Do())
	case "zones":
		return dump(tm.Service().Accounts.Containers.Workspaces.Zones.List(tm.WorkspacePath()).Context(ctx).Do())
	default:
		return "", fmt.Errorf("unknown resource %s", resource)
	}
}
