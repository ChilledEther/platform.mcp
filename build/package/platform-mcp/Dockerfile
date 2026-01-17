# Build stage
FROM golang:1.25.5-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o platform-mcp ./cmd/platform-mcp/main.go

# Final stage
FROM alpine:latest

WORKDIR /

# Copy the binary from the builder stage
COPY --from=builder /app/platform-mcp /platform-mcp

# Expose no ports as it uses stdio
ENTRYPOINT ["/platform-mcp"]
