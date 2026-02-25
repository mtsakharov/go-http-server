package handlers

import (
	"testing"

	"github.com/mtsakharov/go-http-server/internal/server/httpcore"
)

func TestUserAgent_Present(t *testing.T) {
	req := httpcore.Request{
		Path:    "/user-agent",
		Headers: map[string]string{"user-agent": "Mozilla/5.0 (X11; Linux x86_64)"},
	}

	resp := UserAgent(req)

	if resp.Status != httpcore.StatusOK {
		t.Errorf("status = %d, want %d", resp.Status, httpcore.StatusOK)
	}
	if resp.ContentType != "text/plain" {
		t.Errorf("content-type = %q, want %q", resp.ContentType, "text/plain")
	}
	if string(resp.Body) != "Mozilla/5.0 (X11; Linux x86_64)" {
		t.Errorf("body = %q, want %q", resp.Body, "Mozilla/5.0 (X11; Linux x86_64)")
	}
}

func TestUserAgent_Missing(t *testing.T) {
	req := httpcore.Request{
		Path:    "/user-agent",
		Headers: map[string]string{},
	}

	resp := UserAgent(req)

	if resp.Status != httpcore.StatusOK {
		t.Errorf("status = %d, want %d", resp.Status, httpcore.StatusOK)
	}
	if string(resp.Body) != "" {
		t.Errorf("body = %q, want empty when User-Agent header is missing", resp.Body)
	}
}
