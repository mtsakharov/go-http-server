package compress

import (
	"bytes"
	"compress/gzip"
	"io"
	"strings"
	"testing"
)

func TestCompress_RoundTrip(t *testing.T) {
	original := []byte("Hello, World!")

	compressed, err := Compress(original)
	if err != nil {
		t.Fatalf("Compress error: %v", err)
	}

	reader, err := gzip.NewReader(bytes.NewReader(compressed))
	if err != nil {
		t.Fatalf("gzip.NewReader error: %v", err)
	}
	defer reader.Close()

	decompressed, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("ReadAll error: %v", err)
	}

	if !bytes.Equal(decompressed, original) {
		t.Errorf("roundtrip mismatch: got %q, want %q", decompressed, original)
	}
}

func TestCompress_EmptyInput(t *testing.T) {
	compressed, err := Compress([]byte{})
	if err != nil {
		t.Fatalf("Compress error: %v", err)
	}

	// Should produce valid gzip that decompresses to empty
	reader, err := gzip.NewReader(bytes.NewReader(compressed))
	if err != nil {
		t.Fatalf("gzip.NewReader error: %v", err)
	}
	defer reader.Close()

	decompressed, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("ReadAll error: %v", err)
	}

	if len(decompressed) != 0 {
		t.Errorf("expected empty output, got %d bytes", len(decompressed))
	}
}

func TestCompress_LargeInput(t *testing.T) {
	original := []byte(strings.Repeat("abcdefghij", 10000))

	compressed, err := Compress(original)
	if err != nil {
		t.Fatalf("Compress error: %v", err)
	}

	// Compressed should be smaller than original for repetitive data
	if len(compressed) >= len(original) {
		t.Errorf("compressed size (%d) should be smaller than original (%d)", len(compressed), len(original))
	}

	reader, err := gzip.NewReader(bytes.NewReader(compressed))
	if err != nil {
		t.Fatalf("gzip.NewReader error: %v", err)
	}
	defer reader.Close()

	decompressed, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("ReadAll error: %v", err)
	}

	if !bytes.Equal(decompressed, original) {
		t.Error("roundtrip mismatch for large input")
	}
}
