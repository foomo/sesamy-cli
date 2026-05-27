# Emarsys

Emarsys server-side integration.

## Configuration

```yaml
emarsys:
  enabled: true
  merchantId: ''
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
