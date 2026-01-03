#!/bin/sh
set -e

echo "🚧 Building statically linked Go application..."
cd /app

go mod tidy

# Build static binary for Linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /app/ordersystem .

ls -lh /app/ordersystem

echo "✅ Build complete!"