package app

import (
	"github.com/elct9620/gopherday2024/internal/app/rest"
	v1 "github.com/elct9620/gopherday2024/internal/app/rest/v1"
	"github.com/google/wire"
)

var RestSet = wire.NewSet(
	v1.RouteSet,
	ProvideRestRouters,
	rest.New,
)

func ProvideRestRouters(
	v1Api *v1.Router,
) []rest.Router {
	return []rest.Router{
		rest.NewProbeRouter(),
		v1Api,
	}
}
