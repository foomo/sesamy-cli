# Conversion Linker

Google's conversion linker preserves ad-click identifiers (`gclid`, `gbraid`, `wbraid`) across pageviews so downstream conversion tags can attribute correctly.

## Configuration

```yaml
conversionLinker:
  enabled: true
  googleConsent:
    enabled: true
    mode: ad_storage
  enableLinkerParams: true          # accept incoming linker params
```

| Field | Purpose |
|-------|---------|
| `enabled` | Provision the conversion linker tag. |
| `googleConsent` | Gate the linker on `ad_storage` consent. |
| `enableLinkerParams` | Accept linker params arriving via cross-domain redirects. |

::: tip
If you use Google Ads conversion tracking, leave this enabled. Without it, attribution silently degrades.
:::
