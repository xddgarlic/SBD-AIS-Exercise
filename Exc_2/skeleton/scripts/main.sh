#!/bin/sh
set -e

echo "🟢 Starting full Ordersystem stack..."

# Step 0: Make sure helper scripts are executable
chmod +x ./scripts/*.sh 2>/dev/null || true

# Step 1: Run Postgres
echo "📦 Starting Postgres..."
#!/bin/bash
./scripts/run-postgres.sh

# Wait a few seconds for Postgres to initialize
echo "⏳ Waiting 5 seconds for Postgres to be ready..."
sleep 5

# Step 2: Build Orderservice
echo "🏗️ Building Orderservice Docker image..."
./scripts/build-orderservice.sh

# Step 3: Run Orderservice
echo "🚀 Starting Orderservice container..."
./scripts/run-orderservice.sh

echo "✅ Ordersystem stack is up and running!"
echo "API available at http://localhost:3001/"
