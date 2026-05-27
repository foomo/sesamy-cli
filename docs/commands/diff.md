# diff

Show what `sesamy provision` would change, without writing anything.

```bash
sesamy diff web [-c sesamy.yaml ...]
sesamy diff server [-c sesamy.yaml ...]
```

## What it does

Runs the full reconciliation against your live GTM workspace and prints a per-entity report:

- **Create** — entity is missing in GTM.
- **Update** — entity exists but its definition has drifted.
- **No change** — local matches remote.
- **Unmanaged** — entity exists in GTM but isn't described in your config.

No API mutations are made. Read-only API calls still count against quota.

## Flags

| Flag | Default | Purpose |
|------|---------|---------|
| `-c, --config` | `sesamy.yaml` | Config file(s). |
| `-v, --verbose` | `false` | Debug logging. |

## CI integration

Run `sesamy diff` on every PR that touches `sesamy.yaml`. A non-empty diff before merge means the change is real and previewable; a non-empty diff on `main` between deploys means someone hand-edited GTM.

## Exit codes

| Code | Meaning |
|------|---------|
| `0` | Diff produced (may be empty or non-empty — both are non-error). |
| `non-zero` | Could not produce a diff (config invalid, API error, etc.). |

::: tip
The command itself does not exit non-zero when changes are pending. If you want CI to fail on drift, pipe the output and grep for change markers, or wire up your own check.
:::
