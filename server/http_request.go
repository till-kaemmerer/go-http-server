package server

import (
	"net"
	"strings"
)

type HttpRequest struct {
	Method  string
	Path    string
	Headers map[string]string
	Body    string
}

func ReadHttpRequest(conn net.Conn) *HttpRequest {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)

	if err != nil {
		return nil
	}

	request := &HttpRequest{}
	request.Headers = make(map[string]string)

	lines := strings.Split(string(buffer), "\n")

	requestLine := strings.Split(lines[0], " ")
	request.Method = requestLine[0]
	request.Path = requestLine[1]

	for i := 1; i < len(lines); i++ {
		if lines[i] == "\r" {
			break
		}

		header := strings.Split(lines[i], ": ")
		request.Headers[header[0]] = header[1]
	}

	body := strings.Split(string(buffer), "\r\n\r\n")
	if len(body) > 1 {
		request.Body = body[1]
	}

	return request
}
