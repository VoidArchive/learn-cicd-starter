package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_ValidHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey test-key-123")

	key, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if key != "test-key-123" {
		t.Errorf("Expected key 'test-key-123', got: %s", key)
	}
}

func TestGetAPIKey_NoAuthHeader(t *testing.T) {
	headers := http.Header{}

	_, err := GetAPIKey(headers)
	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("Expected ErrNoAuthHeaderIncluded, got: %v", err)
	}
}
