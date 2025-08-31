# Setup GitHub Self-Hosted Runner

This allows GitHub Actions to run on your server/network where it can access your Tailscale hostname.

## Step 1: On Your Server

```bash
# Create runner directory
mkdir -p ~/actions-runner && cd ~/actions-runner

# Download the runner (adjust for your architecture)
curl -o actions-runner-linux-x64-2.319.1.tar.gz -L https://github.com/actions/runner/releases/download/v2.319.1/actions-runner-linux-x64-2.319.1.tar.gz

# Extract it
tar xzf ./actions-runner-linux-x64-2.319.1.tar.gz
```

## Step 2: Get Token from GitHub

1. Go to your repo: https://github.com/ebonfiglio/go-db-demo
2. Settings → Actions → Runners → New self-hosted runner
3. Copy the configuration command they show you

## Step 3: Configure Runner

```bash
# Run the config command from GitHub (it will look like this):
./config.sh --url https://github.com/ebonfiglio/go-db-demo --token YOUR_TOKEN_HERE

# Start the runner
./run.sh
```

## Step 4: Update Workflow

Change `runs-on: ubuntu-latest` to `runs-on: self-hosted` in your workflow file.

## Benefits:
- No SSH needed - direct local deployment
- Works with Tailscale hostnames
- Faster builds (no network transfer)
- More secure (runs on your infrastructure)
