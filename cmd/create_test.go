package cmd

import "testing"

func TestParseDate(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"valid datetime", "2024-01-02 15:04", false},
		{"empty string", "", true},
		{"wrong format", "01/02/2024 3pm", true},
		{"date only", "2024-01-02", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseDate(tt.input)
			if tt.wantErr && err == nil {
				t.Errorf("ParseDate(%q): expected an error, got nil", tt.input)
			}
			if !tt.wantErr && err != nil {
				t.Errorf("ParseDate(%q): unexpected error: %v", tt.input, err)
			}
		})
	}
}

func TestParseDate_Fields(t *testing.T) {
	got, err := ParseDate("2024-03-09 08:30")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Year() != 2024 || got.Month() != 3 || got.Day() != 9 {
		t.Errorf("date = %v, want 2024-03-09", got)
	}
	if got.Hour() != 8 || got.Minute() != 30 {
		t.Errorf("time = %02d:%02d, want 08:30", got.Hour(), got.Minute())
	}
}
