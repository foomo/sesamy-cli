# Introduction

`sesamy-cli` is a command-line tool that brings Google Tag Manager (GTM) under version control. Instead of clicking through the GTM UI to wire up tags, triggers, and variables for every analytics or advertising provider you use, you describe the desired state once in a `sesamy.yaml` file and let the CLI reconcile your GTM containers to match.

## What it does

- **Provisions** both **web** and **server-side** GTM containers from a single YAML file.
- **Diffs** local config against live GTM state so you can preview changes before applying them.
- **Lists** resources in a container — useful for audits and for spotting drift.
- **Generates TypeScript** event type definitions from Go event models (via [`sesamy-go`](https://github.com/foomo/sesamy-go)) so your frontend stays in sync with the tags you ship.
- **Opens** the relevant GTM and Google Analytics consoles for your configured account.

## Supported providers

The CLI ships built-in configuration for these providers, each wired with sensible defaults and Google Consent Mode integration where applicable:

| Category   | Providers |
|------------|-----------|
| Analytics  | Google Analytics, Mixpanel, Umami, Hotjar |
| Advertising | Google Ads, Facebook (Conversions API), Pinterest, Microsoft Ads, Criteo, Emarsys, Tracify |
| Consent / linking | Cookiebot CMP, Conversion Linker, Google Consent |
| Tagging   | Google Tag, Google Tag Manager |

See the [Providers](/providers/) section for per-provider config.

## How it fits together

```
┌──────────────┐    sesamy provision web      ┌──────────────────┐
│ sesamy.yaml  │ ───────────────────────────▶ │  GTM web         │
│              │    sesamy provision server   │  container       │
│              │ ───────────────────────────▶ │                  │
└──────┬───────┘                              └──────────────────┘
       │
       │ sesamy typescript
       ▼
┌──────────────────┐
│ Typed TS events  │ ───▶ used by your frontend / sesamy-go consumers
└──────────────────┘
```

`sesamy-cli` talks to GTM via the Google Tag Manager API v2 using a Service Account. Quota is throttled and cached entities are tracked so unused resources can be detected.

## When to use it

Use `sesamy-cli` if you:

- Run server-side GTM and want the same source-of-truth treatment for it as for your application code.
- Manage multiple providers and want their tags, triggers, and variables generated rather than hand-built.
- Need a reviewable diff before pushing changes to a shared container.
- Generate event types from Go and want GTM, your server, and your frontend to agree on the schema.

Continue to [Installation](./installation) to set it up.
