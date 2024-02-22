package main

import (
	"net/http"
)

func main() {
	app, err := Initialize()
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8080", app)
}
