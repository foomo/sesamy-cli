# Workflow

A typical day with `sesamy-cli` looks like this.

## Local development

```bash
# 1. Edit sesamy.yaml — add or tweak a provider.
$EDITOR sesamy.yaml

# 2. Preview the change.
sesamy diff web
sesamy diff server

# 3. Apply when the diff looks right.
sesamy provision web
sesamy provision server

# 4. Verify.
sesamy list web
```

## In code review

Commit `sesamy.yaml` alongside the application change. Reviewers can run `sesamy diff` locally — or, better, have CI do it on every PR.

## In CI

A minimal CI job:

```yaml
- run: sesamy diff web -c sesamy.yaml -c sesamy.prod.yaml
- run: sesamy diff server -c sesamy.yaml -c sesamy.prod.yaml
```

For a deploy job, run `sesamy provision` against the appropriate workspace. The CLI is idempotent: re-running it without config changes produces no GTM mutations.

::: tip
Run `provision` against a **non-default workspace** in CI and require a human to publish in the GTM UI. That keeps a human in the loop for the actual go-live step.
:::

## When the config changes

Whenever the YAML changes:

1. `sesamy diff` to confirm the change is what you expect.
2. `sesamy provision` to apply.
3. If you generate TypeScript events, run `sesamy typescript` and commit the regenerated files.

## When code in this repo changes

If you're contributing to `sesamy-cli` itself: a change to commands, the config schema, or a provider must be reflected in these docs. Every section under [Commands](/commands/) and [Providers](/providers/) maps directly to a folder under `pkg/` or `cmd/`.
