# Mixpanel

Mixpanel server-side event tracking.

## Configuration

```yaml
mixpanel:
  enabled: true
  projectToken: ''
  googleConsent:
    enabled: true
    mode: analytics_storage
  serverContainer:
    track:
      directory: .
      packages:
        - path: github.com/foomo/sesamy-go/pkg/event
          types:
            - AddPaymentInfo
            - AddShippingInfo
            - AddToCart
            - BeginCheckout
            - PageView
            - Purchase
            - RemoveFromCart
            - Search
            - SelectItem
            - ViewCart
            - ViewItem
            - ViewItemList
```

| Field | Purpose |
|-------|---------|
| `projectToken` | Mixpanel project token. |
| `googleConsent.mode` | Defaults to `analytics_storage`. |
| `serverContainer.track.packages` | Events to forward to Mixpanel's `/track` endpoint. |
