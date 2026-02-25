package handlers

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mtsakharov/go-http-server/internal/server/httpcore"
)

func TestFiles_GETExistingFile(t *testing.T) {
	dir := t.TempDir()
	content := []byte("file content here")
	os.WriteFile(filepath.Join(dir, "test.txt"), content, 0644)

	req := httpcore.Request{
		Method:  httpcore.GET,
		Path:    "/files/test.txt",
		Headers: map[string]string{},
	}

	resp := Files(req, dir)

	if resp.Status != httpcore.StatusOK {
		t.Errorf("status = %d, want %d", resp.Status, httpcore.StatusOK)
	}
	if resp.ContentType != "application/octet-stream" {
		t.Errorf("content-type = %q, want %q", resp.ContentType, "application/octet-stream")
	}
	if string(resp.Body) != "file content here" {
		t.Errorf("body = %q, want %q", resp.Body, "file content here")
	}
}

func TestFiles_GETMissingFile(t *testing.T) {
	dir := t.TempDir()

	req := httpcore.Request{
		Method:  httpcore.GET,
		Path:    "/files/nonexistent.txt",
		Headers: map[string]string{},
	}

	resp := Files(req, dir)

	if resp.Status != httpcore.StatusNotFound {
		t.Errorf("status = %d, want %d", resp.Status, httpcore.StatusNotFound)
	}
}

func TestFiles_POSTCreatesFile(t *testing.T) {
	dir := t.TempDir()

	req := httpcore.Request{
		Method:  httpcore.POST,
		Path:    "/files/new.txt",
		Headers: map[string]string{},
		Body:    "new file content",
	}

	resp := Files(req, dir)

	if resp.Status != httpcore.StatusCreated {
		t.Errorf("status = %d, want %d", resp.Status, httpcore.StatusCreated)
	}

	// Verify file was written
	data, err := os.ReadFile(filepath.Join(dir, "new.txt"))
	if err != nil {
		t.Fatalf("failed to read created file: %v", err)
	}
	if string(data) != "new file content" {
		t.Errorf("file content = %q, want %q", data, "new file content")
	}
}

func TestFiles_PathTraversal(t *testing.T) {
	dir := t.TempDir()

	req := httpcore.Request{
		Method:  httpcore.GET,
		Path:    "/files/../../etc/passwd",
		Headers: map[string]string{},
	}

	resp := Files(req, dir)

	if resp.Status != httpcore.StatusNotFound {
		t.Errorf("status = %d, want %d (path traversal should be blocked)", resp.Status, httpcore.StatusNotFound)
	}
}

func TestFiles_UnsupportedMethod(t *testing.T) {
	dir := t.TempDir()

	req := httpcore.Request{
		Method:  "DELETE",
		Path:    "/files/test.txt",
		Headers: map[string]string{},
	}

	resp := Files(req, dir)

	if resp.Status != httpcore.StatusNotFound {
		t.Errorf("status = %d, want %d for unsupported method", resp.Status, httpcore.StatusNotFound)
	}
}

func TestFiles_POSTEmptyBody(t *testing.T) {
	dir := t.TempDir()

	req := httpcore.Request{
		Method:  httpcore.POST,
		Path:    "/files/empty.txt",
		Headers: map[string]string{},
		Body:    "",
	}

	resp := Files(req, dir)

	if resp.Status != httpcore.StatusCreated {
		t.Errorf("status = %d, want %d", resp.Status, httpcore.StatusCreated)
	}

	data, err := os.ReadFile(filepath.Join(dir, "empty.txt"))
	if err != nil {
		t.Fatalf("failed to read created file: %v", err)
	}
	if len(data) != 0 {
		t.Errorf("file should be empty, got %d bytes", len(data))
	}
}
