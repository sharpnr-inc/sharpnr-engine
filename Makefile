# =========================================================
# SharpNR Engine Makefile
# =========================================================

APP_NAME=sharpnr-engine
BUILD_DIR=build

# =========================================================
# Version info (injected into the binary via -ldflags)
# VERSION  -> nearest git tag + commits since (e.g. v1.0.0-alpha-772-gc87f028)
# COMMIT   -> short SHA; BUILD_TIME -> UTC build timestamp
# =========================================================
META_PKG=sharpnr-engine/meta
VERSION?=$(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
COMMIT?=$(shell git rev-parse --short HEAD 2>/dev/null || echo unknown)
BUILD_TIME?=$(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS=-s -w \
	-X '$(META_PKG).Version=$(VERSION)' \
	-X '$(META_PKG).Commit=$(COMMIT)' \
	-X '$(META_PKG).BuildTime=$(BUILD_TIME)'

.PHONY: build run clean lint test watch version

# =========================================================
# Build Linux production binary into build/ folder
# =========================================================
build:
	mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME) .
	@echo "✅ Built $(APP_NAME) $(VERSION) ($(COMMIT)) — $$(du -sh $(BUILD_DIR)/$(APP_NAME) | cut -f1)"

# =========================================================
# Print the version that a build would embed
# =========================================================
version:
	@echo "Version:    $(VERSION)"
	@echo "Commit:     $(COMMIT)"
	@echo "Build time: $(BUILD_TIME)"

# =========================================================
# Run application locally in development mode
# =========================================================
run:
	go run .

# =========================================================
# Run with hot reload (requires air)
# Install: go install github.com/air-verse/air@latest
# =========================================================
watch:
	air

# =========================================================
# Remove build directory and compiled binaries
# =========================================================
clean:
	rm -rf $(BUILD_DIR)

# =========================================================
# Run tests
# =========================================================
test:
	go test ./... -v

# =========================================================
# Run linter (requires golangci-lint)
# Install: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
# =========================================================
lint:
	golangci-lint run ./...

# =========================================================
# NOTE: Deployment is fully automated by Jenkins (see Jenkinsfile).
# Jenkins and the app run on the same machine, so the old ssh/scp
# remote targets and deploy-be.sh have been removed. This Makefile is
# now for local development, building, and versioning only.
# =========================================================


# ==========================================
# Environment setup
# ==========================================
ENV ?= development
ENV_FILE = .env.$(ENV)

ifneq (,$(wildcard $(ENV_FILE)))
    include $(ENV_FILE)
    export
endif

DB_URL = postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)
MIGRATIONS_PATH = migrations

# ==========================================
# Migration commands
# ==========================================

## Run all pending migrations
migrate-up:
	@echo "▶ Running migrations ($(ENV))..."
	@migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up
	@echo "✓ Migrations applied"

## Rollback last migration
migrate-down:
	@echo "▼ Rolling back 1 migration ($(ENV))..."
	@migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down 1
	@echo "✓ Rolled back"

## Rollback ALL migrations
migrate-down-all:
	@echo "▼ Rolling back ALL migrations ($(ENV))..."
	@migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down -all
	@echo "✓ All rolled back"

## Show current migration version
migrate-version:
	@migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" version

## Create new migration: make migrate-create name=add_cost_center
migrate-create:
	@if [ -z "$(name)" ]; then echo "❌ Usage: make migrate-create name=your_migration_name"; exit 1; fi
	@migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $(name)
	@echo "✓ Created migration: $(name)"

## Force a version (fix dirty state): make migrate-force v=1
migrate-force:
	@if [ -z "$(v)" ]; then echo "❌ Usage: make migrate-force v=1"; exit 1; fi
	@migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" force $(v)
	@echo "✓ Forced version: $(v)"

## Go to a specific version: make migrate-goto v=2
migrate-goto:
	@if [ -z "$(v)" ]; then echo "❌ Usage: make migrate-goto v=2"; exit 1; fi
	@migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" goto $(v)
	@echo "✓ Migrated to version: $(v)"

.PHONY: migrate-up migrate-down migrate-down-all migrate-version migrate-create migrate-force migrate-goto