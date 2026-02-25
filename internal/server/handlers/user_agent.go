package handlers

import "github.com/mtsakharov/go-http-server/internal/server/httpcore"

func UserAgent(req httpcore.Request) httpcore.Response {
	return httpcore.Response{
		Status:      httpcore.StatusOK,
		ContentType: "text/plain",
		Body:        []byte(req.Headers["user-agent"]),
	}
}
