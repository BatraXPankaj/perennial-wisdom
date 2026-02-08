#!/usr/bin/env bash
# Quick redeploy — rebuild binary and push to VM
# Usage: ./redeploy-azure.sh

set -euo pipefail

RESOURCE_GROUP="stoic-wisdom-rg"
VM_NAME="pbatra-vm"
VM_USER="azureuser"
VM_FQDN="${VM_NAME}.eastus.cloudapp.azure.com"
VM_SSH="${VM_USER}@${VM_FQDN}"
APP_DIR="/opt/perennial-wisdom"
BINARY_NAME="perennial-wisdom"

echo "→ Building Go binary for linux/amd64"
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o "${BINARY_NAME}-linux" .

echo "→ Stopping service"
ssh "$VM_SSH" "sudo systemctl stop ${BINARY_NAME}"

echo "→ Uploading binary and templates"
scp "${BINARY_NAME}-linux" "${VM_SSH}:${APP_DIR}/${BINARY_NAME}"
scp -r templates/ "${VM_SSH}:${APP_DIR}/templates/"
ssh "$VM_SSH" "chmod +x ${APP_DIR}/${BINARY_NAME}"

echo "→ Restarting service"
ssh "$VM_SSH" "sudo systemctl start ${BINARY_NAME}"

rm -f "${BINARY_NAME}-linux"

echo "✓ Redeployed: http://${VM_FQDN}"
