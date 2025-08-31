#!/bin/bash
# Test the deployment locally before pushing
set -e

echo "=== Testing deployment build ==="

# Build the binary like the CI does
echo "Building Linux binary..."
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o go-db-demo ./web/cmd

# Check if binary was created
if [ -f "go-db-demo" ]; then
    echo "✅ Binary built successfully ($(du -h go-db-demo | cut -f1))"
    
    # Test if it can start (won't work on non-Linux, but we can check if it's valid)
    if file go-db-demo | grep -q "ELF.*executable"; then
        echo "✅ Binary is valid Linux executable"
    else
        echo "❌ Binary format looks wrong"
        exit 1
    fi
else
    echo "❌ Binary build failed"
    exit 1
fi

# Clean up
rm go-db-demo

echo ""
echo "=== Pre-deployment checklist ==="
echo "✅ Code builds successfully"
echo "✅ Binary is valid Linux executable"
echo ""
echo "Ready to deploy! Make sure you have:"
echo "1. Set up the deploy user on your server"
echo "2. Added SERVER_HOST and DEPLOY_SSH_KEY to GitHub secrets"  
echo "3. SSH access working to deploy@YOUR_SERVER"
echo ""
echo "Then push to main branch to deploy!"
