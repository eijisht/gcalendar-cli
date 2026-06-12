package internal

import "testing"

func TestParseCommand(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{"create maps to c", []string{"gcal", "create"}, "c"},
		{"read maps to r", []string{"gcal", "read"}, "r"},
		{"remove maps to d", []string{"gcal", "remove"}, "d"},
		{"reset passes through", []string{"gcal", "reset"}, "reset"},
		{"no command", []string{"gcal"}, ""},
		{"unknown passthrough", []string{"gcal", "bogus"}, "bogus"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseCommand(tt.args); got != tt.want {
				t.Errorf("ParseCommand(%v) = %q, want %q", tt.args, got, tt.want)
			}
		})
	}
}
