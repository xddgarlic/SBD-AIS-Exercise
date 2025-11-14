#!/bin/sh
set -e

echo "🛑 Stopping any running containers..."
docker compose down

echo "🏗️ Building Ordersystem Docker image..."
docker compose build --no-cache ordersystem

echo "📦 Starting Postgres container..."
docker compose up -d postgres

echo "⏳ Waiting for Postgres to be ready..."

# Poll Postgres until it accepts connections
until docker compose exec -T postgres pg_isready -U docker >/dev/null 2>&1; do
    echo "Waiting for Postgres..."
    sleep 2
done

echo "✅ Postgres is ready!"

echo "🚀 Starting Ordersystem API container..."
# Start the Go API and attach to logs
docker compose up ordersystem


#chmod +x scripts/main.sh
#./scripts/main.sh