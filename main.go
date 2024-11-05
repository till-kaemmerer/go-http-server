package main

import (
	"fmt"
	"github.com/till-kaemmerer/go-http-server/server"
)

func main() {
	fmt.Println("Hello, World!")
	server := &server.Server{Host: "localhost", Port: 8080}
	server.Start()
}
