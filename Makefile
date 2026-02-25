# Variables
APP_NAME := sharpnr-engine
APP_MAIN := ./cmd/server
EC2_USER := ec2-user
EC2_HOST := 13.232.183.253
KEY_PATH := ~/Downloads/sharpnr-engine-key.pem
REMOTE_BINARY := ~/$(APP_NAME)
ENV_FILE_LOCAL := .env
ENV_FILE_REMOTE := ~/.env.production

# Build the Go binary for Linux AMD64
build:
	GOOS=linux GOARCH=amd64 go build -o $(APP_NAME) $(APP_MAIN)

# Copy the binary and .env.production to the EC2 instance
deploy: build
	scp -i $(KEY_PATH) $(APP_NAME) $(EC2_USER)@$(EC2_HOST):$(REMOTE_BINARY)
	scp -i $(KEY_PATH) $(ENV_FILE_LOCAL) $(EC2_USER)@$(EC2_HOST):$(ENV_FILE_REMOTE)

# SSH into EC2
ssh:
	ssh -i $(KEY_PATH) $(EC2_USER)@$(EC2_HOST)

# Run the app on EC2 in foreground
run:
	ssh -i $(KEY_PATH) $(EC2_USER)@$(EC2_HOST) \
		'chmod +x $(REMOTE_BINARY) && APP_ENV=production $(REMOTE_BINARY)'

# Run the app on EC2 in background with nohup
run-bg:
	ssh -i $(KEY_PATH) $(EC2_USER)@$(EC2_HOST) \
		'pkill $(APP_NAME) || true && chmod +x $(REMOTE_BINARY) && nohup env APP_ENV=production $(REMOTE_BINARY) > app.log 2>&1 &'

# Full deploy and run in background
deploy-run: deploy run-bg
