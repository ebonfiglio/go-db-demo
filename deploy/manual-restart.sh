#!/bin/bash
# Manual restart script for go-db-demo
set -e

echo "Stopping current go-db-demo processes..."
pkill -f go-db-demo || true
sleep 2

echo "Starting go-db-demo manually..."
cd /var/www/go-db-demo
nohup ./current > /tmp/go-db-demo.log 2>&1 &
echo $! > /tmp/go-db-demo.pid

sleep 2
echo "Process started. Checking health..."
curl -f http://127.0.0.1:8080/healthz || echo "Health check failed"

echo "Recent logs:"
tail -20 /tmp/go-db-demo.log || true
