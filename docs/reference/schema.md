# JSON schema

The full machine-readable schema for `sesamy.yaml` is published at:

```
https://raw.githubusercontent.com/foomo/sesamy-cli/refs/heads/main/sesamy.schema.json
```

It is generated from the Go struct definitions in `pkg/config/` (see `main_test.go` in the repo). The schema is updated on every release that changes a config field.

## Editor integration

Add a `yaml-language-server` directive at the top of your `sesamy.yaml`:

```yaml
# yaml-language-server: $schema=https://raw.githubusercontent.com/foomo/sesamy-cli/refs/heads/main/sesamy.schema.json
version: '1.1'
```

This unlocks autocomplete, hover docs, and inline validation in:

- VS Code (with the [YAML extension](https://marketplace.visualstudio.com/items?itemName=redhat.vscode-yaml))
- JetBrains IDEs (built-in)
- Neovim (`yamlls`)
- Helix, Zed, Sublime LSP, ...

## Pinning a version

If you want to lock the schema to a tagged release:

```yaml
# yaml-language-server: $schema=https://raw.githubusercontent.com/foomo/sesamy-cli/refs/tags/v0.18.0/sesamy.schema.json
```

Replace the tag with the version of `sesamy-cli` you are running.

## Regenerating the schema

When contributing changes to the config structs:

```bash
make generate
```

This refreshes `sesamy.schema.json` from the Go types. Commit the updated file alongside your code change.
