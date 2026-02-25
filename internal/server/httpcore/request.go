package httpcore

import (
	"fmt"
	"strings"
)

type Method string

const (
	GET  Method = "GET"
	POST Method = "POST"
)

type Request struct {
	Method  Method
	Path    string
	Headers map[string]string
	Body    string
}

func ParseRequest(raw string) (Request, error) {
	requestParts := strings.SplitN(raw, "\r\n\r\n", 2)

	lines := strings.Split(requestParts[0], "\r\n")

	parts := strings.Split(lines[0], " ")
	if len(parts) < 2 {
		return Request{}, fmt.Errorf("malformed request line: %q", lines[0])
	}

	headers := make(map[string]string)
	for _, line := range lines[1:] {
		if line == "" {
			break
		}
		kv := strings.SplitN(line, ": ", 2)
		if len(kv) == 2 {
			headers[strings.ToLower(kv[0])] = kv[1]
		}
	}

	body := ""
	if len(requestParts) == 2 {
		body = requestParts[1]
	}

	return Request{
		Method:  Method(parts[0]),
		Path:    parts[1],
		Headers: headers,
		Body:    body,
	}, nil
}
