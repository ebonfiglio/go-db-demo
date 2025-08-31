#!/bin/bash
# Alternative deployment script that can be run on the server
set -euo pipefail

DEPLOY_DIR="/var/www/go-db-demo"
BINARY_NAME="go-db-demo"

echo "=== Alternative Deployment Script ==="

# Check if we're in the right directory
if [ ! -d "$DEPLOY_DIR" ]; then
    echo "Error: Deploy directory $DEPLOY_DIR not found"
    exit 1
fi

cd "$DEPLOY_DIR"

# Stop current process
echo "Stopping current processes..."
pkill -f "$BINARY_NAME" || true
sleep 3

# Verify the new binary exists and is executable
if [ ! -x "./current" ]; then
    echo "Error: ./current binary not found or not executable"
    ls -la ./current || true
    exit 1
fi

echo "Starting new process..."
# Start the new process in background
nohup ./current > /tmp/go-db-demo.log 2>&1 &
NEW_PID=$!
echo $NEW_PID > /tmp/go-db-demo.pid

echo "Process started with PID: $NEW_PID"

# Wait a moment for startup
sleep 3

# Test the health endpoint
echo "Testing health endpoint..."
if curl -f http://127.0.0.1:8080/healthz >/dev/null 2>&1; then
    echo "✅ Health check passed"
    echo "✅ Deployment successful!"
else
    echo "❌ Health check failed"
    echo "Recent logs:"
    tail -20 /tmp/go-db-demo.log || true
    exit 1
fi

echo "=== Deployment completed successfully ==="
