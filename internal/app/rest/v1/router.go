package v1

import (
	"github.com/elct9620/gopherday2024/internal/app/rest"
	"github.com/elct9620/gopherday2024/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
)

type Route rest.Route

var RouteSet = wire.NewSet(
	ProvideRotues,
	New,
)

var _ rest.Router = &Router{}

type Router struct {
	*chi.Mux
}

func New(routes ...Route) *Router {
	r := chi.NewRouter()

	for _, route := range routes {
		r.Method(route.Method(), route.Path(), route)
	}

	return &Router{r}
}

func (r *Router) Namespace() string {
	return "/v1"
}

func ProvideRotues(
	eventQuery *usecase.EventQuery,
) []Route {
	return []Route{
		NewGetEvents(eventQuery),
	}
}
