# Sesamy CLI

[![Go Report Card](https://goreportcard.com/badge/github.com/foomo/sesamy-cli)](https://goreportcard.com/report/github.com/foomo/sesamy-cli)
[![godoc](https://godoc.org/github.com/foomo/sesamy-cli?status.svg)](https://godoc.org/github.com/foomo/sesamy-cli)
[![goreleaser](https://github.com/foomo/sesamy-cli/workflows/goreleaser/badge.svg)](https://github.com/foomo/sesamy-cli/actions)

> CLI to keep you sane while working with GTM.

## Quickstart

Add a `sesamy.yaml` configurtion

```yaml
google:
  ga4:
    measurement_id: G-PZ5ELRCR31

  gtm:
    account_id: 6099238525
    server:
      container_id: 175348980
      workspace_id: 10
      measurement_id: GTM-5NWPR4QW

    web:
      container_id: 175355532
      workspace_id: 23
      measurement_id: GTM-57BHX34G

  credentials_file: ./tmp/google_service_account_creds.json

events:
  packages:
    - path: "github.com/foomo/sesamy-cli/_example/server"
      output_path: "./_example/client/types.d.ts"
      indent: "\t"
```

## Caveats

You might need to increase your Google Tag Manager API quotas, since they are limited to 15 r/m by default.

## How to Contribute

Make a pull request...

## License

Distributed under MIT License, please see license file within the code for more details.
