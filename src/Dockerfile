# Stage 1: Build
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o mcp-server ./cmd/mcp-server

# Stage 2: Final
FROM alpine:latest

# Install CA certificates for HTTPS support
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/mcp-server .

# Expose port (default for HTTP mode)
EXPOSE 8080

# Environment variable defaults
ENV MCP_TRANSPORT=stdio
ENV MCP_ADDR=:8080

# Run the binary
ENTRYPOINT ["./mcp-server"]
