package main

import (
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}

	server, cleanup, err := Initialize()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}
