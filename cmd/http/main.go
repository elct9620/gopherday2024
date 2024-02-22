package main

import (
	"net/http"

	"github.com/elct9620/gopherday2024/internal/app/rest"
)

func main() {
	app := rest.New(
		rest.NewProbeRouter(),
	)

	http.ListenAndServe(":8080", app)
}
