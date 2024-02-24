//go:build wireinject
// +build wireinject

package app

import (
	"github.com/elct9620/gopherday2024/internal/app"
	"github.com/elct9620/gopherday2024/internal/repository"
	"github.com/elct9620/gopherday2024/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
)

func InitializeTest() (*chi.Mux, error) {
	wire.Build(
		repository.InMemorySet,
		usecase.DefaultSet,
		app.RestTestSet,
	)
	return nil, nil
}
