# ----- Build Stage -----
FROM golang:1.24.4-alpine AS builder
WORKDIR /app

# Install git (required for modules) and ca-certificates
RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY src ./src

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o woka-server ./src/cmd/server
# ----- Final Stage -----
FROM alpine:3.14
WORKDIR /root/

# Install certificates
RUN apk add --no-cache ca-certificates && update-ca-certificates

# Copy binary
COPY --from=builder /app/woka-server .

# Expose application port
EXPOSE 8080
CMD ["./woka-server"]