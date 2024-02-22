package rest

import "net/http"

type Route interface {
	Method() string
	Path() string
	http.Handler
}
