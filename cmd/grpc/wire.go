//go:build wireinject
// +build wireinject

package main

import (
	"github.com/elct9620/gopherday2024/internal/app"
	"github.com/elct9620/gopherday2024/internal/repository"
	"github.com/elct9620/gopherday2024/internal/usecase"
	"github.com/google/wire"
)

func Initialize() (*app.GrpcServer, func(), error) {
	wire.Build(
		repository.DefaultSet,
		usecase.DefaultSet,
		app.DefaultSet,
		app.GrpcSet,
	)
	return nil, nil, nil
}
