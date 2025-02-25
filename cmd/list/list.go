package list

import (
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/alecthomas/chroma/quick"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	"github.com/itchyny/json2yaml"
)

func dump(i interface{ MarshalJSON() ([]byte, error) }, err error) error {
	if err != nil {
		return err
	}
	out, err := i.MarshalJSON()
	if err != nil {
		return err
	}
	// if err := json.Indent(ret, out, "", "  "); err != nil {
	// 	return err
	// }
	// fmt.Println(ret.String())
	var output strings.Builder
	if err := json2yaml.Convert(&output, bytes.NewBuffer(out)); err != nil {
		return err
	}
	// fmt.Print(output.String())
	return quick.Highlight(os.Stdout, output.String(), "yaml", "terminal", "monokai")
}

func list(l *slog.Logger, tm *tagmanager.TagManager, resource string) error {
	switch resource {
	case "status":
		return dump(tm.Service().Accounts.Containers.Workspaces.GetStatus(tm.WorkspacePath()).Do())
	case "clients":
		return dump(tm.Service().Accounts.Containers.Workspaces.Clients.List(tm.WorkspacePath()).Do())
	case "tags":
		return dump(tm.Service().Accounts.Containers.Workspaces.Tags.List(tm.WorkspacePath()).Do())
	case "built-in-variables":
		return dump(tm.Service().Accounts.Containers.Workspaces.BuiltInVariables.List(tm.WorkspacePath()).Do())
	case "folders":
		return dump(tm.Service().Accounts.Containers.Workspaces.Folders.List(tm.WorkspacePath()).Do())
	case "variables":
		return dump(tm.Service().Accounts.Containers.Workspaces.Variables.List(tm.WorkspacePath()).Do())
	case "templates":
		return dump(tm.Service().Accounts.Containers.Workspaces.Templates.List(tm.WorkspacePath()).Do())
	case "templates-data":
		r, err := tm.Service().Accounts.Containers.Workspaces.Templates.List(tm.WorkspacePath()).Do()
		if err != nil {
			return err
		}
		for _, template := range r.Template {
			l.Info("---- Template data: " + template.Name + " ----------------------")
			fmt.Println(template.TemplateData)
		}
		return nil
	case "gtag-config":
		return dump(tm.Service().Accounts.Containers.Workspaces.GtagConfig.List(tm.WorkspacePath()).Do())
	case "triggers":
		return dump(tm.Service().Accounts.Containers.Workspaces.Triggers.List(tm.WorkspacePath()).Do())
	case "transformations":
		return dump(tm.Service().Accounts.Containers.Workspaces.Transformations.List(tm.WorkspacePath()).Do())
	case "zones":
		return dump(tm.Service().Accounts.Containers.Workspaces.Zones.List(tm.WorkspacePath()).Do())
	default:
		return fmt.Errorf("unknown resource %s", resource)
	}
}
