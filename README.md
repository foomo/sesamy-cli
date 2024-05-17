# Sesamy CLI

[![Build Status](https://github.com/foomo/sesamy-cli/actions/workflows/test.yml/badge.svg?branch=main&event=push)](https://github.com/foomo/sesamy-cli/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/foomo/sesamy-cli)](https://goreportcard.com/report/github.com/foomo/sesamy-cli)
[![godoc](https://godoc.org/github.com/foomo/sesamy-cli?status.svg)](https://godoc.org/github.com/foomo/sesamy-cli)
[![goreleaser](https://github.com/foomo/sesamy-cli/actions/workflows/release.yml/badge.svg)](https://github.com/foomo/sesamy-cli/actions)

> CLI to keep you sane while working with GTM.

## Installing

Install the latest release of the cli:

````bash
$ brew update
$ brew install foomo/tap/sesamy-cli
````

## Usage

Add a `sesamy.yaml` configurtion

```yaml
google:
  gt:
    send_page_views: true
    server_container_url: https://sst.your.domain.com

  ga4:
    measurement_id: G-PZ5ELRCR31

  gtm:
    account_id: 6099238525
    web:
      container_id: 175355532
      measurement_id: GTM-57BHX34G
      workspace_id: 23
    server:
      container_id: 175348980
      measurement_id: GTM-5NWPR4QW
      workspace_id: 10

  request_quota: 15
  credentials_file: ./google_service_account_creds.json


typescript:
  packages:
    - path: 'github.com/username/repository/event'
      events:
        - Custom
    - path: 'github.com/foomo/sesamy-go/event'
      events:
        - PageView
        - SelectItem
  output_path: '/path/to/index.ts'

tagmanager:
  packages:
    - path: 'github.com/username/repository/event'
      events:
        - Custom
    - path: 'github.com/foomo/sesamy-go/event'
      events:
        - AddPaymentInfo
        - AddShippingInfo
        - AddToCart
        - AddToWishlist
        - AdImpression
        - BeginCheckout
        - CampaignDetails
        - Click
        - EarnVirtualMoney
        - FileDownload
        - FormStart
        - FormSubmit
        - GenerateLead
        - JoinGroup
        - LevelEnd
        - LevelStart
        - LevelUp
        - Login
        - PageView
        - PostScore
        - Purchase
        - Refund
        - RemoveFromCart
        - ScreenView
        - Scroll
        - Search
        - SelectContent
        - SelectItem
        - SelectPromotion
        - SessionStart
        - Share
        - SignUp
        - SpendVirtualCurrency
        - TutorialBegin
        - TutorialComplete
        - UnlockAchievement
        - UserEngagement
        - VideoComplete
        - VideoProgress
        - VideoStart
        - ViewCart
        - ViewItem
        - ViewItemList
        - ViewPromotion
        - ViewSearchResults
  prefixes:
    client: ''
    folder: ''
    tags:
      ga4_event: 'GA4 - '
      google_tag: ''
      server_ga4_event: 'GA4 - '
    triggers:
      client: ''
      custom_event: 'Event - '
    variables:
      constant: ''
      event_model: 'dlv.eventModel.'
      gt_event_settings: 'Event Settings - '
      gt_settings: 'Settings - '
```

## Caveats

You might need to increase your Google Tag Manager API quotas, since they are limited to 15 r/m by default.

## How to Contribute

Make a pull request...

## License

Distributed under MIT License, please see license file within the code for more details.
