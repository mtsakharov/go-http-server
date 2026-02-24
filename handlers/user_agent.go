package handlers

import (
	"net"

	"github.com/codecrafters-io/http-server-starter-go/httpcore"
)

func UserAgent(req httpcore.Request, conn net.Conn) {
	body := req.Headers["User-Agent"]

	httpcore.Response{
		Status:      httpcore.StatusOK,
		ContentType: "text/plain",
		Body:        []byte(body),
	}.Write(conn)
}
