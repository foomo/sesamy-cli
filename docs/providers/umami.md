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
