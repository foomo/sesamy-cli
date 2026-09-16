# OpenAI Ads

OpenAI Ads server-side conversion tracking via the Conversions API.

::: warning
This provider relies on the **OpenAI Ads Conversions API by Stape** tag template. Install it into your server
container from the GTM Community Template Gallery before provisioning, or point `templates.conversionsApiTag` at a
local copy.
:::

## Configuration

```yaml
openAiAds:
  enabled: true
  pixelId: ''
  apiKey: ''
  templates:
    conversionsApiTag: ''
  googleConsent:
    enabled: true
    mode: ad_storage
  conversion:
    enabled: true
    validateOnly: false
    serverContainer:
      directory: .
      settings:
        page_view:
          eventName: page_viewed
      packages:
        - path: github.com/foomo/sesamy-go/pkg/event
          types:
            - Purchase
```

| Field | Purpose |
|-------|---------|
| `pixelId` | OpenAI Ads pixel ID. |
| `apiKey` | OpenAI Ads API key. |
| `googleConsent.mode` | Defaults to `ad_storage`. |
| `conversion.validateOnly` | Validate events without persisting them. Applies to every event. |
| `conversion.serverContainer.settings.<event>.eventName` | Standard OpenAI Ads event name. Defaults to the event name. |
| `conversion.serverContainer.packages` | Events to wire conversion tags for. |
| `templates.conversionsApiTag` | Path to a custom conversions api tag template file. Empty (default) uses the manually installed one. |

::: tip
Field names are based on the current schema. Run `sesamy config -c sesamy.yaml` to see the resolved values your build will use.
:::
