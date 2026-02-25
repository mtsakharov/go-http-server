package httpcore

import (
	"net"
	"strings"
	"testing"
)

func TestResponseWrite_AllFields(t *testing.T) {
	server, client := net.Pipe()
	defer server.Close()
	defer client.Close()

	resp := Response{
		Status:      StatusOK,
		ContentType: "text/plain",
		Encoding:    "gzip",
		Connection:  "close",
		Body:        []byte("hello"),
	}

	go func() {
		resp.Write(server)
		server.Close()
	}()

	buf := make([]byte, 4096)
	n, _ := client.Read(buf)
	result := string(buf[:n])

	if !strings.HasPrefix(result, "HTTP/1.1 200 OK\r\n") {
		t.Errorf("response should start with status line, got: %q", result)
	}
	if !strings.Contains(result, "Content-Type: text/plain\r\n") {
		t.Error("missing Content-Type header")
	}
	if !strings.Contains(result, "Content-Encoding: gzip\r\n") {
		t.Error("missing Content-Encoding header")
	}
	if !strings.Contains(result, "Connection: close\r\n") {
		t.Error("missing Connection header")
	}
	if !strings.Contains(result, "Content-Length: 5\r\n") {
		t.Error("missing or incorrect Content-Length header")
	}
	if !strings.HasSuffix(result, "\r\n\r\nhello") {
		t.Errorf("response should end with blank line + body, got: %q", result)
	}
}

func TestResponseWrite_StatusOnly(t *testing.T) {
	server, client := net.Pipe()
	defer server.Close()
	defer client.Close()

	resp := Response{Status: StatusNotFound}

	go func() {
		resp.Write(server)
		server.Close()
	}()

	buf := make([]byte, 4096)
	n, _ := client.Read(buf)
	result := string(buf[:n])

	if !strings.HasPrefix(result, "HTTP/1.1 404 Not Found\r\n") {
		t.Errorf("unexpected status line: %q", result)
	}
	if !strings.Contains(result, "Content-Length: 0\r\n") {
		t.Error("Content-Length: 0 should always be present for empty body")
	}
	// No optional headers
	if strings.Contains(result, "Content-Type:") {
		t.Error("Content-Type should be omitted when empty")
	}
	if strings.Contains(result, "Content-Encoding:") {
		t.Error("Content-Encoding should be omitted when empty")
	}
	if strings.Contains(result, "Connection:") {
		t.Error("Connection should be omitted when empty")
	}
}

func TestResponseWrite_EmptyBody_HasContentLengthZero(t *testing.T) {
	server, client := net.Pipe()
	defer server.Close()
	defer client.Close()

	resp := Response{
		Status:      StatusOK,
		ContentType: "text/plain",
	}

	go func() {
		resp.Write(server)
		server.Close()
	}()

	buf := make([]byte, 4096)
	n, _ := client.Read(buf)
	result := string(buf[:n])

	if !strings.Contains(result, "Content-Length: 0\r\n") {
		t.Error("Content-Length: 0 should be present even with empty body")
	}
}

func TestResponseWrite_CreatedStatus(t *testing.T) {
	server, client := net.Pipe()
	defer server.Close()
	defer client.Close()

	resp := Response{Status: StatusCreated}

	go func() {
		resp.Write(server)
		server.Close()
	}()

	buf := make([]byte, 4096)
	n, _ := client.Read(buf)
	result := string(buf[:n])

	if !strings.HasPrefix(result, "HTTP/1.1 201 Created\r\n") {
		t.Errorf("unexpected status line: %q", result)
	}
}
