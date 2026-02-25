package handlers

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/mtsakharov/go-http-server/internal/server/httpcore"
)

func Files(req httpcore.Request, dir string) httpcore.Response {
	filename := strings.TrimPrefix(req.Path, "/files/")
	fullPath := filepath.Join(dir, filename)
	if !strings.HasPrefix(fullPath, filepath.Clean(dir)) {
		return httpcore.Response{Status: httpcore.StatusNotFound}
	}

	switch req.Method {
	case httpcore.GET:
		data, err := os.ReadFile(fullPath)
		if err != nil {
			return httpcore.Response{Status: httpcore.StatusNotFound}
		}
		return httpcore.Response{
			Status:      httpcore.StatusOK,
			ContentType: "application/octet-stream",
			Body:        data,
		}

	case httpcore.POST:
		err := os.WriteFile(fullPath, []byte(req.Body), 0644)
		if err != nil {
			return httpcore.Response{Status: httpcore.StatusInternalServerError}
		}
		return httpcore.Response{Status: httpcore.StatusCreated}
	}

	return httpcore.Response{Status: httpcore.StatusNotFound}
}
