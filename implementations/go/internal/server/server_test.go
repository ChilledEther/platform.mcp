package server

import (
	"testing"
)

func TestNewServer(t *testing.T) {
	cfg := &Config{Transport: TransportStdio}
	s := NewServer(cfg)
	if s == nil {
		t.Fatal("NewServer returned nil")
	}
	if s.server == nil {
		t.Error("MCPServer.server is nil")
	}
	if s.Registry == nil {
		t.Error("MCPServer.Registry is nil")
	}
}
