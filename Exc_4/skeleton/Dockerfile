# --- Stage 1: Builder ---
FROM golang:1.25 AS builder
WORKDIR /app

COPY . .

RUN chmod +x scripts/build-application.sh && ./scripts/build-application.sh

# --- Stage 2: Runner ---
FROM alpine:latest
WORKDIR /app

# Copy compiled binary from builder
COPY --from=builder /app/ordersystem /app/ordersystem

# Optional: verify it's executable
RUN chmod +x /app/ordersystem && ls -l /app/ordersystem

EXPOSE 3000

CMD ["/app/ordersystem"]
