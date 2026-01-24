# ---------- Builder stage ----------
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go module files from repo root
COPY go.mod go.sum ./
RUN go mod download

# Copy shared and service-specific code
COPY common ./common
COPY proto ./proto
COPY order-service ./order-service

# Build order-service binary
WORKDIR /app/order-service
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app


# ---------- Runtime stage ----------
FROM alpine:3.19

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/order-service/app .

# Expose port (change if your order-service uses a different one)
EXPOSE 50052

# Run service
CMD ["./app"]
