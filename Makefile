APP_VERSION ?= 0.1-dev
BINARY_NAME ?= binary-name
CONTAINER_NAME ?= binary-name
GITHUB_REPO ?= github.com/organization/repo-name

BUILD_FLAGS=-X=main.CompileDate=$(COMPILE_DATE) -X=main.GitCommit=$(GIT_COMMIT) -X=main.Version=$(APP_VERSION)
GIT_COMMIT=$(shell git rev-parse HEAD)
COMPILE_DATE=$(shell date -u +%Y%m%d.%H%M%S)
UNAME=$(shell uname -s | tr '[:upper:]' '[:lower:]')

all: build

deps: # Install all dependencies.
	dep ensure

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

clean: ## Remove compiled binaries.
	rm -f bin/$(BINARY_NAME) || true
	rm -f bin/$(BINARY_NAME)*gz || true

build: clean ## Build.
	go build -ldflags "$(BUILD_FLAGS)" -o bin/$(BINARY_NAME) main.go

rebuild: clean ## Force rebuild of all packages.
	go build -a -ldflags "$(BUILD_FLAGS)" -o bin/$(BINARY_NAME) main.go

linux: clean ## Cross compile for linux.
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$(BUILD_FLAGS)" -o bin/$(BINARY_NAME) main.go

gzip: ## Compress current compiled binary.
	gzip bin/$(BINARY_NAME)
	mv bin/$(BINARY_NAME).gz bin/$(BINARY_NAME)-$(APP_VERSION)-$(UNAME)-amd64.gz

release: clean build gzip ## Full release process.

unit: ## Run unit tests.
	ci/tests.sh

test: unit ## Run all tests.

.PHONY: help all deps clean build gzip release unit test
