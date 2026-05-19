# =========================================================
# SharpNR Engine Makefile
# =========================================================

APP_NAME=sharpnr-engine
BUILD_DIR=build
SERVER=sharpnrllp@192.168.1.51
APP_DIR=/opt/apps/sharpnr-engine

.PHONY: build run clean deploy restart status logs scp install lint test watch env

# =========================================================
# Build Linux production binary into build/ folder
# =========================================================
build:
	mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(APP_NAME) .
	@echo "✅ Build size: $$(du -sh $(BUILD_DIR)/$(APP_NAME) | cut -f1)"

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
# Execute full deployment script
# =========================================================
deploy:
	chmod +x deploy-be.sh
	./deploy-be.sh

# =========================================================
# Restart remote systemd service
# =========================================================
restart:
	ssh $(SERVER) "sudo systemctl restart $(APP_NAME)"

# =========================================================
# Check remote service status
# =========================================================
status:
	ssh $(SERVER) "sudo systemctl status $(APP_NAME) --no-pager"

# =========================================================
# Stream live logs from remote server
# =========================================================
logs:
	ssh $(SERVER) "journalctl -u $(APP_NAME) -f"

# =========================================================
# Show last 50 log lines
# =========================================================
logs-tail:
	ssh $(SERVER) "journalctl -u $(APP_NAME) -n 50 --no-pager"

# =========================================================
# Show only errors in logs
# =========================================================
logs-errors:
	ssh $(SERVER) "journalctl -u $(APP_NAME) -n 100 --no-pager | grep -i 'error\|failed\|panic'"

# =========================================================
# Upload compiled binary to remote home directory
# =========================================================
scp:
	scp $(BUILD_DIR)/$(APP_NAME) $(SERVER):~/

# =========================================================
# Move uploaded binary into application directory
# =========================================================
install:
	ssh $(SERVER) "\
		mv ~/$(APP_NAME) $(APP_DIR)/$(APP_NAME) && \
		chmod +x $(APP_DIR)/$(APP_NAME)"

# =========================================================
# Show remote .env (without secrets)
# =========================================================
env:
	ssh $(SERVER) "grep -v 'PASSWORD\|SECRET\|KEY' $(APP_DIR)/.env"

# =========================================================
# Check if port 8080 is listening on remote
# =========================================================
check:
	ssh $(SERVER) "ss -tlnp | grep 8080 && echo '✅ Port 8080 open' || echo '❌ Port 8080 not found'"

# =========================================================
# Full pipeline: lint → test → build → deploy
# =========================================================
ship: lint test deploy


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