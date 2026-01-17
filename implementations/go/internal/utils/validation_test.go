package utils

import "testing"

func TestIsValidIPOrCIDR(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"127.0.0.1", true},
		{"10.0.0.0/24", true},
		{"2001:db8::/32", true},
		{"8.8.8.8", true},
		{"999.999.999.999", false},
		{"invalid", false},
		{"10.0.0.256", false},
		{"10.0.0.0/33", false},
		{"", false},
	}

	for _, tt := range tests {
		got := IsValidIPOrCIDR(tt.input)
		if got != tt.want {
			t.Errorf("IsValidIPOrCIDR(%q) = %v; want %v", tt.input, got, tt.want)
		}
	}
}
