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