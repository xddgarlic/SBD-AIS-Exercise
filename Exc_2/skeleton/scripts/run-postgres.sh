#!/bin/sh
docker rm -f postgres18 2>/dev/null || true
docker network create ordersystem-net 2>/dev/null || true

docker run -d \
  --name postgres18 \
  --network ordersystem-net \
  -e POSTGRES_DB=order \
  -e POSTGRES_USER=docker \
  -e POSTGRES_PASSWORD=docker \
  -v pgdata:/var/lib/postgresql/18/docker \
  -p 5432:5432 \
  postgres:18
