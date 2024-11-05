package server

import (
	"fmt"
	"net"
	"os"
	"strings"
)

type HttpResponse struct {
	StatusCode int
	Headers    map[string]string
	Body       string
}

const dir = "www"

func GetIndexResponse() HttpResponse {
	return GetFileResponse("/index.html")
}

func GetFileResponse(path string) HttpResponse {
	file, err := os.Open(dir + path)

	if err != nil {
		return HttpResponse{
			StatusCode: 404,
			Headers: map[string]string{
				"Content-Type": "text/plain",
			},
			Body: "Not Found",
		}
	}

	defer file.Close()

	fileInfo, _ := file.Stat()

	buffer := make([]byte, fileInfo.Size())
	file.Read(buffer)

	return HttpResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "text/" + strings.Split(fileInfo.Name(), ".")[1],
		},
		Body: string(buffer),
	}
}

func GetResponse(request HttpRequest) HttpResponse {
	if request.Path == "/" {
		return GetIndexResponse()
	}

	return GetFileResponse(request.Path)
}

var STATUS_CODES = map[int]string{
	200: "OK",
	204: "No Content",
	404: "Not Found",
}

func (res *HttpResponse) Write(conn net.Conn) {
	fmt.Fprintf(conn, "HTTP/1.1 %d %s\r\n", res.StatusCode, STATUS_CODES[res.StatusCode])

	for key, value := range res.Headers {
		fmt.Fprintf(conn, "%s: %s\r\n", key, value)
	}

	fmt.Fprintf(conn, "\r\n%s", res.Body)
}
