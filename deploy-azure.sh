#!/usr/bin/env bash
# Perennial Wisdom — Azure VM Deployment Script
# First-time setup: creates a B1s VM and deploys the Go binary
#
# Prerequisites: az login (service principal or interactive)
# Usage: ./deploy-azure.sh
#
# Estimated cost: ~$7.50/month (B1s VM, covered by VS Pro credits)

set -euo pipefail

# ---- Configuration ----
RESOURCE_GROUP="stoic-wisdom-rg"
LOCATION="eastus"
VM_NAME="pbatra-vm"
VM_SIZE="Standard_B1s"
VM_USER="azureuser"
VM_FQDN="${VM_NAME}.${LOCATION}.cloudapp.azure.com"
APP_DIR="/opt/perennial-wisdom"
BINARY_NAME="perennial-wisdom"

echo "=== Perennial Wisdom — Azure VM Deployment ==="
echo ""

# ---- 1. Check if VM already exists ----
if az vm show -g "$RESOURCE_GROUP" -n "$VM_NAME" &>/dev/null; then
  echo "→ VM '$VM_NAME' already exists. Use ./redeploy-azure.sh for updates."
  echo "  URL: http://$VM_FQDN"
  exit 0
fi

# ---- 2. Create VM (B1s ~$7.50/mo) ----
echo "→ Creating VM: $VM_NAME ($VM_SIZE) in $RESOURCE_GROUP"
az vm create \
  --resource-group "$RESOURCE_GROUP" \
  --name "$VM_NAME" \
  --image Ubuntu2404 \
  --size "$VM_SIZE" \
  --admin-username "$VM_USER" \
  --generate-ssh-keys \
  --public-ip-sku Standard \
  --public-ip-address-dns-name "$VM_NAME" \
  --nsg-rule SSH \
  --output none

# ---- 3. Open ports 80, 8080 ----
echo "→ Opening ports 80 and 8080"
az vm open-port -g "$RESOURCE_GROUP" -n "$VM_NAME" --port 80 --priority 900 --output none
az vm open-port -g "$RESOURCE_GROUP" -n "$VM_NAME" --port 8080 --priority 901 --output none

# ---- 4. Cross-compile Go binary ----
echo "→ Building Go binary for linux/amd64"
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o "${BINARY_NAME}-linux" .

# ---- 5. Deploy to VM ----
echo "→ Deploying to $VM_FQDN"
VM_SSH="${VM_USER}@${VM_FQDN}"

ssh -o StrictHostKeyChecking=no "$VM_SSH" "sudo mkdir -p ${APP_DIR}/data && sudo chown -R ${VM_USER}:${VM_USER} ${APP_DIR}"
scp "${BINARY_NAME}-linux" "${VM_SSH}:${APP_DIR}/${BINARY_NAME}"
scp -r templates/ "${VM_SSH}:${APP_DIR}/templates/"
ssh "$VM_SSH" "chmod +x ${APP_DIR}/${BINARY_NAME}"

# ---- 6. Create systemd service ----
echo "→ Setting up systemd service"
ssh "$VM_SSH" "cat > /tmp/${BINARY_NAME}.service << 'UNIT'
[Unit]
Description=Perennial Wisdom API
After=network.target

[Service]
Type=simple
User=${VM_USER}
WorkingDirectory=${APP_DIR}
ExecStart=${APP_DIR}/${BINARY_NAME}
Environment=DB_PATH=${APP_DIR}/data/wisdom.db
Environment=GIN_MODE=release
Environment=PORT=8080
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
UNIT
sudo mv /tmp/${BINARY_NAME}.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable ${BINARY_NAME}
sudo systemctl start ${BINARY_NAME}"

# ---- 7. Install & configure nginx reverse proxy ----
echo "→ Installing nginx"
ssh "$VM_SSH" "sudo apt-get update -qq && sudo DEBIAN_FRONTEND=noninteractive apt-get install -y -qq nginx > /dev/null 2>&1"

echo "→ Configuring nginx reverse proxy (port 80 → 8080)"
ssh "$VM_SSH" "sudo tee /etc/nginx/sites-available/perennial-wisdom > /dev/null << 'NGINX'
server {
    listen 80;
    server_name ${VM_FQDN};

    location / {
        proxy_pass http://127.0.0.1:8080;
        proxy_set_header Host \$host;
        proxy_set_header X-Real-IP \$remote_addr;
        proxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto \$scheme;
    }
}
NGINX
sudo ln -sf /etc/nginx/sites-available/perennial-wisdom /etc/nginx/sites-enabled/
sudo rm -f /etc/nginx/sites-enabled/default
sudo nginx -t && sudo systemctl restart nginx"

# ---- 8. Verify ----
echo "→ Verifying deployment..."
sleep 2
HTTP_STATUS=$(curl -s -o /dev/null -w '%{http_code}' "http://${VM_FQDN}/")
if [ "$HTTP_STATUS" = "200" ]; then
  echo ""
  echo "=== Deployment complete! ==="
  echo "URL:  http://${VM_FQDN}"
else
  echo "⚠ Got HTTP $HTTP_STATUS — check logs with: ssh ${VM_SSH} 'sudo journalctl -u ${BINARY_NAME} -n 20'"
fi

# Cleanup build artifact
rm -f "${BINARY_NAME}-linux"

echo ""
echo "Useful commands:"
echo "  SSH:      ssh ${VM_SSH}"
echo "  Logs:     ssh ${VM_SSH} 'sudo journalctl -u ${BINARY_NAME} -f'"
echo "  Redeploy: ./redeploy-azure.sh"
echo "  Stop:     az vm deallocate -g $RESOURCE_GROUP -n $VM_NAME"
echo "  Delete:   az vm delete -g $RESOURCE_GROUP -n $VM_NAME --yes"
