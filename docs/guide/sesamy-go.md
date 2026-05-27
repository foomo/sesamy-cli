# sesamy-go integration

`sesamy-cli` is one half of a pair. The other is [`sesamy-go`](https://github.com/foomo/sesamy-go), a Go library that defines the **event model** — strongly typed structs for `PageView`, `AddToCart`, `Purchase`, and so on.

## The contract

`sesamy-go` owns event **types**. `sesamy-cli` reads those types and:

1. **Generates the GTM entities** (tags, triggers, variables) that handle each event.
2. **Generates TypeScript definitions** so your frontend can emit those events with the same shape.

You point the CLI at the Go package containing your event structs by listing them under each provider's `packages` block:

```yaml
googleTag:
  typeScript:
    outputPath: ./web/src/sesamy
    directory: .
    packages:
      - path: github.com/foomo/sesamy-go/pkg/event
        types:
          - PageView
          - AddToCart
          - Purchase
```

`directory` is the path to a `go.mod` whose dependency graph includes the listed packages. Usually the same repo as your application code.

## Per-provider packages

Each provider that consumes events (Google Analytics, Facebook, Pinterest, Mixpanel, ...) has its own `packages` block. Only the event types you list there are wired up for that provider.

```yaml
facebook:
  enabled: true
  pixelId: ''
  apiAccessToken: ''
  serverContainer:
    directory: .
    packages:
      - path: github.com/foomo/sesamy-go/pkg/event
        types:
          - AddToCart
          - Purchase
          - ViewItem
```

This means a Facebook tag, trigger, and matching variables are generated for `AddToCart`, `Purchase`, and `ViewItem` — and nothing else.

## Custom event types

You can mix in your own event types. As long as the Go package lives in (or is a dependency of) the module pointed to by `directory`, list it under `packages`:

```yaml
packages:
  - path: github.com/myorg/myapp/pkg/events
    types:
      - CustomCheckoutStep
```

The struct shape determines the generated DataLayer variables and the TypeScript output.

## Generating TypeScript

```bash
sesamy typescript
```

This reads `googleTag.typeScript.packages`, runs [`gocontemplate`](https://github.com/foomo/gocontemplate) over the listed Go types, and writes `.ts` files to `googleTag.typeScript.outputPath`. Commit the output — your frontend imports from it.

## Standard event catalog

The full list of GA4-recommended events supported by `sesamy-go` is in its README. Common ones:

- `PageView`, `ScreenView`, `UserEngagement`
- `AddToCart`, `RemoveFromCart`, `ViewCart`, `BeginCheckout`, `AddPaymentInfo`, `AddShippingInfo`, `Purchase`, `Refund`
- `ViewItem`, `ViewItemList`, `SelectItem`, `SelectPromotion`, `ViewPromotion`
- `Login`, `SignUp`, `Share`, `Search`, `SelectContent`
- `GenerateLead`, `QualifyLead`, `WorkingLead`, `CloseConvertLead`
