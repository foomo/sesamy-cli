# Configuration reference

The full top-level schema of `sesamy.yaml` (version `1.1`). For provider-specific blocks see the [Providers](/providers/) section.

## Top-level fields

| Field | Type | Purpose |
|-------|------|---------|
| `version` | string | Schema version. Currently `1.1`. |
| `redactVisitorIp` | bool | Globally redact visitor IPs in collected events. |
| `enableGeoResolution` | bool | Enable [region-specific settings](https://developers.google.com/tag-platform/tag-manager/server-side/enable-region-specific-settings). |
| `googleApi` | object | Service account credentials and API quota. |
| `googleTagManager` | object | GTM account / container IDs and shared container variables. |
| `googleTag` | object | The site-wide Google Tag (gtag) and TypeScript generation. |
| `googleAnalytics` | object | [GA4 provider config](/providers/google-analytics). |
| `googleAds` | object | [Google Ads provider config](/providers/google-ads). |
| `conversionLinker` | object | [Conversion Linker provider config](/providers/conversion-linker). |
| `facebook` | object | [Facebook provider config](/providers/facebook). |
| `pinterest` | object | [Pinterest provider config](/providers/pinterest). |
| `microsoftAds` | object | [Microsoft Ads provider config](/providers/microsoft-ads). |
| `criteo` | object | [Criteo provider config](/providers/criteo). |
| `emarsys` | object | [Emarsys provider config](/providers/emarsys). |
| `mixpanel` | object | [Mixpanel provider config](/providers/mixpanel). |
| `umami` | object | [Umami provider config](/providers/umami). |
| `tracify` | object | [Tracify provider config](/providers/tracify). |
| `hotjar` | object | [Hotjar provider config](/providers/hotjar). |
| `cookiebot` | object | [Cookiebot provider config](/providers/cookiebot). |

## googleApi

```yaml
googleApi:
  credentials: '{...}'         # inline service account JSON (single line)
  credentialsFile: ./creds.json # OR path to the JSON file
  requestQuota: 15              # GTM API requests per minute
```

Use `credentialsFile` for local dev, `credentials` for CI. The CLI throttles itself to `requestQuota` r/m.

## googleTagManager

```yaml
googleTagManager:
  accountId: '6099238525'
  webContainer:
    tagId: GTM-XXXXXXXX
    containerId: '175355532'
    workspaceId: '23'            # optional
    workspace: 'Default'         # optional, by name
  serverContainer:
    tagId: GTM-YYYYYYYY
    containerId: '175348980'
    workspaceId: '10'
    workspace: 'Default'
  webContainerVariables:
    dataLayer: [link_url]
    lookupTables:
      link_url_conversion_label:
        input: '{{dlv.link_url}}'
        valueTable:
          123456: 'https://foomo.org/'
  serverContainerVariables:
    eventData: [link_url]
    lookupTables: { ... }
```

| Field | Purpose |
|-------|---------|
| `accountId` | GTM account ID. |
| `<container>.tagId` | The `GTM-` container tag identifier. |
| `<container>.containerId` | Numeric container ID. |
| `<container>.workspaceId` / `workspace` | Optional workspace selector. Defaults to the default workspace. |
| `<container>Variables.dataLayer` | Container-level DataLayer variables (web) or event data (server). |
| `<container>Variables.lookupTables` | Reusable GTM lookup tables. |

## googleTag

Documented in [Google Tag provider](/providers/google-tag).

## Schema validation

Point your editor at the JSON schema:

```yaml
# yaml-language-server: $schema=https://raw.githubusercontent.com/foomo/sesamy-cli/refs/heads/main/sesamy.schema.json
version: '1.1'
```

VS Code (YAML extension), JetBrains IDEs (built-in), and Neovim (`yamlls`) will autocomplete and validate inline.

## Multiple files

```bash
sesamy provision web -c base.yaml -c env/prod.yaml
```

Files are merged in argument order. Later files override earlier ones. Run `sesamy config -c ...` to see the resolved result.
