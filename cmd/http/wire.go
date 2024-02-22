//go:build wireinject
// +build wireinject

package main

import (
	"github.com/elct9620/gopherday2024/internal/app/rest"
	v1 "github.com/elct9620/gopherday2024/internal/app/rest/v1"
	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
)

func Initialize() (*chi.Mux, error) {
	wire.Build(
		v1.RouteSet,
		ProvideRouters,
		rest.New,
	)
	return nil, nil
}

func ProvideRouters(
	v1Api *v1.Router,
) []rest.Router {
	return []rest.Router{
		rest.NewProbeRouter(),
		v1Api,
	}
}
