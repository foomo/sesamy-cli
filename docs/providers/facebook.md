# Facebook

[Facebook Conversions API](https://developers.facebook.com/docs/marketing-api/conversions-api/guides/gtm-server-side) — server-side tagging only.

## Configuration

```yaml
facebook:
  enabled: true
  pixelId: ''
  apiAccessToken: ''
  testEventToken: ''
  googleConsent:
    enabled: true
    mode: ad_storage
  serverContainer:
    directory: .
    packages:
      - path: github.com/foomo/sesamy-go/pkg/event
        types:
          - AddPaymentInfo
          - AddToCart
          - AddToWishlist
          - PageView
          - Purchase
          - Search
          - BeginCheckout
          - GenerateLead
          - ViewItem
```

| Field | Purpose |
|-------|---------|
| `pixelId` | Facebook Pixel ID. |
| `apiAccessToken` | Conversions API access token (server-side). |
| `testEventToken` | Optional — use Meta's test events tool to verify the payload. |
| `googleConsent.mode` | Defaults to `ad_storage`. |
| `serverContainer.packages` | Events to send via Conversions API. |

## Test events

Set `testEventToken` to the value from Meta Events Manager's **Test Events** tab. With that set, sent events appear under "Test events" rather than counting toward your live data. Remove the token before going live.
