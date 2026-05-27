# provision

Reconcile a GTM container to match your `sesamy.yaml`.

```bash
sesamy provision web [-c sesamy.yaml ...]
sesamy provision server [-c sesamy.yaml ...]
```

## What it does

For each enabled provider, the CLI:

1. Reads the desired tags, triggers, variables, and templates from your config.
2. Fetches the current state of the selected workspace via the GTM API.
3. Creates entities that don't exist, updates entities that have drifted, and leaves untouched anything that already matches.

The CLI tracks every entity it accesses so it can detect resources it did **not** create — those are reported but never modified.

## Flags

| Flag | Default | Purpose |
|------|---------|---------|
| `-c, --config` | `sesamy.yaml` | One or more config files (later files override earlier). |
| `-v, --verbose` | `false` | Debug logging — useful when diagnosing quota or API errors. |

## Tips

- Run `sesamy diff` first. It uses the same reconciliation logic but never writes.
- Bump `googleApi.requestQuota` if you hit `429 Too Many Requests`.
- Provision against a non-default workspace (`googleTagManager.webContainer.workspaceId`) when you want a human to publish from the GTM UI.

## Exit codes

| Code | Meaning |
|------|---------|
| `0` | All changes applied successfully (or nothing needed to change). |
| `non-zero` | Config invalid, credentials missing, or a GTM API call failed. |
