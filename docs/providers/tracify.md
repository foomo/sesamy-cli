# Tracify

Tracify attribution tracking, server-side.

## Configuration

```yaml
tracify:
  enabled: true
  token: ''
  customerSiteId: ''
  googleConsent:
    enabled: true
    mode: analytics_storage
  serverContainer:
    directory: .
    packages:
      - path: github.com/foomo/sesamy-go/pkg/event
        types:
          - AddToCart
          - PageView
          - ViewItem
          - Purchase
```

| Field | Purpose |
|-------|---------|
| `token` | Tracify API token. |
| `customerSiteId` | Tracify customer site identifier. |
| `googleConsent.mode` | Defaults to `analytics_storage`. |
| `serverContainer.packages` | Events to forward. |
