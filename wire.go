//go:build wireinject
// +build wireinject

package app

import (
	"github.com/elct9620/gopherday2024/internal/app"
	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
)

func InitializeTest() (*chi.Mux, error) {
	wire.Build(
		app.RestTestSet,
	)
	return nil, nil
}
