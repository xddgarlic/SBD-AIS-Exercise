# --- Stage 1: Builder ---
FROM golang:1.25 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ordersystem .

# --- Stage 2: Runner ---
FROM alpine:latest
WORKDIR /app
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/ordersystem /app/ordersystem
RUN chmod +x /app/ordersystem

EXPOSE 3000
CMD ["/app/ordersystem"]
