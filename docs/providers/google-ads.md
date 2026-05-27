# Google Ads

Remarketing + conversion tracking, web and server.

## Configuration

```yaml
googleAds:
  enabled: true
  conversionId: ''                # Google Ads Conversion Tracking ID
  googleConsent:
    enabled: true
    mode: ad_storage
  remarketing:
    enabled: true
    enableConversionLinker: true
  conversion:
    enabled: true
    serverContainer:
      directory: .
      settings:
        add_to_cart:
          - label: ''
          - conversionId: ''
            label: ''
        purchase:
          - label: ''
          - conversionId: ''
            label: ''
      packages:
        - path: github.com/foomo/sesamy-go/pkg/event
          types:
            - AddToCart
            - Purchase
```

| Field | Purpose |
|-------|---------|
| `conversionId` | Default Google Ads conversion ID. |
| `remarketing.enabled` | Provision the remarketing tag. |
| `remarketing.enableConversionLinker` | Enable conversion linker integration. |
| `conversion.enabled` | Provision conversion tracking tags. |
| `conversion.serverContainer.settings` | Per-event list of `{ conversionId?, label }` pairs. The CLI generates a conversion tag per pair. |
| `conversion.serverContainer.packages` | Event types to wire conversion tags for. |

## Multi-account conversions

`settings.<event>` is a list, so you can fire the same event into multiple Ads accounts by adding more `{ conversionId, label }` entries.
