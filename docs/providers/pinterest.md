# Pinterest

Server-side Pinterest Conversions API.

## Configuration

```yaml
pinterest:
  enabled: true
  advertiserId: ''
  apiAccessToken: ''
  testModeEnabled: false
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
          - AddToCart
          - GenerateLead
          - PageView
          - Purchase
          - Search
          - SignUp
          - ViewItemList
```

| Field | Purpose |
|-------|---------|
| `advertiserId` | Pinterest advertiser (tag) ID. |
| `apiAccessToken` | Pinterest API access token. |
| `testModeEnabled` | Send events with the `test` flag set — won't count toward live data. |
| `googleConsent.mode` | Defaults to `analytics_storage`. |
| `serverContainer.packages` | Events to send. |
| `templates.tag` | Path to a custom conversions tag template file. Empty (default) looks up the manually installed template. |
