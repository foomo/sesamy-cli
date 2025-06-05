package diff

import (
	"bytes"
	"context"
	"log/slog"
	"strings"

	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	"github.com/itchyny/json2yaml"
	"github.com/sters/yaml-diff/yamldiff"
)

func diff(ctx context.Context, l *slog.Logger, tm *tagmanager.TagManager) (string, error) {
	l.Info("└  ⬇︎ Loading status")
	s, err := tm.Service().Accounts.Containers.Workspaces.GetStatus(tm.WorkspacePath()).Context(ctx).Do()
	if err != nil {
		return "", err
	} else if len(s.WorkspaceChange) == 0 {
		return "", nil
	}

	l.Info("└  ⬇︎ Loading live version")
	live, err := tm.Service().Accounts.Containers.Versions.Live(tm.ContainerPath()).Do()
	if err != nil {
		return "", err
	}

	var res []string
	for _, entity := range s.WorkspaceChange {
		switch {
		case entity.Tag != nil:
			res = append(res, "  # Tag: "+entity.Tag.Name+" ("+entity.ChangeStatus+")\n")

			// unset props
			entity.Tag.Path = ""
			entity.Tag.Fingerprint = ""
			entity.Tag.WorkspaceId = ""
			entity.Tag.TagManagerUrl = ""

			var changed string
			if entity.ChangeStatus != "deleted" {
				changed, err = ToYalm(entity.Tag)
				if err != nil {
					return "", err
				}
			}

			var original string
			for _, value := range live.Tag {
				if value.Name == entity.Tag.Name {
					// unset props
					value.Fingerprint = ""

					original, err = ToYalm(value)
					if err != nil {
						return "", err
					}
					break
				}
			}

			d, err := ToDiff(original, changed)
			if err != nil {
				return "", err
			}
			res = append(res, d...)
		case entity.Folder != nil:
			res = append(res, "  # Folder: "+entity.Folder.Name+" ("+entity.ChangeStatus+")")

			// unset props
			entity.Folder.Path = ""
			entity.Folder.Fingerprint = ""
			entity.Folder.WorkspaceId = ""
			entity.Folder.TagManagerUrl = ""

			var changed string
			if entity.ChangeStatus != "deleted" {
				changed, err = ToYalm(entity.Folder)
				if err != nil {
					return "", err
				}
			}

			var original string
			for _, value := range live.Folder {
				if value.Name == entity.Folder.Name {
					// unset props
					value.Fingerprint = ""

					original, err = ToYalm(value)
					if err != nil {
						return "", err
					}
					break
				}
			}

			d, err := ToDiff(original, changed)
			if err != nil {
				return "", err
			}
			res = append(res, d...)
		case entity.Trigger != nil:
			res = append(res, "  # Trigger: "+entity.Trigger.Name+" ("+entity.ChangeStatus+")")

			// unset props
			entity.Trigger.Path = ""
			entity.Trigger.Fingerprint = ""
			entity.Trigger.WorkspaceId = ""
			entity.Trigger.TagManagerUrl = ""

			var changed string
			if entity.ChangeStatus != "deleted" {
				changed, err = ToYalm(entity.Trigger)
				if err != nil {
					return "", err
				}
			}

			var original string
			for _, value := range live.Trigger {
				if value.Name == entity.Trigger.Name {
					// unset props
					value.Fingerprint = ""

					original, err = ToYalm(value)
					if err != nil {
						return "", err
					}
					break
				}
			}

			d, err := ToDiff(original, changed)
			if err != nil {
				return "", err
			}
			res = append(res, d...)
		case entity.Variable != nil:
			res = append(res, "  # Variable: "+entity.Variable.Name+" ("+entity.ChangeStatus+")")

			// unset props
			entity.Variable.Path = ""
			entity.Variable.Fingerprint = ""
			entity.Variable.WorkspaceId = ""
			entity.Variable.TagManagerUrl = ""

			var changed string
			if entity.ChangeStatus != "deleted" {
				changed, err = ToYalm(entity.Variable)
				if err != nil {
					return "", err
				}
			}

			var original string
			for _, value := range live.Variable {
				if value.Name == entity.Variable.Name {
					// unset props
					value.Fingerprint = ""

					original, err = ToYalm(value)
					if err != nil {
						return "", err
					}
					break
				}
			}

			d, err := ToDiff(original, changed)
			if err != nil {
				return "", err
			}
			res = append(res, d...)
		case entity.Client != nil:
			res = append(res, "  # Client: "+entity.Client.Name+" ("+entity.ChangeStatus+")")

			// unset props
			entity.Client.Path = ""
			entity.Client.Fingerprint = ""
			entity.Client.WorkspaceId = ""
			entity.Client.TagManagerUrl = ""

			var changed string
			if entity.ChangeStatus != "deleted" {
				changed, err = ToYalm(entity.Client)
				if err != nil {
					return "", err
				}
			}

			var original string
			for _, value := range live.Client {
				if value.Name == entity.Client.Name {
					// unset props
					value.Fingerprint = ""

					original, err = ToYalm(value)
					if err != nil {
						return "", err
					}
					break
				}
			}

			d, err := ToDiff(original, changed)
			if err != nil {
				return "", err
			}
			res = append(res, d...)
		case entity.GtagConfig != nil:
			res = append(res, "  # GtagConfig: "+entity.GtagConfig.AccountId+" ("+entity.ChangeStatus+")")

			// unset props
			entity.GtagConfig.Path = ""
			entity.GtagConfig.Fingerprint = ""
			entity.GtagConfig.WorkspaceId = ""
			entity.GtagConfig.TagManagerUrl = ""

			var changed string
			if entity.ChangeStatus != "deleted" {
				changed, err = ToYalm(entity.GtagConfig)
				if err != nil {
					return "", err
				}
			}

			var original string
			for _, value := range live.GtagConfig {
				if value.AccountId == entity.GtagConfig.AccountId {
					// unset props
					value.Fingerprint = ""

					original, err = ToYalm(value)
					if err != nil {
						return "", err
					}
					break
				}
			}

			d, err := ToDiff(original, changed)
			if err != nil {
				return "", err
			}
			res = append(res, d...)
		case entity.BuiltInVariable != nil:
			res = append(res, "  # BuiltInVariable: "+entity.BuiltInVariable.Name+" ("+entity.ChangeStatus+")")

			// unset props
			entity.BuiltInVariable.Path = ""
			entity.BuiltInVariable.WorkspaceId = ""

			var changed string
			if entity.ChangeStatus != "deleted" {
				changed, err = ToYalm(entity.BuiltInVariable)
				if err != nil {
					return "", err
				}
			}

			var original string
			for _, value := range live.BuiltInVariable {
				if value.Name == entity.BuiltInVariable.Name {
					original, err = ToYalm(value)
					if err != nil {
						return "", err
					}
					break
				}
			}

			d, err := ToDiff(original, changed)
			if err != nil {
				return "", err
			}
			res = append(res, d...)
		case entity.CustomTemplate != nil:
			res = append(res, "  # CustomTemplate: "+entity.CustomTemplate.Name+" ("+entity.ChangeStatus+")")

			// unset props
			entity.CustomTemplate.Path = ""
			entity.CustomTemplate.Fingerprint = ""
			entity.CustomTemplate.WorkspaceId = ""
			entity.CustomTemplate.TagManagerUrl = ""

			var changed string
			if entity.ChangeStatus != "deleted" {
				changed, err = ToYalm(entity.CustomTemplate)
				if err != nil {
					return "", err
				}
			}

			var original string
			for _, value := range live.CustomTemplate {
				if value.Name == entity.CustomTemplate.Name {
					// unset props
					value.Fingerprint = ""

					original, err = ToYalm(value)
					if err != nil {
						return "", err
					}
					break
				}
			}

			d, err := ToDiff(original, changed)
			if err != nil {
				return "", err
			}
			res = append(res, d...)
		case entity.Transformation != nil:
			res = append(res, "  # Transformation: "+entity.Transformation.Name+" ("+entity.ChangeStatus+")")

			// unset props
			entity.Transformation.Path = ""
			entity.Transformation.Fingerprint = ""
			entity.Transformation.WorkspaceId = ""
			entity.Transformation.TagManagerUrl = ""

			var changed string
			if entity.ChangeStatus != "deleted" {
				changed, err = ToYalm(entity.Transformation)
				if err != nil {
					return "", err
				}
			}

			var original string
			for _, value := range live.Transformation {
				if value.Name == entity.Transformation.Name {
					// unset props
					value.Fingerprint = ""

					original, err = ToYalm(value)
					if err != nil {
						return "", err
					}
					break
				}
			}

			d, err := ToDiff(original, changed)
			if err != nil {
				return "", err
			}
			res = append(res, d...)
		}
	}
	return strings.Join(res, "  ---\n"), nil
}

type Marshelable interface {
	MarshalJSON() ([]byte, error)
}

func ToDiff(original, changed string) ([]string, error) {
	yamls1, err := yamldiff.Load(original)
	if err != nil {
		return nil, err
	}

	yamls2, err := yamldiff.Load(changed)
	if err != nil {
		return nil, err
	}

	var ret []string
	for _, d := range yamldiff.Do(yamls1, yamls2) {
		if value := d.Dump(); len(value) > 4 {
			ret = append(ret, value)
		}
	}
	return ret, nil
}

func ToYalm(m Marshelable) (string, error) {
	if m == nil {
		return "", nil
	}

	out, err := m.MarshalJSON()
	if err != nil {
		return "", err
	}

	var ret bytes.Buffer
	if err := json2yaml.Convert(&ret, bytes.NewBuffer(out)); err != nil {
		return "", err
	}

	return ret.String(), nil
}
