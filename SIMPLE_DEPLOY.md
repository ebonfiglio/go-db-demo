# Simple Deployment Guide

## Quick Setup (5 minutes)

### 1. Prepare your Ubuntu server
Run this on your Ubuntu server:
```bash
curl -O https://raw.githubusercontent.com/ebonfiglio/go-db-demo/main/setup-server.sh
chmod +x setup-server.sh
./setup-server.sh
```

### 2. Set up SSH access
```bash
# On your local machine, generate SSH key if you don't have one:
ssh-keygen -t ed25519 -f ~/.ssh/deploy_key

# Copy public key to server:
ssh-copy-id -i ~/.ssh/deploy_key.pub deploy@YOUR_SERVER_IP
```

### 3. Configure GitHub Secrets
In your GitHub repository, go to Settings → Secrets and add:
- `SERVER_HOST`: Your server's IP address
- `DEPLOY_SSH_KEY`: Contents of your private key (`cat ~/.ssh/deploy_key`)

### 4. Deploy!
Push to main branch and watch it deploy automatically.

## How it works

1. **Build**: Compiles Go binary for Linux
2. **Deploy**: Copies binary to server in user's home directory
3. **Run**: Starts the application with health checks
4. **Verify**: Confirms deployment worked

## Files created on server:
- `~/app/go-db-demo` - The application binary
- `~/app/.env` - Environment configuration
- `~/app/app.log` - Application logs
- `~/app/app.pid` - Process ID file

## Troubleshooting

**Check if app is running:**
```bash
ssh deploy@YOUR_SERVER curl http://localhost:8080/healthz
```

**View logs:**
```bash
ssh deploy@YOUR_SERVER tail -f ~/app/app.log
```

**Manual restart:**
```bash
ssh deploy@YOUR_SERVER 'cd ~/app && pkill -f go-db-demo && nohup ./go-db-demo > app.log 2>&1 &'
```

That's it! Simple, reliable, no complex systemd or sudo required.
