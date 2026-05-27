# Providers

Each provider is a self-contained module under `pkg/provider/` that knows how to wire its tags, triggers, and variables into both the web and server GTM containers (where applicable). You enable them with `enabled: true` in the relevant section of `sesamy.yaml`.

## Catalog

| Provider | Web | Server | Consent | Server events |
|----------|:---:|:------:|:-------:|---------------|
| [Google Tag](./google-tag) | ✔ | — | — | — |
| [Google Analytics](./google-analytics) | ✔ | ✔ | analytics | PageView, SelectItem, ... |
| [Google Ads](./google-ads) | ✔ | ✔ | ad | AddToCart, Purchase |
| [Conversion Linker](./conversion-linker) | ✔ | — | ad | — |
| [Facebook](./facebook) | — | ✔ | ad | AddToCart, Purchase, ... |
| [Pinterest](./pinterest) | — | ✔ | analytics | AddToCart, Purchase, ... |
| [Microsoft Ads](./microsoft-ads) | — | ✔ | ad | conversions |
| [Criteo](./criteo) | — | ✔ | — | AddToCart, Purchase, ... |
| [Emarsys](./emarsys) | — | ✔ | analytics | Purchase, ViewItem, ... |
| [Mixpanel](./mixpanel) | — | ✔ | analytics | PageView, Purchase, ... |
| [Umami](./umami) | — | ✔ | analytics | PageView, SelectItem |
| [Tracify](./tracify) | — | ✔ | analytics | AddToCart, Purchase, ... |
| [Hotjar](./hotjar) | ✔ | — | — | — |
| [Cookiebot](./cookiebot) | ✔ | — | CMP | — |

> "Consent" indicates the Google Consent Mode signal the provider respects by default. You can override per provider.

## Shape of a provider config

Most providers share the same shape:

```yaml
<providerName>:
  enabled: true
  # provider-specific identifiers (pixel IDs, tokens, ...)
  googleConsent:
    enabled: true
    mode: ad_storage          # or analytics_storage
  serverContainer:            # if the provider supports server-side
    directory: .
    packages:
      - path: github.com/foomo/sesamy-go/pkg/event
        types: [PageView, AddToCart, Purchase]
```

- **`enabled`** toggles the provider entirely. Disabled providers are not provisioned.
- **`googleConsent`** wires consent mode for the provider's tags. Disable it if your consent flow is handled differently.
- **`serverContainer.packages`** lists the `sesamy-go` (or your own) event types this provider should handle.

## Adding a provider

1. Set `enabled: true` and fill in the identifiers.
2. Add the event types you care about under `serverContainer.packages` (and `webContainer.packages` for providers that have a web component).
3. `sesamy diff` → review → `sesamy provision`.
