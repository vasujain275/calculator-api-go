# Makefile for calculator-api-go

APP_NAME := calculator-api-go
VERSION  ?= 1.0.0
BUILD_DIR := build
GO       := go

LDFLAGS := -s -w -X main.version=$(VERSION)

.PHONY: all
all: dev

.PHONY: build
build:
	@mkdir -p $(BUILD_DIR)
	$(GO) build -o $(BUILD_DIR)/$(APP_NAME) .

.PHONY: dev
dev:
	@echo "Building in development mode..."
	@mkdir -p $(BUILD_DIR)
	$(GO) build -race -o $(BUILD_DIR)/$(APP_NAME) .
	@echo "Starting $(APP_NAME) in development mode..."
	./$(BUILD_DIR)/$(APP_NAME)

.PHONY: prod
prod:
	@echo "Building production binary..."
	@mkdir -p $(BUILD_DIR)
	$(GO) build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME) .
	@echo "Production build created at $(BUILD_DIR)/$(APP_NAME)"

.PHONY: test
test:
	$(GO) test -v ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)
