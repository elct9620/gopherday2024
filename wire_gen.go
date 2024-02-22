// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/elct9620/gopherday2024/internal/app"
	"github.com/elct9620/gopherday2024/internal/app/rest"
	"github.com/elct9620/gopherday2024/internal/app/rest/v1"
	"github.com/go-chi/chi/v5"
)

// Injectors from wire.go:

func InitializeTest() (*chi.Mux, error) {
	v := v1.ProvideRotues()
	router := v1.New(v...)
	v2 := app.ProvideRestRouters(router)
	mux := rest.NewTest(v2...)
	return mux, nil
}
