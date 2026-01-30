.PHONY: help build test test-coverage lint fmt clean install run-examples

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-20s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install: ## Install dependencies
	go mod download
	go mod tidy

build: ## Build the project
	go build ./...

test: ## Run tests
	go test ./...

test-coverage: ## Run tests with coverage
	go test -cover ./...
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

lint: ## Run linter
	@which golangci-lint > /dev/null || (echo "golangci-lint not installed. Install from https://golangci-lint.run/usage/install/" && exit 1)
	golangci-lint run

fmt: ## Format code
	go fmt ./...
	gofmt -s -w .

clean: ## Clean build artifacts
	rm -f coverage.out coverage.html
	go clean

run-basic: ## Run basic example
	cd examples/basic && go run main.go

run-top100: ## Run top100 example
	cd examples/top100 && go run main.go

run-error-handling: ## Run error handling example
	cd examples/error_handling && go run main.go

run-global-metrics: ## Run global metrics example
	cd examples/global_metrics && go run main.go

run-with-env: ## Run with env example
	cd examples/with_env && go run main.go
