package server

import (
	"fmt"
	"net"
)

type Server struct {
	Host string
	Port int
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New connection from", conn.RemoteAddr())

	request := ReadHttpRequest(conn)
	if request == nil {
		fmt.Println("Error reading request")
		return
	}

	fmt.Println("Request:", request.Method, request.Path)

	response := GetResponse(*request)

	fmt.Println("Response:", response.StatusCode, response.Body)

	response.Write(conn)
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Host, s.Port))

	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	fmt.Println("Server started on", s.Host, "port", s.Port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return
		}
		go handleConnection(conn)
	}
}
