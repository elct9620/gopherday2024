package main

import (
	"net"

	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}
