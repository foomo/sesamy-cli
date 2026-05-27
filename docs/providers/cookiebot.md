# Cookiebot

Cookiebot CMP integration. The CLI does **not** install the Cookiebot tag template for you — install [Cookiebot CMP](https://tagmanager.google.com/gallery/#/owners/cybotcorp/templates/gtm-templates-cookiebot-cmp) from the GTM gallery first.

## Configuration

```yaml
cookiebot:
  enabled: true
  templateName: Cookiebot CMP     # must match the imported template's name
  cookiebotId: ''
  cdnRegion: eu                    # eu or com
  urlPassthrough: false
  advertiserConsentModeEnabled: false
  regionSettings:
    - region: ''                   # blank = global default
      preferences: denied          # functionality_storage, personalization_storage
      statistics: denied           # analytics_storage
      marketing: denied            # ad_storage
      adUserData: denied
      adPersonalization: denied
```

| Field | Purpose |
|-------|---------|
| `templateName` | Name of the manually-imported Cookiebot CMP template. Must match exactly. |
| `cookiebotId` | Your Cookiebot CBID. |
| `cdnRegion` | `eu` or `com` — pick the region matching your Cookiebot subscription. |
| `urlPassthrough` | Pass identifiers via URL parameters when consent is denied. |
| `advertiserConsentModeEnabled` | Enable advertiser consent mode signals. |
| `regionSettings[]` | Per-region default consent states. Leave `region` blank for the global default. |
