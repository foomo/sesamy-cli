# Quick start

This walkthrough takes you from zero to a provisioned GTM web container in a few minutes.

## Prerequisites

- A Google Tag Manager **account** with both a **web** and **server** container created.
- A Google Cloud **Service Account** with access to the GTM API. See [Google API setup](/reference/google-api).
- The service account email added as a **user** of your GTM account (with publish rights).

## 1. Create `sesamy.yaml`

In your project root, create a minimal `sesamy.yaml`:

```yaml
# yaml-language-server: $schema=https://raw.githubusercontent.com/foomo/sesamy-cli/refs/heads/main/sesamy.schema.json
version: '1.1'

redactVisitorIp: true
enableGeoResolution: true

googleApi:
  credentialsFile: ./google_service_account_creds.json
  requestQuota: 15

googleTagManager:
  accountId: '6099238525'
  webContainer:
    tagId: GTM-XXXXXXXX
    containerId: '175355532'
  serverContainer:
    tagId: GTM-YYYYYYYY
    containerId: '175348980'

googleTag:
  tagId: G-XXXXXXXXXX
  sendPageView: true

googleAnalytics:
  enabled: true
  googleConsent:
    enabled: true
    mode: analytics_storage
```

The `$schema` comment gives you autocomplete and validation in editors with YAML language server support (VS Code, JetBrains, Neovim).

## 2. Preview the diff

Before touching GTM, see what `sesamy` plans to change:

```bash
sesamy diff web
sesamy diff server
```

The first run will show every tag, trigger, and variable as **created**.

## 3. Provision

Apply the changes:

```bash
sesamy provision web
sesamy provision server
```

This creates / updates entities in the configured workspace. Rerun `sesamy diff` afterward — it should report no changes.

## 4. Inspect

```bash
sesamy list web
sesamy list server
```

Lists every tag, trigger, and variable managed under the workspace, including ones the CLI considers unused.

## 5. Generate TypeScript events (optional)

If you use [`sesamy-go`](https://github.com/foomo/sesamy-go) to define your event model, generate TypeScript types:

```bash
sesamy typescript
```

Output goes to the `outputPath` configured under `googleTag.typeScript`. See the [sesamy-go integration](./sesamy-go) guide for details.

## Next steps

- Browse the full [Configuration](./configuration) reference.
- Enable additional [Providers](/providers/) one at a time and re-provision.
- Wire `sesamy diff` into CI to catch drift.
