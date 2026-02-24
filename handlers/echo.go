package handlers

import (
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/compress"
	"github.com/codecrafters-io/http-server-starter-go/httpcore"
)

func Echo(req httpcore.Request) httpcore.Response {
	body := strings.TrimPrefix(req.Path, "/echo/")
	encoding := req.Headers["Accept-Encoding"]

	if strings.Contains(encoding, "gzip") {
		compressed, err := compression.Compress([]byte(body))
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
