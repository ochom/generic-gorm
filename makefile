# Makefile

SHELL := /bin/bash

tidy:
	go mod tidy

lint:
	@echo "Linting ..."
	@golangci-lint run --timeout 5m

test:
	@echo "Running tests ..."
	@go test ./...
 