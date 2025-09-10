.DEFAULT_GOAL	:= help
-include .makerc
PATH:=bin:$(PATH)

# --- Targets -----------------------------------------------------------------

# This allows us to accept extra arguments
%: .husky
	@:

.PHONY: .husky
# Configure git hooks for husky
.husky:
	@if ! command -v husky &> /dev/null; then \
		echo "ERROR: missing executeable 'husky', please run:"; \
		echo "\n$ make brew\n"; \
	fi
	@git config core.hooksPath .husky

### Tasks

.PHONY: doc
## Open go docs
doc:
	@open "http://localhost:6060/pkg/github.com/foomo/sesamy-cli/"
	@godoc -http=localhost:6060 -play

.PHONY: test
## Run tests
test:
	@GO_TEST_TAGS=-skip go test -tags=safe -coverprofile=coverage.out -race -json -v ./... 2>&1 | tee /tmp/gotest.log | gotestfmt

.PHONY: lint
## Run linter
lint:
	@golangci-lint run

.PHONY: lint.fix
## Fix lint violations
lint.fix:
	@golangci-lint run --fix

.PHONY: tidy
## Run go mod tidy
tidy:
	@go mod tidy

.PHONY: outdated
## Show outdated direct dependencies
outdated:
	@go list -u -m -json all | go-mod-outdated -update -direct

.PHONY: build
## Build binary
build:
	@mkdir -p bin
	@echo "building: bin/sesamy"
	@go build -tags=safe -o bin/sesamy main.go

.PHONY: install
## Install binary
install:
	@echo "installing: ${GOPATH}/bin/sesamy"
	@go build -tags=safe -o ${GOPATH}/bin/sesamy main.go

.PHONY: install.debug
## Install debug binary
install.debug:
	@go build -tags=safe -gclags="all=-N -l" -o ${GOPATH}/bin/sesamy main.go

### Utils

.PHONY: brew
## Install project binaries
brew:
	@ownbrew install

.PHONY: help
## Show help text
help:
	@echo "\033[1;36mSesamy CLI\033[0m"
	@awk '{ \
		if($$0 ~ /^### /){ \
			if(help) printf "\033[36m%-23s\033[0m %s\n\n", cmd, help; help=""; \
			printf "\n\033[1;36m%s\033[0m\n", substr($$0,5); \
		} else if($$0 ~ /^[a-zA-Z0-9._-]+:/){ \
			cmd = substr($$0, 1, index($$0, ":")-1); \
			if(help) printf "  \033[36m%-23s\033[0m %s\n", cmd, help; help=""; \
		} else if($$0 ~ /^##/){ \
			help = help ? help "\n                        " substr($$0,3) : substr($$0,3); \
		} else if(help){ \
			print "\n                        " help "\n"; help=""; \
		} \
	}' $(MAKEFILE_LIST)
