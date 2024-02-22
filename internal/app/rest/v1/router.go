package v1

import (
	"github.com/elct9620/gopherday2024/internal/app/rest"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	*chi.Mux
}

func New(routes ...rest.Route) *Router {
	r := chi.NewRouter()

	for _, route := range routes {
		r.Method(route.Method(), route.Path(), route)
	}

	return &Router{r}
}

func (r *Router) Namespace() string {
	return "/v1"
}
