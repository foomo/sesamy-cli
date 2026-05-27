---
layout: home

hero:
  name: sesamy-cli
  text: GTM, version-controlled.
  tagline: Provision, diff, and inspect Google Tag Manager server-side and web containers from a single YAML file.
  image:
    src: /logo.png
    alt: sesamy-cli
  actions:
    - theme: brand
      text: Get started
      link: /guide/introduction
    - theme: alt
      text: View commands
      link: /commands/
    - theme: alt
      text: GitHub
      link: https://github.com/foomo/sesamy-cli

features:
  - title: Declarative GTM
    details: One sesamy.yaml describes both your web and server containers. The CLI reconciles GTM state to match.
  - title: Multi-provider out of the box
    details: Google Analytics, Google Ads, Facebook, Pinterest, Microsoft Ads, Criteo, Emarsys, Mixpanel, Umami, Tracify, Hotjar, Cookiebot — wired with sensible defaults.
  - title: Diff before you ship
    details: See exactly which tags, triggers, and variables will change before they hit your container.
  - title: Typed events
    details: Generate TypeScript event definitions from your sesamy-go event models so frontend and GTM stay in lockstep.
---
