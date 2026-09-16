# Criteo

Criteo OneTag, server-side.

## Configuration

```yaml
criteo:
  enabled: true
  callerId: 123
  partnerId: 123456
  applicationId: com.foomo
  templates:
    eventsApiTag: ''
    userIdentification: ''
  serverContainer:
    directory: .
    packages:
      - path: github.com/foomo/sesamy-go/pkg/event
        types:
          - AddToCart
          - BeginCheckout
          - PageView
          - Purchase
          - ViewItem
          - ViewItemList
          - ViewCart
```

| Field | Purpose |
|-------|---------|
| `callerId` | Criteo caller ID issued for your integration. |
| `partnerId` | Criteo partner ID. |
| `applicationId` | Application identifier (often the reversed-domain form). |
| `serverContainer.packages` | Events to forward to Criteo. |
| `templates.eventsApiTag` | Path to a custom Events API tag template file. Empty (default) looks up the manually installed template. |
| `templates.userIdentification` | Path to a custom User Identification template file. Empty (default) looks up the manually installed template. |
