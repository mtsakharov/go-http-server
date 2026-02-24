package handlers

import (
	"net"
	"os"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/httpcore"
)

func Files(req httpcore.Request, conn net.Conn, dir string) {
	filename := strings.TrimPrefix(req.Path, "/files/")
	fullPath := dir + filename

	switch req.Method {
	case httpcore.GET:
		data, err := os.ReadFile(fullPath)
		if err != nil {
			httpcore.Response{Status: httpcore.StatusNotFound}.Write(conn)
			return
		}
		httpcore.Response{
			Status:      httpcore.StatusOK,
			ContentType: "application/octet-stream",
			Body:        data,
		}.Write(conn)

	case httpcore.POST:
		err := os.WriteFile(fullPath, []byte(req.Body), 0644)
		if err != nil {
			httpcore.Response{Status: httpcore.StatusInternalServerError}.Write(conn)
			return
		}
		httpcore.Response{Status: httpcore.StatusCreated}.Write(conn)
	}
}
