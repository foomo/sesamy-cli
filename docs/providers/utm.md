# UTM

UTM attribution for the server container. Persists the incoming UTM parameters in a cookie and exposes them as a variable other providers can consume.

## Configuration

```yaml
utm:
  enabled: true
  templates:
    attributionVariable: ''
    cookieWriterTag: ''
  googleConsent:
    enabled: true
    mode: ad_storage
```

| Field | Purpose |
|-------|---------|
| `googleConsent.mode` | Defaults to `ad_storage`. |
| `templates.attributionVariable` | Path to a custom attribution variable template file. Empty (default) provisions the embedded template. |
| `templates.cookieWriterTag` | Path to a custom cookie writer tag template file. Empty (default) provisions the embedded template. |

## What gets provisioned

| Entity | Name |
|--------|------|
| Variable template | `UTM Attribution` |
| Variable | `UTM Attribution` |
| Tag template | `UTM Attribution Cookie Writer` |
| Trigger | `UTM Attribution Cookie` (all events) |
| Tag | `UTM Attribution Cookie` |

The cookie writer tag fires on every event and stores the attribution data; the `UTM Attribution` variable reads it back.

## Dependants

[Mixpanel](./mixpanel) track tags send the `UTM Attribution` variable as the `utmAttribution` parameter. Enable and provision `utm` before provisioning Mixpanel, otherwise provisioning fails with a lookup error.
