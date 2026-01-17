package server

import (
	"flag"
	"os"
)

type TransportMode string

const (
	TransportStdio TransportMode = "stdio"
	TransportHTTP  TransportMode = "http"
)

// Config holds the server configuration.
// Matches TypeScript ServerConfig interface.
type Config struct {
	Transport TransportMode
	Addr      string
}

// NewConfig creates a new Config with values from flags and environment variables.
// Priority: Environment Variables > CLI Flags > Defaults
// Matches TypeScript parseConfig behavior.
func NewConfig() *Config {
	c := &Config{}

	// Flags with defaults
	transportFlag := flag.String("transport", "stdio", "Transport mode: stdio or http")
	addrFlag := flag.String("addr", ":8080", "Address to listen on (for http transport)")
	flag.Parse()

	// Apply flag values first
	c.Transport = TransportMode(*transportFlag)
	c.Addr = *addrFlag

	// Environment Variables override flags (higher priority)
	if envTransport := os.Getenv("MCP_TRANSPORT"); envTransport != "" {
		c.Transport = TransportMode(envTransport)
	}

	if envAddr := os.Getenv("MCP_ADDR"); envAddr != "" {
		c.Addr = envAddr
	}

	return c
}
