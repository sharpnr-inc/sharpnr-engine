# =========================================================
# SharpNR Engine Makefile
# =========================================================

# Application binary name
APP_NAME=sharpnr-engine

# Build output directory
BUILD_DIR=build

# Remote deployment server
SERVER=sharpnrllp@192.168.1.51

# Remote application directory
APP_DIR=/opt/apps/sharpnr-engine

# Marks these as commands instead of file names
.PHONY: build run clean deploy restart status logs scp install

# =========================================================
# Build Linux production binary into build/ folder
#
# Usage:
# make build
# =========================================================
build:
	mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME) .

# =========================================================
# Run application locally in development mode
#
# Usage:
# make run
# =========================================================
run:
	go run .

# =========================================================
# Remove build directory and compiled binaries
#
# Usage:
# make clean
# =========================================================
clean:
	rm -rf $(BUILD_DIR)

# =========================================================
# Execute deployment script
#
# Usage:
# make deploy
# =========================================================
deploy:
	./deploy-be.sh

# =========================================================
# Restart remote systemd service
#
# Usage:
# make restart
# =========================================================
restart:
	ssh $(SERVER) "sudo systemctl restart $(APP_NAME)"

# =========================================================
# Check remote service status
#
# Usage:
# make status
# =========================================================
status:
	ssh $(SERVER) "sudo systemctl status $(APP_NAME) --no-pager"

# =========================================================
# Stream live logs from remote server
#
# Usage:
# make logs
# =========================================================
logs:
	ssh $(SERVER) "journalctl -u $(APP_NAME) -f"

# =========================================================
# Upload compiled binary to remote home directory
#
# Usage:
# make scp
# =========================================================
scp:
	scp $(BUILD_DIR)/$(APP_NAME) $(SERVER):~/

# =========================================================
# Move uploaded binary into application directory
# and make it executable
#
# Usage:
# make install
# =========================================================
install:
	ssh $(SERVER) "\
		mv ~/$(APP_NAME) $(APP_DIR)/$(APP_NAME) && \
		chmod +x $(APP_DIR)/$(APP_NAME)"