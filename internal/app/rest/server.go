package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router interface {
	Namespace() string
	chi.Router
}

func New(routers ...Router) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	for _, router := range routers {
		r.Mount(router.Namespace(), router)
	}

	return r
}
