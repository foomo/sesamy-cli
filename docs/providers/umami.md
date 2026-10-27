# Umami

Self-hosted, privacy-friendly analytics via Umami's API.

## Configuration

```yaml
umami:
  enabled: true
  domain: your-domain.com
  websiteId: ''
  endpointUrl: https://umami.your-domain.com
  templates:
    tag: ''
  googleConsent:
    enabled: true
    mode: analytics_storage
  serverContainer:
    directory: .
    packages:
      - path: github.com/foomo/sesamy-go/pkg/event
        types:
          - PageView
          - SelectItem
```

| Field | Purpose |
|-------|---------|
| `domain` | Optional override for the event's domain field. |
| `websiteId` | Umami website ID. |
| `endpointUrl` | Base URL of your Umami instance. |
| `googleConsent.mode` | Defaults to `analytics_storage`. |
| `serverContainer.packages` | Events to forward. |
| `templates.tag` | Path to a custom tag template file. Empty (default) provisions the embedded template. |

## Event data

Each tag sends the `Params` fields of its event type as Umami event data (`payload.data`).
The tag reads the values by their JSON name from the server container event data. It omits missing and empty values.
Events without `Params` fields send no event data.

## Visitor IP

The tag sends the first address of the `ip_override` event data value as `payload.ip`.
This requires Umami v2.10.0 or later. If `payload.ip` is set, Umami v2.17.0 and later resolve the location from the IP and ignore the `cf-ip*` location headers.
