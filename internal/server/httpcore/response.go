package httpcore

import (
	"fmt"
	"net"
)

type Response struct {
	Status      StatusCode
	ContentType string
	Encoding    string
	Connection  string
	Body        []byte
}

func (r Response) Write(conn net.Conn) error {
	text := statusText[r.Status]
	result := fmt.Sprintf("HTTP/1.1 %d %s\r\n", r.Status, text)

	if r.ContentType != "" {
		result += fmt.Sprintf("Content-Type: %s\r\n", r.ContentType)
	}
	if r.Encoding != "" {
		result += fmt.Sprintf("Content-Encoding: %s\r\n", r.Encoding)
	}
	if r.Connection != "" {
		result += fmt.Sprintf("Connection: %s\r\n", r.Connection)
	}

	result += fmt.Sprintf("Content-Length: %d\r\n", len(r.Body))
	result += "\r\n"

	buf := make([]byte, 0, len(result)+len(r.Body))
	buf = append(buf, result...)
	buf = append(buf, r.Body...)

	_, err := conn.Write(buf)
	return err
}
