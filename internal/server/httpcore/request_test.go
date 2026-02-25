package httpcore

import (
	"testing"
)

func TestParseRequest_ValidGET(t *testing.T) {
	raw := "GET /index.html HTTP/1.1\r\nHost: localhost\r\nAccept: text/html\r\n\r\n"

	req, err := ParseRequest(raw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if req.Method != GET {
		t.Errorf("method = %q, want %q", req.Method, GET)
	}
	if req.Path != "/index.html" {
		t.Errorf("path = %q, want %q", req.Path, "/index.html")
	}
	if req.Headers["host"] != "localhost" {
		t.Errorf("Host header = %q, want %q", req.Headers["host"], "localhost")
	}
	if req.Headers["accept"] != "text/html" {
		t.Errorf("Accept header = %q, want %q", req.Headers["accept"], "text/html")
	}
	if req.Body != "" {
		t.Errorf("body = %q, want empty", req.Body)
	}
}

func TestParseRequest_ValidPOST(t *testing.T) {
	raw := "POST /files/test.txt HTTP/1.1\r\nContent-Length: 13\r\n\r\nHello, World!"

	req, err := ParseRequest(raw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if req.Method != POST {
		t.Errorf("method = %q, want %q", req.Method, POST)
	}
	if req.Path != "/files/test.txt" {
		t.Errorf("path = %q, want %q", req.Path, "/files/test.txt")
	}
	if req.Body != "Hello, World!" {
		t.Errorf("body = %q, want %q", req.Body, "Hello, World!")
	}
}

func TestParseRequest_HeadersNormalizedToLowercase(t *testing.T) {
	raw := "GET / HTTP/1.1\r\nUser-Agent: Go-Test\r\nAccept-Encoding: gzip\r\n\r\n"

	req, err := ParseRequest(raw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if req.Headers["user-agent"] != "Go-Test" {
		t.Errorf("user-agent = %q, want %q", req.Headers["user-agent"], "Go-Test")
	}
	if req.Headers["accept-encoding"] != "gzip" {
		t.Errorf("accept-encoding = %q, want %q", req.Headers["accept-encoding"], "gzip")
	}

	// Original casing should not exist
	if _, ok := req.Headers["User-Agent"]; ok {
		t.Error("header key should be lowercase, but found original casing")
	}
}

func TestParseRequest_MalformedRequestLine(t *testing.T) {
	raw := "INVALID\r\n\r\n"

	_, err := ParseRequest(raw)
	if err == nil {
		t.Fatal("expected error for malformed request line, got nil")
	}
}

func TestParseRequest_EmptyInput(t *testing.T) {
	_, err := ParseRequest("")
	if err == nil {
		t.Fatal("expected error for empty input, got nil")
	}
}

func TestParseRequest_HeaderMissingColon(t *testing.T) {
	raw := "GET / HTTP/1.1\r\nBadHeader\r\nGood-Header: value\r\n\r\n"

	req, err := ParseRequest(raw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Bad header should be skipped
	if _, ok := req.Headers["badheader"]; ok {
		t.Error("header without colon should be skipped")
	}
	if req.Headers["good-header"] != "value" {
		t.Errorf("good-header = %q, want %q", req.Headers["good-header"], "value")
	}
}

func TestParseRequest_NoBody(t *testing.T) {
	raw := "GET / HTTP/1.1\r\nHost: localhost\r\n\r\n"

	req, err := ParseRequest(raw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if req.Body != "" {
		t.Errorf("body = %q, want empty", req.Body)
	}
}
