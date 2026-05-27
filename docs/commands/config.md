# config

Print the merged, effective configuration after layering all `-c` files.

```bash
sesamy config [-c sesamy.yaml ...] [-c sesamy.prod.yaml]
```

## What it does

Loads each `--config` file in order, merges them, validates against the JSON schema, and prints the resolved YAML. Use it to:

- Verify what a multi-file config actually resolves to before provisioning.
- Debug environment-specific overrides.
- Pipe into `diff` against a known-good config to spot drift.

## Flags

| Flag | Default | Purpose |
|------|---------|---------|
| `-c, --config` | `sesamy.yaml` | Config file(s). |
| `-v, --verbose` | `false` | Include debug-level fields. |

::: warning
The output includes any credentials inlined under `googleApi.credentials`. Don't paste it into shared channels or tickets without redacting.
:::
