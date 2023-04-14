# force bash to use build-ins
SHELL:=bash

# go options
GO          ?= go

help: ## Display this help
	@ echo "Please use \`make <target>' where <target> is one of:"
	@ echo
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@ echo

.DEFAULT_GOAL := help

test-unit: ## Execute unit tests
	$(GO) test -race -v -p=1 ./...

build:
	$(GO) build  -o ./myhttp main.go

.PHONY: build help test-unit
