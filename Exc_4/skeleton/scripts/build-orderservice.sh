#!/bin/sh
docker rm -f orderservice 2>/dev/null || true
docker build --no-cache -t orderservice:latest .
