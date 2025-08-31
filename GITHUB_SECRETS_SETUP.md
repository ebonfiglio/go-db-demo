# 🔑 GitHub Secrets Setup Guide

The deployment failed because GitHub secrets aren't set up yet. Here's how to fix it:

## Step 1: Set up your Ubuntu server

On your Ubuntu server, run:
```bash
# Create deploy user
sudo useradd -m -s /bin/bash deploy 2>/dev/null || echo "Deploy user exists"
sudo -u deploy mkdir -p /home/deploy/app
sudo -u deploy mkdir -p /home/deploy/.ssh
```

## Step 2: Create SSH key pair

On your local machine:
```bash
# Generate SSH key for deployment
ssh-keygen -t ed25519 -f ~/.ssh/deploy_key -N ""

# Copy public key to server
ssh-copy-id -i ~/.ssh/deploy_key.pub deploy@YOUR_SERVER_IP

# Test connection
ssh -i ~/.ssh/deploy_key deploy@YOUR_SERVER_IP "echo 'SSH works!'"
```

## Step 3: Add GitHub Secrets

1. Go to your GitHub repository
2. Click **Settings** → **Secrets and variables** → **Actions**
3. Click **New repository secret**

Add these two secrets:

### SECRET 1: `TS_HOST`
- **Name**: `TS_HOST`
- **Value**: Your server's IP address (e.g., `192.168.1.100`)

### SECRET 2: `DEPLOY_SSH_KEY` 
- **Name**: `DEPLOY_SSH_KEY`
- **Value**: Contents of your private key
```bash
# Copy this output:
cat ~/.ssh/deploy_key
```

## Step 4: Test deployment

Push any change to main branch:
```bash
git add .
git commit -m "Test deployment"
git push origin main
```

## Troubleshooting

**SSH connection issues:**
```bash
# Test SSH manually:
ssh -i ~/.ssh/deploy_key deploy@YOUR_SERVER_IP

# If it fails, check:
# 1. Server IP is correct
# 2. deploy user exists on server
# 3. SSH service is running on server
```

**Secret issues:**
- Make sure secret names are exactly: `TS_HOST` and `DEPLOY_SSH_KEY`
- No extra spaces or characters
- Private key should start with `-----BEGIN OPENSSH PRIVATE KEY-----`

That's it! Once secrets are set, deployment will work automatically.
