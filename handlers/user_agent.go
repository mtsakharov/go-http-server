package handlers

import "github.com/codecrafters-io/http-server-starter-go/httpcore"

func UserAgent(req httpcore.Request) httpcore.Response {
	return httpcore.Response{
		Status:      httpcore.StatusOK,
		ContentType: "text/plain",
		Body:        []byte(req.Headers["User-Agent"]),
	}
}
