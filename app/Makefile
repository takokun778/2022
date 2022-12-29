include .env
export

.PHONY: name
name: ## display app name
	@echo ${APP_NAME}

.PHONY: compile
compile: ## go compile
	@go build -v ./... && go clean

.PHONY: fmt
fmt: ## go format
	@go fmt ./...

.PHONY: lint
lint: ## go lint
	@golangci-lint run --fix

.PHONY: tidy
tidy: ## go mod tidy
	@go mod tidy

.PHONY: run
run:
	@go run cmd/main.go

.PHONY: mod
mod: ## go modules list
	@go list -u -m all

.PHONY: update
update: ## go modules update
	@go get -u -t ./...

.PHONY: test
test: ## go test
	@$(call _test,${c})

define _test
if [ -z "$1" ]; then \
	go test ./... ; \
else \
	go test ./... -count=1 ; \
fi
endef

.PHONY: help
help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
