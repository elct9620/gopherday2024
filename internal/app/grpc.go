package app

import (
	"net"

	rpc "github.com/elct9620/gopherday2024/internal/app/grpc"
	"github.com/elct9620/gopherday2024/pkg/events"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var GrpcSet = wire.NewSet(
	rpc.NewEventsServer,
	ProvideGrpcServer,
	NewGrpcServer,
)

func ProvideGrpcServer(eventServer *rpc.EventsServer) *grpc.Server {
	server := grpc.NewServer()

	events.RegisterEventsServer(server, eventServer)
	reflection.Register(server)

	return server
}

type GrpcServer struct {
	*grpc.Server
}

func NewGrpcServer(server *grpc.Server) *GrpcServer {
	return &GrpcServer{Server: server}
}

func (s *GrpcServer) Serve() error {
	socket, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	return s.Server.Serve(socket)
}
