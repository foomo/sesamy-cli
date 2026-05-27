# open

Open the relevant GTM or Google Analytics console for the configured account in your default browser.

```bash
sesamy open <target> [-c sesamy.yaml ...]
```

## Targets

| Target | Opens |
|--------|-------|
| `ga` | Google Analytics property dashboard. Requires `googleAnalytics.propertyId`. |
| `gtm-web` | The web container in GTM. |
| `gtm-server` | The server container in GTM. |

## Example

```bash
sesamy open gtm-web
sesamy open ga
```

## Flags

| Flag | Default | Purpose |
|------|---------|---------|
| `-c, --config` | `sesamy.yaml` | Config file(s). |
