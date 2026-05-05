#!/bin/bash
set -e

SERVER="sharpnrllp@192.168.1.51"
APP_DIR="/opt/apps/sharpnr-engine"
BACKUP_DIR="/media/sharpnrllpDrive/engines"
TIMESTAMP=$(date +"%Y%m%d_%H%M%S")

echo "▶ Building sharpnr-engine..."
GOOS=linux GOARCH=amd64 go build -o sharpnr-engine .

echo "▶ Backing up existing engine to $BACKUP_DIR/$TIMESTAMP..."
ssh $SERVER "
    mkdir -p $BACKUP_DIR/$TIMESTAMP
    cp $APP_DIR/sharpnr-engine $BACKUP_DIR/$TIMESTAMP/sharpnr-engine 2>/dev/null || true
    cp -r $APP_DIR/sharp_mailer_engine $BACKUP_DIR/$TIMESTAMP/sharp_mailer_engine 2>/dev/null || true
    echo '✅ Backup done: $TIMESTAMP'
"

echo "▶ Uploading new binary..."
scp sharpnr-engine $SERVER:$APP_DIR/sharpnr-engine

echo "▶ Uploading email templates..."
ssh $SERVER "rm -rf $APP_DIR/sharp_mailer_engine"
scp -r sharp_mailer_engine $SERVER:$APP_DIR/sharp_mailer_engine

echo "▶ Restarting service..."
ssh $SERVER "sudo systemctl restart sharpnr-engine"

echo "▶ Waiting for service to stabilize..."
sleep 3

echo "▶ Checking status..."
ssh $SERVER "sudo systemctl status sharpnr-engine --no-pager"

echo "✅ sharpnr-engine deployed! Backup saved at $BACKUP_DIR/$TIMESTAMP"