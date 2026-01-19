# Build stage
FROM golang:1.25.5-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the MCP server binary
RUN CGO_ENABLED=0 GOOS=linux go build -o platform-mcp ./cmd/platform-mcp/main.go

# Final stage
FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/platform-mcp .

# Run the MCP server
ENTRYPOINT ["./platform-mcp"]
