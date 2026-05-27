# tags

Print every provider the CLI knows about and the GTM tag identifier it uses.

```bash
sesamy tags
```

## What it does

Prints a table of `Name → Tag`. The `Tag` column is the GTM template identifier that `sesamy provision` looks up (or creates) when generating tags for that provider. Useful when:

- Onboarding to confirm a provider is supported.
- Debugging — match a tag in the GTM UI back to the provider that owns it.

## Example output

```
NAME              | TAG
------------------|----------------------------------
conversionlinker  | gtag-conversion-linker
cookiebot         | cookiebot-cmp
criteo            | criteo-onetag
emarsys           | emarsys
facebook          | facebook-capi
googleads         | google-ads
googleanalytics   | google-analytics
googleconsent     | google-consent
googletag         | google-tag
googletagmanager  | google-tag-manager
hotjar            | hotjar
microsoftads      | microsoft-ads
mixpanel          | mixpanel
pinterest         | pinterest
tracify           | tracify
umami             | umami
```

The actual identifiers come from each provider's `constants.go` under `pkg/provider/*/`.
