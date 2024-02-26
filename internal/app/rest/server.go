package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/wire"
)

var DefaultSet = wire.NewSet(
	OpenApiSet,
	New,
)

var TestSet = wire.NewSet(
	OpenApiSet,
	NewTest,
)

type Router interface {
	Namespace() string
	chi.Router
}

func New(routers []Router, openApiMiddleware OpenApiMiddleware) *chi.Mux {
	r := chi.NewRouter()

	r.Use(openApiMiddleware)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	for _, router := range routers {
		r.Mount(router.Namespace(), router)
	}

	return r
}

func NewTest(routers []Router, openApiMiddleware OpenApiMiddleware) *chi.Mux {
	r := chi.NewRouter()

	r.Use(openApiMiddleware)
	r.Use(middleware.Recoverer)

	for _, router := range routers {
		r.Mount(router.Namespace(), router)
	}

	return r
}
