# Google Tag

The Google Tag (`gtag`) is the foundational tag that loads on every page and routes events to GA4, Google Ads, and any other Google product you've connected. `sesamy-cli` provisions this in the **web** container.

## Configuration

```yaml
googleTag:
  tagId: G-PZ5ELRCR31         # your GA4 / Google Tag measurement ID
  sendPageView: true          # emit a page_view on initial load
  serverContainerUrl: ''      # optional: route gtag through your server container
  typeScript:
    outputPath: ./web/src/sesamy
    directory: .
    packages:
      - path: github.com/foomo/sesamy-go/pkg/event
        types:
          - PageView
          - AddToCart
          - Purchase
          # ... see sesamy-go for the full GA4 event catalog
```

| Field | Purpose |
|-------|---------|
| `tagId` | Google Tag / GA4 measurement ID. Required. |
| `sendPageView` | If `true`, fires a `page_view` automatically on initial page load. |
| `serverContainerUrl` | Optional custom server container URL (e.g. your first-party tagging domain). |
| `typeScript` | Controls `sesamy typescript` output. See [`sesamy typescript`](/commands/typescript). |

The full GA4 event list — automatically collected and recommended — is supported by [`sesamy-go`](https://github.com/foomo/sesamy-go). See the README there for the catalog.
