package handlers

import (
	"strings"

	"github.com/mtsakharov/go-http-server/internal/server/compress"
	"github.com/mtsakharov/go-http-server/internal/server/httpcore"
)

func Echo(req httpcore.Request) httpcore.Response {
	body := strings.TrimPrefix(req.Path, "/echo/")
	encoding := req.Headers["accept-encoding"]

	if strings.Contains(encoding, "gzip") {
		compressed, err := compress.Compress([]byte(body))
		if err != nil {
			return httpcore.Response{Status: httpcore.StatusInternalServerError}
		}
		return httpcore.Response{
			Status:      httpcore.StatusOK,
			ContentType: "text/plain",
			Encoding:    "gzip",
			Body:        compressed,
		}
	}
	return httpcore.Response{
		Status:      httpcore.StatusOK,
		ContentType: "text/plain",
		Body:        []byte(body),
	}
}
