package main

import (
	"net/http"

	"github.com/elct9620/gopherday2024/internal/app/rest"
	v1 "github.com/elct9620/gopherday2024/internal/app/rest/v1"
)

func main() {
	getEvents := v1.NewGetEvents()

	v1 := v1.New(
		getEvents,
	)

	app := rest.New(
		rest.NewProbeRouter(),
		v1,
	)

	http.ListenAndServe(":8080", app)
}
