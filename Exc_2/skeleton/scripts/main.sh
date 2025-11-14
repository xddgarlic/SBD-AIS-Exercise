#!/bin/sh
set -e

echo "🛑 Stopping any running containers..."
docker compose down

echo "🏗️ Building Ordersystem Docker image..."
docker compose build --no-cache ordersystem

echo "📦 Starting Postgres container..."
docker compose up -d postgres

echo "⏳ Waiting for Postgres to initialize..."
until docker compose exec -T postgres pg_isready -U docker >/dev/null 2>&1; do
    echo "Waiting for Postgres..."
    sleep 2
done
echo "✅ Postgres is ready!"

echo "🚀 Starting Ordersystem, sws, and Traefik..."
docker compose up -d ordersystem sws traefik

# Print service URLs
echo ""
echo "🌐 Services are running! Access them here:"
echo "-------------------------------------------"
echo "Frontend (SWS): http://localhost"
echo "Go API (Ordersystem): http://orders.localhost"
echo "Postgres: running inside Docker container (port 5432 mapped internally)"
echo "Traefik Dashboard: http://localhost:8080"
echo "-------------------------------------------"
echo ""
echo "Logs can be viewed with: docker compose logs -f"
