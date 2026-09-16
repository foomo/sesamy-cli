# Emarsys

Emarsys server-side integration.

## Configuration

```yaml
emarsys:
  enabled: true
  merchantId: ''
  templates:
    webExtendTag: ''
    initializationClient: ''
    initializationTag: ''
  googleConsent:
    enabled: true
    mode: analytics_storage
  serverContainer:
    directory: .
    packages:
      - path: github.com/foomo/sesamy-go/pkg/event
        types:
          - Purchase
          - ViewItem
          - ViewItemList
```

| Field | Purpose |
|-------|---------|
| `merchantId` | Emarsys merchant ID. |
| `googleConsent.mode` | Defaults to `analytics_storage`. |
| `serverContainer.packages` | Events to forward to Emarsys. |
| `templates.webExtendTag` | Path to a custom web extend tag template file. Empty (default) provisions the embedded template. |
| `templates.initializationClient` | Path to a custom initialization client template file. Empty (default) provisions the embedded template. |
| `templates.initializationTag` | Path to a custom initialization tag template file. Empty (default) provisions the embedded template. |
