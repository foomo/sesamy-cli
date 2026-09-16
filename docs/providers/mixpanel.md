# Mixpanel

Mixpanel server-side event tracking.

Track tags send UTM attribution data, so the [UTM](./utm) provider must be enabled and provisioned first.

## Configuration

```yaml
mixpanel:
  enabled: true
  projectToken: ''
  templates:
    tag: ''
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
| `templates.tag` | Path to a custom tag template file. Empty (default) provisions the embedded template. |
| `googleConsent.mode` | Defaults to `analytics_storage`. |
| `serverContainer.userId` | Event-data property holding the user ID, forwarded to Mixpanel. |
| `serverContainer.track.packages` | Events to forward to Mixpanel's `/track` endpoint. |

## Tag template

By default the `Mixpanel` tag template is provisioned by the CLI — no manual template install required.

Set `templates.tag` to the path of a template file to provision your own instead:

```yaml
mixpanel:
  templates:
    tag: ./templates/mixpanel.tpl
```

The file is uploaded verbatim as the template named `Mixpanel`, so it must be a complete GTM template
(`___INFO___`, `___TEMPLATE_PARAMETERS___`, `___SANDBOXED_JS_FOR_SERVER___`, …). Use the
[Mixpanel](https://github.com/stape-io/mixpanel-tag) template by stape-io as a starting point.
