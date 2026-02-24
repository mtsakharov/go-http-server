package handlers

import (
	"net"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/compress"
	"github.com/codecrafters-io/http-server-starter-go/httpcore"
)

func Echo(req httpcore.Request, conn net.Conn) {
	body := strings.TrimPrefix(req.Path, "/echo/")
	encoding := req.Headers["Accept-Encoding"]

	if strings.Contains(encoding, "gzip") {
		compressed, err := compression.Compress([]byte(body))
		if err != nil {
			httpcore.Response{Status: httpcore.StatusInternalServerError}.Write(conn)
			return
		}
		httpcore.Response{
			Status:      httpcore.StatusOK,
			ContentType: "text/plain",
			Encoding:    "gzip",
			Body:        compressed,
		}.Write(conn)
	} else {
		httpcore.Response{
			Status:      httpcore.StatusOK,
			ContentType: "text/plain",
			Body:        []byte(body),
		}.Write(conn)
	}
}
