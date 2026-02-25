package handlers

import (
	"bytes"
	"compress/gzip"
	"io"
	"testing"

	"github.com/mtsakharov/go-http-server/internal/server/httpcore"
)

func TestEcho_SimpleText(t *testing.T) {
	req := httpcore.Request{
		Path:    "/echo/hello",
		Headers: map[string]string{},
	}

	resp := Echo(req)

	if resp.Status != httpcore.StatusOK {
		t.Errorf("status = %d, want %d", resp.Status, httpcore.StatusOK)
	}
	if resp.ContentType != "text/plain" {
		t.Errorf("content-type = %q, want %q", resp.ContentType, "text/plain")
	}
	if string(resp.Body) != "hello" {
		t.Errorf("body = %q, want %q", resp.Body, "hello")
	}
	if resp.Encoding != "" {
		t.Errorf("encoding = %q, want empty", resp.Encoding)
	}
}

func TestEcho_EmptyPath(t *testing.T) {
	req := httpcore.Request{
		Path:    "/echo/",
		Headers: map[string]string{},
	}

	resp := Echo(req)

	if resp.Status != httpcore.StatusOK {
		t.Errorf("status = %d, want %d", resp.Status, httpcore.StatusOK)
	}
	if string(resp.Body) != "" {
		t.Errorf("body = %q, want empty", resp.Body)
	}
}

func TestEcho_WithGzipEncoding(t *testing.T) {
	req := httpcore.Request{
		Path:    "/echo/hello",
		Headers: map[string]string{"accept-encoding": "gzip"},
	}

	resp := Echo(req)

	if resp.Status != httpcore.StatusOK {
		t.Errorf("status = %d, want %d", resp.Status, httpcore.StatusOK)
	}
	if resp.Encoding != "gzip" {
		t.Errorf("encoding = %q, want %q", resp.Encoding, "gzip")
	}

	// Decompress and verify content
	reader, err := gzip.NewReader(bytes.NewReader(resp.Body))
	if err != nil {
		t.Fatalf("gzip.NewReader error: %v", err)
	}
	defer reader.Close()

	decompressed, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("ReadAll error: %v", err)
	}

	if string(decompressed) != "hello" {
		t.Errorf("decompressed body = %q, want %q", decompressed, "hello")
	}
}

func TestEcho_WithMultipleEncodings(t *testing.T) {
	req := httpcore.Request{
		Path:    "/echo/test",
		Headers: map[string]string{"accept-encoding": "deflate, gzip, br"},
	}

	resp := Echo(req)

	if resp.Encoding != "gzip" {
		t.Errorf("encoding = %q, want %q", resp.Encoding, "gzip")
	}
}

func TestEcho_WithUnsupportedEncoding(t *testing.T) {
	req := httpcore.Request{
		Path:    "/echo/test",
		Headers: map[string]string{"accept-encoding": "deflate, br"},
	}

	resp := Echo(req)

	if resp.Encoding != "" {
		t.Errorf("encoding = %q, want empty for unsupported encoding", resp.Encoding)
	}
	if string(resp.Body) != "test" {
		t.Errorf("body = %q, want %q", resp.Body, "test")
	}
}

func TestEcho_NoAcceptEncoding(t *testing.T) {
	req := httpcore.Request{
		Path:    "/echo/world",
		Headers: map[string]string{},
	}

	resp := Echo(req)

	if resp.Encoding != "" {
		t.Errorf("encoding = %q, want empty", resp.Encoding)
	}
	if string(resp.Body) != "world" {
		t.Errorf("body = %q, want %q", resp.Body, "world")
	}
}
