# list

Print the resources currently in a GTM container's selected workspace.

```bash
sesamy list web [-c sesamy.yaml ...]
sesamy list server [-c sesamy.yaml ...]
```

## What it does

Fetches and tabulates tags, triggers, variables, templates, and folders. Each entity is shown alongside the provider that owns it. Entities the CLI doesn't recognize are flagged as **unmanaged** so you can decide whether to clean them up or add them to your config.

## Flags

| Flag | Default | Purpose |
|------|---------|---------|
| `-c, --config` | `sesamy.yaml` | Config file(s) — needed to know which account / container to query. |
| `-v, --verbose` | `false` | Debug logging. |

## Use cases

- **Audit** an existing container before adopting `sesamy-cli`.
- **Drift detection** — combine with `sesamy diff` to spot manual changes.
- **Cleanup** — find leftover entities from removed providers.
