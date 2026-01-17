package server

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jjr/mcp.github.agentic/internal/tools"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type MCPServer struct {
	server   *mcp.Server
	Registry *tools.Registry
	Config   *Config
}

func NewServer(cfg *Config) *MCPServer {
	// Initialize server with basic options
	s := mcp.NewServer(&mcp.Implementation{
		Name:    "infra-mcp-server",
		Version: "0.1.0",
	}, nil)

	// Create registry wrapper
	registry := tools.NewRegistry(s)
	tools.RegisterHealthTool(registry)
	tools.RegisterAzureFirewallTool(registry)

	// Add Resources support
	// 1. Data Model / Schema Specification (FR-FW-017)
	// Uses the centralized schema function from tools package
	s.AddResource(&mcp.Resource{
		Name:        "Azure Firewall Rule Schema",
		URI:         "mcp://azure-firewall/schema",
		Description: "The JSON schema defining the required structure for firewall rules.",
		MIMEType:    "application/schema+json",
	}, func(ctx context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
		return &mcp.ReadResourceResult{
			Contents: []*mcp.ResourceContents{
				{
					URI:      "mcp://azure-firewall/schema",
					MIMEType: "application/schema+json",
					Text:     tools.GetAzureFirewallSchema(),
				},
			},
		}, nil
	})

	return &MCPServer{
		server:   s,
		Registry: registry,
		Config:   cfg,
	}
}

func (s *MCPServer) Run(ctx context.Context) error {
	switch s.Config.Transport {
	case TransportHTTP:
		return s.runHTTP(ctx)
	case TransportStdio:
		return s.runStdio(ctx)
	default:
		return fmt.Errorf("unknown transport: %s", s.Config.Transport)
	}
}

func (s *MCPServer) runStdio(ctx context.Context) error {
	fmt.Fprintln(os.Stderr, "Starting MCP server on stdio...")
	return s.server.Run(ctx, &mcp.StdioTransport{})
}

func (s *MCPServer) runHTTP(ctx context.Context) error {
	fmt.Fprintf(os.Stderr, "Starting MCP server on http %s (Streamable HTTP /mcp)...\n", s.Config.Addr)

	handler := mcp.NewStreamableHTTPHandler(func(request *http.Request) *mcp.Server {
		return s.server
	}, nil)

	mux := http.NewServeMux()
	mux.Handle("/mcp", handler)
	mux.Handle("/messages/", handler) // Streamable handler needs /messages too for SSE mode

	// Basic Health check endpoint for US2
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	srv := &http.Server{
		Addr:    s.Config.Addr,
		Handler: mux,
	}

	errChan := make(chan error, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Fprintln(os.Stderr, "Shutting down HTTP server...")
		return srv.Shutdown(context.Background())
	case err := <-errChan:
		return err
	}
}
