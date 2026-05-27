# Google Analytics

GA4 web + server-side tagging.

## Configuration

```yaml
googleAnalytics:
  enabled: true
  propertyId: ''                  # used by `sesamy open ga`
  googleConsent:
    enabled: true
    mode: analytics_storage
  googleGTagJSOverride:
    enabled: true                 # provision a custom gtag.js client
    priority: 10
    ecommerceItems: true          # patch ecommerce items in transit
  webContainer:
    directory: .
    packages:
      - path: github.com/foomo/sesamy-go/pkg/event
        types:
          - PageView
          - SelectItem
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
| `enabled` | Enable / disable the provider. |
| `propertyId` | GA4 property ID. Required by `sesamy open ga`. |
| `googleConsent` | Google Consent Mode integration. |
| `googleGTagJSOverride.enabled` | Provision a custom `gtag.js` client tag with higher priority. |
| `googleGTagJSOverride.priority` | GTM tag priority — higher fires earlier. |
| `googleGTagJSOverride.ecommerceItems` | Normalize ecommerce items in the override. |
| `webContainer.packages` | Events handled in the web container. |
| `serverContainer.packages` | Events handled in the server container. |
