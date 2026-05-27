# Configuration

`sesamy-cli` reads a single YAML file (default: `sesamy.yaml`) that describes both your web and server GTM containers and every provider you want enabled. The file is validated against [`sesamy.schema.json`](https://github.com/foomo/sesamy-cli/blob/main/sesamy.schema.json) — point your editor at it with the `# yaml-language-server: $schema=...` comment.

## File layout

```yaml
version: '1.1'           # schema version
redactVisitorIp: true    # global GTM setting
enableGeoResolution: true

googleApi: { ... }       # service account credentials + quota
googleTagManager: { ... } # account / container IDs and shared variables
googleTag: { ... }        # tag ID, page-view defaults, typescript output

# Providers (each can be enabled: false)
googleAnalytics: { ... }
googleAds: { ... }
facebook: { ... }
pinterest: { ... }
microsoftAds: { ... }
criteo: { ... }
emarsys: { ... }
mixpanel: { ... }
umami: { ... }
tracify: { ... }
hotjar: { ... }
cookiebot: { ... }
conversionLinker: { ... }
```

See [Reference → Configuration](/reference/configuration) for the field-by-field schema.

## Multiple configs

The `--config` (`-c`) flag accepts multiple files and merges them in order. Later files override earlier ones:

```bash
sesamy provision web -c base.yaml -c production.yaml
```

This lets you keep secret-free defaults checked in and layer environment-specific values on top via CI.

## Credentials

`googleApi.credentialsFile` points to a Service Account JSON file, or you can inline the JSON as a single-line string in `googleApi.credentials`. Inline is convenient for CI; file paths are convenient for local development. Both are equivalent at runtime.

::: warning
Never commit credential files. Add `google_service_account_creds.json` (or whatever you call it) to `.gitignore`.
:::

## Quotas

GTM API quota is **15 requests per minute** by default. Provisioning a fully-loaded container easily exceeds that. Either:

- Request a higher quota in the Google Cloud console for the service account's project, then set `googleApi.requestQuota` to the new value.
- Leave it at `15` and let the CLI throttle. Larger configs will simply take a few minutes.

See [Caveats](#caveats) below.

## Workspaces

`googleTagManager.webContainer.workspaceId` (or `workspace`, by name) selects a non-default workspace. Useful for staging changes in a separate workspace before publishing.

## Container variables

`googleTagManager.webContainerVariables` and `serverContainerVariables` declare DataLayer variables and lookup tables that providers can reference. Define them once at the container level instead of repeating them per provider.

```yaml
googleTagManager:
  webContainerVariables:
    dataLayer:
      - link_url
    lookupTables:
      link_url_conversion_label:
        input: '{{dlv.link_url}}'
        valueTable:
          123456: 'https://foomo.org/'
```

## Caveats

- The default GTM API quota (15 r/m) is low. Bump it for any non-trivial config.
- The CLI manages entities it creates. Hand-edited tags from the GTM UI are left alone but reported by `sesamy list` and `diff` so you can audit drift.
- Provisioning publishes changes to the selected workspace but does **not** create a GTM version or publish to live. Use the GTM UI (or `sesamy open gtm-web`) to publish.
