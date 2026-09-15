# Mixpanel

Mixpanel server-side event tracking.

Track tags send UTM attribution data, so the [UTM](./utm) provider must be enabled and provisioned first.

## Configuration

```yaml
mixpanel:
  enabled: true
  projectToken: ''
  lookupTagTemplate: false
  googleConsent:
    enabled: true
    mode: analytics_storage
  serverContainer:
    userId: user_id
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
| `lookupTagTemplate` | When `true`, look up an existing `Mixpanel` tag template instead of provisioning the built-in one. Defaults to `false`. |
| `googleConsent.mode` | Defaults to `analytics_storage`. |
| `serverContainer.userId` | Event-data property holding the user ID, forwarded to Mixpanel. |
| `serverContainer.track.packages` | Events to forward to Mixpanel's `/track` endpoint. |

## Tag template

By default the `Mixpanel` tag template is provisioned by the CLI — no manual template install required.

Set `lookupTagTemplate: true` to use a template you manage yourself instead. The CLI then looks up a
template named `Mixpanel` in the server container and fails if it is missing — install the
[Mixpanel](https://github.com/stape-io/mixpanel-tag) template by stape-io from the GTM gallery first.
