package main

import (
	"net/http"
)

func main() {
	app, cleanup, err := Initialize()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	http.ListenAndServe(":8080", app)
}
