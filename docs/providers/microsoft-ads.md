# Microsoft Ads

Microsoft Ads (Bing UET) server-side conversion tracking.

## Configuration

```yaml
microsoftAds:
  enabled: true
  tagId: ''
  googleConsent:
    enabled: true
    mode: ad_storage
  conversion:
    enabled: true
    serverContainer:
      directory: .
      settings:
        purchase:
          - conversionLabel: ''
      packages:
        - path: github.com/foomo/sesamy-go/pkg/event
          types:
            - Purchase
```

| Field | Purpose |
|-------|---------|
| `tagId` | Microsoft UET tag ID. |
| `googleConsent.mode` | Defaults to `ad_storage`. |
| `conversion.serverContainer.settings.<event>` | List of `{ conversionLabel }` pairs — one tag per entry. |
| `conversion.serverContainer.packages` | Events to wire conversion tags for. |

::: tip
Field names are based on the current schema. Run `sesamy config -c sesamy.yaml` to see the resolved values your build will use.
:::
