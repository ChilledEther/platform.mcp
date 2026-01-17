package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/jjr/mcp.github.agentic/internal/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Fprintln(os.Stderr, "Received signal, shutting down...")
		cancel()
	}()

	cfg := server.NewConfig()
	srv := server.NewServer(cfg)
	if err := srv.Run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Server exited with error: %v\n", err)
		os.Exit(1)
	}
}
