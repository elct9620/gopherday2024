package app

import (
	"net/http"

	"github.com/elct9620/gopherday2024/internal/app/rest"
	v1 "github.com/elct9620/gopherday2024/internal/app/rest/v1"
	v2 "github.com/elct9620/gopherday2024/internal/app/rest/v2"
	v3 "github.com/elct9620/gopherday2024/internal/app/rest/v3"
	"github.com/elct9620/gopherday2024/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
)

var RestServerSet = wire.NewSet(
	v1.RouteSet,
	v2.RouteSet,
	v3.RouteSet,
	ProvideRestRouters,
	rest.DefaultSet,
	NewRestServerConfig,
	NewRestServer,
)

var RestTestSet = wire.NewSet(
	v1.RouteSet,
	v2.RouteSet,
	v3.RouteSet,
	ProvideRestRouters,
	rest.TestSet,
)

func ProvideRestRouters(
	v1Api *v1.Router,
	v2Api *v2.Router,
	v3Api *v3.Router,
) []rest.Router {
	return []rest.Router{
		rest.NewProbeRouter(),
		v1Api,
		v2Api,
		v3Api,
	}
}

type RestServerConfig struct {
	Address string
}

func NewRestServerConfig(config *config.Config) *RestServerConfig {
	return &RestServerConfig{
		Address: config.HttpAddr,
	}
}

type RestServer struct {
	chi.Router
	config *RestServerConfig
}

func NewRestServer(router *chi.Mux, config *RestServerConfig) *RestServer {
	return &RestServer{Router: router, config: config}
}

func (s *RestServer) Serve() error {
	return http.ListenAndServe(s.config.Address, s.Router)
}
