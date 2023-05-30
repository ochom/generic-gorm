SHELL=/bin/bash

pull:
	@echo "Pulling latest changes..."
	@git pull

pull-dev:
	@echo "Pulling latest changes..."
	@git pull origin dev

tidy:
	@echo "Tidying up..."
	@go mod tidy

test:
	@echo "Running tests..."
	@go test ./...

lint:
	@echo "Linting ..."
	@golangci-lint run --timeout 5m
