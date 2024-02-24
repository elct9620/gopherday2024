package main

func main() {
	server, cleanup, err := Initialize()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := server.Serve(); err != nil {
		panic(err)
	}
}
