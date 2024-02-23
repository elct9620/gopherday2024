//go:build wireinject
// +build wireinject

package main

import (
	"github.com/elct9620/gopherday2024/internal/app"
	"github.com/elct9620/gopherday2024/internal/repository"
	"github.com/elct9620/gopherday2024/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
)

func Initialize() (*chi.Mux, error) {
	wire.Build(
		repository.DefaultSet,
		usecase.DefaultSet,
		app.RestSet,
	)
	return nil, nil
}
