package httpcore

import (
	"fmt"
	"net"
)

type Response struct {
	Status      StatusCode
	ContentType string
	Encoding    string
	Body        []byte
}

func (r Response) Write(conn net.Conn) {
	text := statusText[r.Status]
	result := fmt.Sprintf("HTTP/1.1 %d %s\r\n", r.Status, text)

	if r.ContentType != "" {
		result += fmt.Sprintf("Content-Type: %s\r\n", r.ContentType)
	}
	if r.Encoding != "" {
		result += fmt.Sprintf("Content-Encoding: %s\r\n", r.Encoding)
	}
	if len(r.Body) > 0 {
		result += fmt.Sprintf("Content-Length: %d\r\n", len(r.Body))
	}

	result += "\r\n"

	conn.Write([]byte(result))
	if len(r.Body) > 0 {
		conn.Write(r.Body)
	}
}
