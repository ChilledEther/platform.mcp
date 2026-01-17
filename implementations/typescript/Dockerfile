# Stage 1: Build
FROM oven/bun:alpine AS builder
WORKDIR /app

# Copy source (files now at root, no src/ subdirectory)
COPY package.json bun.lock* ./
COPY index.ts ./
COPY server ./server
COPY tools ./tools
COPY resources ./resources
COPY utils ./utils

# Compile to a standalone executable for Musl (Alpine)
# This removes node_modules entirely and includes the runtime
# --minify: Reduces code size
# --sourcemap=none: Saves space
# --target=bun-linux-x64-musl: Targets Alpine Linux
RUN bun install --frozen-lockfile && \
    bun build --compile --minify --sourcemap=none --target=bun-linux-x64-musl --outfile mcp-server ./index.ts

# Stage 2: Production
FROM alpine:latest
WORKDIR /app

# Install minimal dependencies required by Bun binary (libstdc++, libgcc)
# --no-cache: Keep image size small
RUN apk add --no-cache libstdc++ libgcc

# Copy only the compiled binary
COPY --from=builder /app/mcp-server .

# Create non-root user for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

# Expose port (default for HTTP mode)
EXPOSE 8080

# Environment variable defaults
ENV MCP_TRANSPORT=stdio
ENV MCP_ADDR=:8080
ENV BUN_CONFIG_MAX_MEMORY=128mb

# Run the compiled server
CMD ["./mcp-server"]
