package app

import (
	"github.com/elct9620/gopherday2024/internal/app/rest"
	v1 "github.com/elct9620/gopherday2024/internal/app/rest/v1"
	v2 "github.com/elct9620/gopherday2024/internal/app/rest/v2"
	"github.com/google/wire"
)

var RestSet = wire.NewSet(
	v1.RouteSet,
	v2.RouteSet,
	ProvideRestRouters,
	rest.DefaultSet,
)

var RestTestSet = wire.NewSet(
	v1.RouteSet,
	v2.RouteSet,
	ProvideRestRouters,
	rest.TestSet,
)

func ProvideRestRouters(
	v1Api *v1.Router,
	v2Api *v2.Router,
) []rest.Router {
	return []rest.Router{
		rest.NewProbeRouter(),
		v1Api,
		v2Api,
	}
}
