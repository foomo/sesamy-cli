[![Build Status](https://github.com/foomo/sesamy-cli/actions/workflows/test.yml/badge.svg?branch=main&event=push)](https://github.com/foomo/sesamy-cli/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/foomo/sesamy-cli)](https://goreportcard.com/report/github.com/foomo/sesamy-cli)
[![GoDoc](https://godoc.org/github.com/foomo/sesamy-cli?status.svg)](https://godoc.org/github.com/foomo/sesamy-cli)
[![goreleaser](https://github.com/foomo/sesamy-cli/actions/workflows/release.yml/badge.svg)](https://github.com/foomo/sesamy-cli/actions)

<p align="center">
  <img alt="sesamy" src=".github/assets/sesamy.png"/>
</p>

# Sesamy CLI

> CLI to keep you sane while working with GTM.

## Installing

Install the latest release of the cli:

````bash
$ brew update
$ brew install foomo/tap/sesamy-cli
````

## Usage

```shell
$ sesamy help
Server Side Tag Management System

Usage:
  sesamy [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  config      Print config
  help        Help about any command
  list        List Google Tag Manager containers
  provision   Provision Google Tag Manager containers
  tags        Print out all available tags
  typescript  Generate typescript events
  version     Print version

Flags:
  -c, --config string   config file (default is sesamy.yaml) (default "sesamy.yaml")
  -h, --help            help for sesamy
  -v, --verbose         output debug information

Use "sesamy [command] --help" for more information about a command.
```

## Configuration

Add a `sesamy.yaml` configuration

```yaml
# yaml-language-server: $schema=https://raw.githubusercontent.com/foomo/sesamy-cli/refs/heads/main/sesamy.schema.json
version: '1.0'

# Whether to redact the visitor ip
redactVisitorIp: true
# Enable region specific settings
# https://developers.google.com/tag-platform/tag-manager/server-side/enable-region-specific-settings
enableGeoResolution: true

# --- Google API settings
googleApi:
  # Single line Service Account credentials
  credentials: '{...\\n...\\n...}'
  # Path to the Service Account credentials json file
  credentialsFile: google_service_account_creds.json
  # Current API request quota (send a request to increase the quota)
  requestQuota: 15

# --- Google Tag Manager settings
googleTagManager:
  # The account id
  accountId: '6099238525'
  # Web container settings
  webContainer:
    # The container tag id
    tagId: GTM-57BHX34G
    # The container id
    containerId: '175355532'
    # The workspace id that should be used by the api
    workspaceId: '23'
  # Server container settings
  serverContainer:
    # The container tag id
    tagId: GTM-5NWPR4QW
    # The container id
    containerId: '175348980'
    # The workspace id that should be used by the api
    workspaceId: '10'
  # Web container variables
  webContainerVariables:
    dataLayer:
      - link_url
    lookupTables:
      link_url_conversion_label:
        input: '{{dlv.link_url}}'
        valueTable:
          123456: 'https://foomo.org/'
  # Server container variables
  serverContainerVariables:
    eventData:
      - link_url
    lookupTables:
      link_url_conversion_label:
        input: '{{event.link_url}}'
        valueTable:
          123456: 'https://foomo.org/'

# --- Google Tag settings
googleTag:
  # A tag ID is an identifier that you put on your page to load a given Google tag
  tagId: G-PZ5ELRCR31
  # Whether a page_view should be sent on initial load
  sendPageView: true
  # Optional custom server container url
  serverContainerUrl: ''
  # TypeScript settings
  typeScript:
    # Target directory for generate files
    outputPath: path/to/target
    # Path to the go.mod file
    directory: .
    # Contemplate package config for generated events
    packages:
      - path: github.com/foomo/sesamy-go/pkg/event
        types:
          ## GA4 Automatically collected events
          ## https://support.google.com/analytics/answer/9234069
          - Click
          - FileDownload
          - FirstVisit
          - FormStart
          - FormSubmit
          - PageView
          - Scroll
          - UserEngagement
          - VideoComplete
          - VideoProgress
          - VideoStart
          - ViewSearchResults
          ## Recommended events
          ## https://developers.google.com/tag-platform/gtagjs/reference/events
          - AdImpression
          - AddPaymentInfo
          - AddShippingInfo
          - AddToCart
          - AddToWishlist
          - BeginCheckout
          - CampaignDetails
          - CloseConvertLead
          - CloseUnconvertLead
          - DisqualifyLead
          - EarnVirtualMoney
          - Exception
          - GenerateLead
          - JoinGroup
          - LevelEnd
          - LevelStart
          - LevelUp
          - Login
          - PostScore
          - Purchase
          - QualifyLead
          - Refund
          - RemoveFromCart
          - ScreenView
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
          - ViewCart
          - ViewItem
          - ViewItemList
          - ViewPromotion
          - WorkingLead

# --- Google Analytics settings
googleAnalytics:
  # Enable provider
  enabled: true
  # Google Consent settings
  googleConsent:
    # Enable consent mode
    enabled: true
    # Consent mode name
    mode: analytics_storage
  # Google GTag.js override settings
  googleGTagJSOverride:
    # Provision custom client
    enabled: true
    # Client priority
    priority: 10
    # Patch ecommerce items
    ecommerceItems: true
  # Google Tag Manager web container settings
  webContainer:
    # Path to the go.mod file
    directory: .
    # Contemplate package config for generated events
    packages:
      - path: github.com/foomo/sesamy-go/pkg/event
        types:
          - PageView
          - SelectItem
  # Google Tag Manager server container settings
  serverContainer:
    # Path to the go.mod file
    directory: .
    # Contemplate package config for generated events
    packages:
      - path: github.com/foomo/sesamy-go/pkg/event
        types:
          - PageView
          - SelectItem

# --- Google Ads
googleAds:
  # Enable provider
  enabled: true
  # Google Ads Conversion Tracking ID
  conversionId: ''
  # Google Consent settings
  googleConsent:
    # Enable consent mode
    enabled: true
    # Consent mode name
    mode: ad_storage
  # Google Ads Remarketing settings
  remarketing:
    # Enable Google Ads Remarketing
    enabled: true
    # Enable conversion linking
    enableConversionLinker: true
  # Google Ads Conversion settings
  conversion:
    # Enable Google Ads Conversion
    enabled: true
    # Google Ads Conversion Tracking Label
    conversionLabel: ''
    # Google Tag Manager server container settings
    serverContainer:
      # Path to the go.mod file
      directory: .
      # Contemplate package config for generated events
      packages:
        - path: github.com/foomo/sesamy-go/pkg/event
          types:
            - AddToCart
            - Purchase

# --- Conversion Linker settings
conversionLinker:
  # Enable provider
  enabled: true
  # Google Consent settings
  googleConsent:
    # Enable consent mode
    enabled: true
    # Consent mode name
    mode: ad_storage

# --- Umami settings
umami:
  # Enable provider
  enabled: true
  # Enter an optional fixed domain to override event data
  domain: your-domain.com
  # Paste ID for your website from the Umami settings
  websiteId: ''
  # Endpoint url of the umami api
  endpointUrl: https://umami.your-domain.com
  # Google Consent settings
  googleConsent:
    # Enable consent mode
    enabled: true
    # Consent mode name
    mode: analytics_storage
  # Google Tag Manager server container settings
  serverContainer:
    # Path to the go.mod file
    directory: .
    # Contemplate package config for generated events
    packages:
      - path: github.com/foomo/sesamy-go/pkg/event
        types:
          - PageView
          - SelectItem
# --- Criteo
criteo:
  # Enable provider
  enabled: true
  # Criteo caller id
  callerId: 123
  # Criteo partner id
  partnerId: 123456
  # Criteo applicaiton id
  applicationId: com.foomo
  # Google Tag Manager server container settings
  serverContainer:
    # Path to the go.mod file
    directory: .
    # Contemplate package config for generated events
    packages:
      - path: github.com/foomo/sesamy-go/pkg/event
        types:
          - AddToCart
          - BeginCheckout
          - PageView
          - Purchase
          - ViewItem
          - ViewItemList
          - ViewCart

# --- Facebook
# https://developers.facebook.com/docs/marketing-api/conversions-api/guides/gtm-server-side
facebook:
  # Enable provider
  enabled: true
  # Facebook pixel id
  pixelId: ''
  # To use the Conversions API, you need an access token.
  apiAccessToken: ''
  # Code used to verify that your server events are received correctly by Conversions API
  testEventToken: ''
  # Google Tag Manager server container settings
  # Google Consent settings
  googleConsent:
    # Enable consent mode
    enabled: true
    # Consent mode name
    mode: ad_storage
  serverContainer:
    # Path to the go.mod file
    directory: .
    # Contemplate package config for generated events
    packages:
      - path: github.com/foomo/sesamy-go/pkg/event
        types:
          - AddPaymentInfo
          - AddToCart
          - AddToWishlist
          - PageView
          - Purchase
          - Search
          - BeginCheckout
          - GenerateLead
          - ViewItem

# --- Emarsys
emarsys:
  # Enable provider
  enabled: true
  # Emarsys merchant id
  merchantId: ''
  # Google Consent settings
  googleConsent:
    # Enable consent mode
    enabled: true
    # Consent mode name
    mode: analytics_storage
  # Google Tag Manager server container settings
  serverContainer:
    # Path to the go.mod file
    directory: .
    # Contemplate package config for generated events
    packages:
      - path: github.com/foomo/sesamy-go/pkg/event
        types:
          - Purchase
          - ViewItem
          - ViewItemList

# --- Tracify
tracify:
  # Enable provider
  enabled: true
  # Tracify token
  token: ''
  # Tracify customer site id
  customerSiteId: ''
  # Google Consent settings
  googleConsent:
    # Enable consent mode
    enabled: true
    # Consent mode name
    mode: analytics_storage
  # Google Tag Manager server container settings
  serverContainer:
    # Path to the go.mod file
    directory: .
    # Contemplate package config for generated events
    packages:
      - path: github.com/foomo/sesamy-go/pkg/event
        types:
          - AddToCart
          - PageView
          - ViewItem
          - Purchase

# --- Hotjar
hotjar:
  # Enable provider
  enabled: true
  # Hotjar site id
  siteId: 123456

# --- Cookiebot CMP
cookiebot:
  # Enable provider
  enabled: true
  # Name of the manually installed Cookiebot CMP tag template
  # "https://tagmanager.google.com/gallery/#/owners/cybotcorp/templates/gtm-templates-cookiebot-cmp
  templateName: Cookiebot CMP
  # Cookiebot id
  cookiebotId: ''
  # CDN Region (eu, com)
  cdnRegion: eu
  # Enable URL passthrough
  urlPassthrough: false
  # Enable advertiser consent mode
  advertiserConsentModeEnabled: false
```

## Caveats

You might need to increase your Google Tag Manager API quotas, since they are limited to 15 r/m by default.

## How to Contribute

Please refer to the [CONTRIBUTING](.gihub/CONTRIBUTING.md) details and follow the [CODE_OF_CONDUCT](.gihub/CODE_OF_CONDUCT.md) and [SECURITY](.github/SECURITY.md) guidelines.

## License

Distributed under MIT License, please see license file within the code for more details.

_Made with ♥ [foomo](https://www.foomo.org) by [bestbytes](https://www.bestbytes.com)_
