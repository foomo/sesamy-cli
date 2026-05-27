# typescript

Generate TypeScript event type definitions from Go event structs.

```bash
sesamy typescript [-c sesamy.yaml ...]
```

## What it does

Reads `googleTag.typeScript.packages` from your config, runs [`gocontemplate`](https://github.com/foomo/gocontemplate) over the listed Go types, and writes `.ts` files to `googleTag.typeScript.outputPath`. Existing files in the output directory are overwritten.

The point: your frontend, your Go services, and the tags you push to GTM all agree on the shape of every event.

## Configuration

```yaml
googleTag:
  typeScript:
    outputPath: ./web/src/sesamy
    directory: .                         # path to the go.mod
    packages:
      - path: github.com/foomo/sesamy-go/pkg/event
        types:
          - PageView
          - AddToCart
          - Purchase
```

| Field | Purpose |
|-------|---------|
| `outputPath` | Directory for generated `.ts` files (relative paths resolved against CWD). |
| `directory` | Path to a `go.mod` whose dependency graph includes the listed packages. |
| `packages[].path` | Go import path of the event package. |
| `packages[].types` | Exported Go type names to generate. |

## Flags

| Flag | Default | Purpose |
|------|---------|---------|
| `-c, --config` | `sesamy.yaml` | Config file(s). |

## Tips

- Commit the generated `.ts` files. They're an artifact your frontend depends on.
- Re-run `sesamy typescript` in CI and fail if the working tree is dirty — that catches missed regenerations.
- Custom event types from your own Go packages work just like `sesamy-go` types. See [sesamy-go integration](/guide/sesamy-go).
