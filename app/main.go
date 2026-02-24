package main

import (
	_ "bytes"
	_ "compress/gzip"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/handlers"
	"github.com/codecrafters-io/http-server-starter-go/httpcore"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	var dir string
	for i, arg := range os.Args {
		if arg == "--directory" && i+1 < len(os.Args) {
			dir = os.Args[i+1]
		}
	}

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleConn(conn, dir)
	}
}

func handleConn(conn net.Conn, dir string) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		read, err := conn.Read(buf)
		if err != nil {
			return
		}

		req := httpcore.ParseRequest(string(buf[:read]))

		var resp httpcore.Response
		switch {
		case req.Path == "/":
			resp = httpcore.Response{Status: httpcore.StatusOK}
		case strings.HasPrefix(req.Path, "/echo/"):
			resp = handlers.Echo(req)
		case strings.HasPrefix(req.Path, "/files/"):
			resp = handlers.Files(req, dir)
		case req.Path == "/user-agent":
			resp = handlers.UserAgent(req)
		default:
			resp = httpcore.Response{Status: httpcore.StatusNotFound}
		}

		if req.Headers["Connection"] == "close" {
			resp.Connection = "close"
		}

		resp.Write(conn)

		if req.Headers["Connection"] == "close" {
			return
		}
	}
}
}
