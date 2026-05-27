# Installation

`sesamy-cli` is distributed as a single Go binary.

## Homebrew (recommended)

```bash
brew update
brew install foomo/tap/sesamy-cli
```

## Go install

```bash
go install github.com/foomo/sesamy-cli@latest
```

Builds from source use the `-tags=safe` build tag. The `Makefile` in the repo wires this up; if you build manually, pass `-tags=safe`:

```bash
go build -tags=safe -o sesamy .
```

## From release archive

Pre-built binaries for macOS, Linux, and Windows are attached to each [GitHub release](https://github.com/foomo/sesamy-cli/releases). Download the archive for your platform, extract, and place `sesamy` somewhere on your `$PATH`.

## Verify the install

```bash
sesamy version
sesamy help
```

You should see the available commands listed:

```
Available Commands:
  completion  Generate the autocompletion script for the specified shell
  config      Print config
  diff        Print Google Tag Manager container status diff
  help        Help about any command
  list        List Google Tag Manager containers
  open        Open links in the browser
  provision   Provision Google Tag Manager containers
  tags        Print out all available tags
  typescript  Generate typescript events
  version     Print version
```

## Shell completion

`sesamy` ships standard Cobra completion:

```bash
# zsh
sesamy completion zsh > "${fpath[1]}/_sesamy"

# bash
sesamy completion bash > /etc/bash_completion.d/sesamy
```

## Next

Head to [Quick start](./quick-start) to wire up your first config.
