#!/bin/bash
# Simple server setup script
# Run this ONCE on your Ubuntu server to prepare for deployments

set -e

echo "=== Setting up Ubuntu server for go-db-demo deployment ==="

# Create deploy user if it doesn't exist
if ! id deploy &>/dev/null; then
    echo "Creating deploy user..."
    sudo useradd -m -s /bin/bash deploy
    echo "Deploy user created"
else
    echo "Deploy user already exists"
fi

# Create app directory
echo "Setting up application directory..."
sudo -u deploy mkdir -p /home/deploy/app
sudo -u deploy mkdir -p /home/deploy/.ssh

echo "=== Setup complete! ==="
echo ""
echo "Next steps:"
echo "1. Add your SSH public key to /home/deploy/.ssh/authorized_keys"
echo "2. Set up your GitHub secrets:"
echo "   - SERVER_HOST: your server's IP address"
echo "   - DEPLOY_SSH_KEY: your private SSH key"
echo "3. Push to main branch to trigger deployment"
echo ""
echo "That's it! No sudo permissions needed for deployment."
