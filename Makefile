SHELL := /bin/bash
include .env
export
export APP_NAME := $(basename $(notdir $(shell pwd)))

.PHONY: help
help: ## display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: aqua
aqua: ## install aqua ref. https://aquaproj.github.io/
	@brew install aquaproj/aqua/aqua

.PHONY: tool
tool: ## install tool
	@aqua install

.PHONY: build
build: ## go build
	@go build -v ./... && go clean

.PHONY: fmt
fmt: ## go format
	@go fmt ./...

.PHONY: lint
lint: ## go lint ref. https://golangci-lint.run/
	@golangci-lint run ./... --fix

.PHONY: mod
mod: ## go mod tidy & go mod vendor
	@go mod tidy
	@go mod vendor

.PHONY: update
update: ## go modules update
	@go get -u -t ./...
	@go mod tidy
	@go mod vendor

.PHONY: gen
gen: ## Generate code.
	@go generate ./...
	@go mod tidy
	@go mod vendor

.PHONY: load
load: ## api server reload
	@touch cmd/api/main.go

.PHONY: test
test: ## unit test
	@$(call _test,${c})

define _test
if [ -z "$1" ]; then \
	go test ./internal/... ; \
else \
	go test ./internal/... -count=1 ; \
fi
endef

.PHONY: integration
integration: ## run integration test. If you want to invalidate the cache, please specify an argument like `make integration c=c`.
	@$(call _integration,${c})

define _integration
if [ -z "$1" ]; then \
	go test ./test/integration/... ; \
else \
	go test ./test/integration/... -count=1 ; \
fi
endef

.PHONY: e2e
e2e: ## run e2e test. If you want to invalidate the cache, please specify an argument like `make e2e c=c`.
	@$(call _e2e,${c})

define _e2e
if [ -z "$1" ]; then \
	go test ./test/e2e/... ; \
else \
	go test ./test/e2e/... -count=1 ; \
fi
endef

.PHONY: up
up: ## docker compose up with air hot reload
	@docker compose --project-name ${APP_NAME} --file ./.docker/compose.yaml up -d

.PHONY: down
down: ## docker compose down
	@docker compose --project-name ${APP_NAME} down --volumes

.PHONY: balus
balus: ## destroy everything about docker. (containers, images, volumes, networks.)
	@docker compose --project-name ${APP_NAME} down --rmi all --volumes

.PHONY: ymlfmt
ymlfmt: ## format yaml file
	@yamlfmt

.PHONY: ymlint
ymlint: ## lint yaml file
	@yamlfmt -lint

.PHONY: psql
psql:
	@docker exec -it ${APP_NAME}-postgres psql -U postgres

.PHONY: migrate
migrate: ## migrate prisma schema
	@(cd schema && bun run prisma db push)

.PHONY: prismastudio
prismastudio: ## execute prisma studio
	@(cd schema && bun run prisma studio)

.PHONY: prismapull
prismapull: ## import prisma schema
	@(cd schema && bun run prisma db pull)

.PHONY: schemafmt
schemafmt: 
	@(cd schema && bun run prisma format)
