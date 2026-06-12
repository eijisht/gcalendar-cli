package internal

import (
	"os"
	"path/filepath"
	"testing"
)

// tokenFromFile already returns errors, so these tests pass before and after
// the commit-4 refactor. They lock in the "return the error, don't exit"
// contract for the file-loading path.

func TestTokenFromFile_MissingFile(t *testing.T) {
	_, err := tokenFromFile(filepath.Join(t.TempDir(), "does-not-exist.json"))
	if err == nil {
		t.Fatal("expected an error for a missing token file, got nil")
	}
}

func TestTokenFromFile_MalformedJSON(t *testing.T) {
	path := filepath.Join(t.TempDir(), "token.json")
	if err := os.WriteFile(path, []byte("this is not json"), 0600); err != nil {
		t.Fatalf("setup: %v", err)
	}

	_, err := tokenFromFile(path)
	if err == nil {
		t.Fatal("expected an error for malformed JSON, got nil")
	}
}

func TestTokenFromFile_Valid(t *testing.T) {
	path := filepath.Join(t.TempDir(), "token.json")
	contents := `{"access_token":"abc123","token_type":"Bearer"}`
	if err := os.WriteFile(path, []byte(contents), 0600); err != nil {
		t.Fatalf("setup: %v", err)
	}

	tok, err := tokenFromFile(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tok.AccessToken != "abc123" {
		t.Errorf("AccessToken = %q, want %q", tok.AccessToken, "abc123")
	}
}

// TestGetCalendarService_MissingCredentials is the actual commit-4 test: the
// function must RETURN an error when credentials.json is absent, instead of
// calling log.Fatalf (which would os.Exit and kill the test binary).
//
// This test only passes AFTER you refactor GetCalendarService to return errors.
// Until then it will abort the whole test run.
func TestGetCalendarService_MissingCredentials(t *testing.T) {
	// Run from an empty dir so the hard-coded "credentials.json" path is missing.
	orig, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(t.TempDir()); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.Chdir(orig) })

	if _, err := GetCalendarService(); err == nil {
		t.Fatal("expected an error when credentials.json is missing, got nil")
	}
}
