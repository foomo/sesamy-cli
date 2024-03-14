# Sesamy CLI

[![Build Status](https://github.com/foomo/sesamy-cli/actions/workflows/test.yml/badge.svg?branch=main&event=push)](https://github.com/foomo/sesamy-cli/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/foomo/sesamy-cli)](https://goreportcard.com/report/github.com/foomo/sesamy-cli)
[![godoc](https://godoc.org/github.com/foomo/sesamy-cli?status.svg)](https://godoc.org/github.com/foomo/sesamy-cli)
[![goreleaser](https://github.com/foomo/sesamy-cli/actions/workflows/release.yml/badge.svg)](https://github.com/foomo/sesamy-cli/actions)

> CLI to keep you sane while working with GTM.

## Installing

Install the latest release of the cli:

```bash
$ brew update
$ brew install foomo/tap/sesamy-cli
```

## Caveats

You might need to increase your Google Tag Manager API quotas, since they are limited to 15 r/m by default.

Currently the sesamy-cli doesn't submit the changes automatically. This needs to be done manually in Google Tag Manager

## How to Contribute

Make a pull request...

## License

Distributed under MIT License, please see license file within the code for more details.

## Glossary

As some terms for Google Analytics and GTM events and ressources can be confusing we want to clarify the following definitions that are used in this README:

- Google Tag Manager: A configuration interface for different kind of containers (Web, Server etc.). The server container configuration is represented by a Tagging server
- Web container: This is a tagging container hosted by Google and can be configured via the Google Tag Manager
- Tagging server: This is a server that you setup yourself based on an docker image from Google that can receive and process "/g/collect" and other requests. You have to set environment variables to connect your server with the Google Tag Manager configuration interface
- Preview server: Similar to the tagging server, with the same docker image, but different environment variables to be needed. This server provides you with a Web UI to preview incoming and outgoing requests of the Tagging server

## What problem does it solve?

### Automate the creating of web and server container ressources

If tracking events need params that are not standard or e-commerce (enhanced) in Google Analytics then you have to configure those events and params on the Web container GA4 Tag manually. Same is true if it is a standard parameter, but not part of the recommended parameters of the event type.

In order to associate a parameter with an event, the parameter needs to be configured as an additional ressource in the web container by creating a new datalayer variable. The data layer variables name needs to start with "eventModel." e.g. "eventModel.user_group".

Sesamy CLI automates this process

### Guarantee typesafety when tracking events with gtag()

Auto generate typescript types based on the event structs that you define with sesamy-go (https://github.com/foomo/sesamy-go/). When using gtag() the generated types make sure that you get code completion and type checks on gtag() calls

```typescript
//common.tsx

import { EventNames, EventParams } from "events";

export const gtag = (name: EventNames, params: EventParams) => {
  window.dataLayer = window.dataLayer || [];
  window.gtag =
    window.gtag ||
    function () {
      // @ts-ignore
      // eslint-disable-next-line prefer-rest-params
      window.dataLayer?.push(arguments);
    };
  window.gtag("event", name, params);
};
```

```typescript
//e.g. checkout.tsx

gtag("begin_checkout", {
  currency: props.summary.total.currency,
  value: props.summary.total.cents / 100,
  items: mapCartItems(props.items),
} as BeginCheckout);
```

## Configuration

You need a `sesamy.yaml` configuration. Check out the tygo documentation (https://github.com/gzuidhof/tygo) for additional, optional configurations.

Explanation:

- account_id: Open Server container in tagmanager -> url parameter `accountId`
- container_id: Open Server/web container in tagmanager -> url param: `container`
- workspace_id: Open Server/web container in tagmanager -> url param: `workspace`
- ga4.measurement_id: The Google Analytics measurment ID
- gtm.[server/web].measurement_id: The identifier of the web or server container
- credentials_file: You need a credentials.json from a Google Cloud Platform service account in order to connect to the web and server container. You need to register that user with edit rights on the web and the server container
- server_container_url: URL of your server container (has to be https)

<br>

```yml
# sesamy.yaml

# NOTE: needs to be placed next to the go.mod file

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
  server_container_url: http://sst.my-domain.com/

# Generates typescript types
typescript:
  packages:
    - path: "github.com/foomo/sesamy-cli/_example/server"
      output_path: "./_example/client/types.d.ts"
      indent: "\t"

# Provisions backend structs. Exclude allnon event structs
tagmanager:
  packages:
    - path: "github.com/foomo/sesamy-cli/_example/server"
      output_path: "./_example/client/types.d.ts"
      exclude_files:
        - item.go
```

<br>

Currently to be created manually: typescript.go file that exports all the event names and the event parameters

```go
// typescript.go

package event

//tygo:emit
var _ = `export type EventNames =
	| 'add_payment_info'
	| 'add_shipping_ingo'
	| 'add_to_cart'
	| 'add_to_wishlist'
	| 'ce_account_signup'
	| 'ce_login'
	| 'purchase'
	| 'refund'
	| 'remove_from_cart'
	| 'select_item'
	| 'view_cart'
	| (string & NonNullable<unknown>);

export type EventParams =
	| AddPaymentInfo
	| AddShippingInfo
	| AddToCart
	| AddToWishlist
	| CEAccountSignup
	| CELogin
	| Purchase
	| Refund
	| RemoveFromCart
	| SelectItem
	| ViewCart;
```

## How to run

The following two command setup Google Tag Manager ressources in the server and in the web container

```bash
sesamy tagmanager server
```

Creates following ressources on the server container

- Folder Sesamy
- Variable for MeasurementID
- "Google Tag Manager Web Container" client
- "GA4" client
- "GA4" trigger
- "GA4" tag -> "GA4 trigger" is used here

<br>

```bash
sesamy tagmanager client
```

Creates following ressources on the web container

- Folder Sesamy
- Variable for MeasurementID
- Event params for GA4 Tag, Tag, Trigger and Variables
