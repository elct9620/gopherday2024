package app

import (
	rpc "github.com/elct9620/gopherday2024/internal/app/grpc"
	"github.com/elct9620/gopherday2024/pkg/events"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var GrpcSet = wire.NewSet(
	rpc.NewEventsServer,
	ProvideGrpcServer,
)

func ProvideGrpcServer(eventServer *rpc.EventsServer) *grpc.Server {
	server := grpc.NewServer()

	events.RegisterEventsServer(server, eventServer)
	reflection.Register(server)

	return server
}
