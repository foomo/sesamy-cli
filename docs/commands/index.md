# Commands

`sesamy-cli` is built with [Cobra](https://cobra.dev). Every command supports `--help` and shell completion.

## Global flags

| Flag | Shorthand | Default | Purpose |
|------|-----------|---------|---------|
| `--config` | `-c` | `sesamy.yaml` | Path to a config file. Repeat to layer multiple files. |
| `--verbose` | `-v` | `false` | Print debug output. |
| `--help` | `-h` | — | Show help for any command. |

## Available commands

| Command | What it does |
|---------|--------------|
| [`provision web`](./provision) / `provision server` | Apply local config to a GTM container. |
| [`list web`](./list) / `list server` | List entities currently in a container. |
| [`diff web`](./diff) / `diff server` | Show planned changes without applying them. |
| [`typescript`](./typescript) | Generate TypeScript event types from `sesamy-go`. |
| [`tags`](./tags) | Print every provider tag the CLI knows about. |
| [`open`](./open) | Open the GTM or GA console for the configured account. |
| [`config`](./config) | Print the merged effective config. |
| `version` | Print the binary version. |
| `completion <shell>` | Emit shell completion script. |

## The web / server pair

Most commands have `web` and `server` subcommands, one per GTM container type. Provider behavior differs between them — e.g. Facebook's Conversions API only makes sense in the server container — so the CLI keeps them distinct.

```bash
sesamy provision web      # web container
sesamy provision server   # server container
```

Run both whenever you change the config.
