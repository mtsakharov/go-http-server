package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/mtsakharov/go-http-server/internal/server/handlers"
	"github.com/mtsakharov/go-http-server/internal/server/httpcore"
)

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
			continue
		}
		go handleConn(conn, dir)
	}
}

func handleConn(conn net.Conn, dir string) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		raw, err := readHTTPRequest(reader)
		if err != nil {
			return
		}

		req, err := httpcore.ParseRequest(raw)
		if err != nil {
			return
		}

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

		if req.Headers["connection"] == "close" {
			resp.Connection = "close"
		}

		if err := resp.Write(conn); err != nil {
			return
		}

		if req.Headers["connection"] == "close" {
			return
		}
	}
}

func readHTTPRequest(reader *bufio.Reader) (string, error) {
	var builder strings.Builder

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		builder.WriteString(line)
		if line == "\r\n" {
			break
		}
	}

	headers := builder.String()

	contentLength := 0
	for _, line := range strings.Split(headers, "\r\n") {
		lower := strings.ToLower(line)
		if strings.HasPrefix(lower, "content-length: ") {
			val := strings.TrimPrefix(lower, "content-length: ")
			n, err := strconv.Atoi(strings.TrimSpace(val))
			if err != nil {
				return "", fmt.Errorf("invalid Content-Length: %q", val)
			}
			contentLength = n
			break
		}
	}

	if contentLength > 0 {
		body := make([]byte, contentLength)
		_, err := io.ReadFull(reader, body)
		if err != nil {
			return "", err
		}
		builder.Write(body)
	}

	return builder.String(), nil
}
