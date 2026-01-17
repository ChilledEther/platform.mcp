package utils

import (
	"net"
)

// IsValidIPOrCIDR checks if the string is a valid IPv4/v6 address or CIDR block.
func IsValidIPOrCIDR(s string) bool {
	if net.ParseIP(s) != nil {
		return true
	}
	_, _, err := net.ParseCIDR(s)
	return err == nil
}
