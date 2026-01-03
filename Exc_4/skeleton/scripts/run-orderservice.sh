#!/bin/sh
docker rm -f orderservice 2>/dev/null || true
docker run -d \
  --name orderservice \
  --network ordersystem-net \
  -e POSTGRES_DB=order \
  -e POSTGRES_USER=docker \
  -e POSTGRES_PASSWORD=docker \
  -e POSTGRES_TCP_PORT=5432 \
  -e DB_HOST=postgres18 \
  -p 3001:3000 \
  orderservice:latest
