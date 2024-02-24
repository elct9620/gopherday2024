package app

import (
	"net"

	rpc "github.com/elct9620/gopherday2024/internal/app/grpc"
	"github.com/elct9620/gopherday2024/internal/config"
	"github.com/elct9620/gopherday2024/pkg/events"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var GrpcSet = wire.NewSet(
	rpc.NewEventsServer,
	ProvideGrpcServer,
	NewGrpcServerConfig,
	NewGrpcServer,
)

func ProvideGrpcServer(eventServer *rpc.EventsServer) *grpc.Server {
	server := grpc.NewServer()

	events.RegisterEventsServer(server, eventServer)
	reflection.Register(server)

	return server
}

type GrpcServerConfig struct {
	Address string
}

func NewGrpcServerConfig(config *config.Config) *GrpcServerConfig {
	return &GrpcServerConfig{
		Address: config.GrpcAddr,
	}
}

type GrpcServer struct {
	*grpc.Server
	config *GrpcServerConfig
}

func NewGrpcServer(server *grpc.Server, config *GrpcServerConfig) *GrpcServer {
	return &GrpcServer{Server: server, config: config}
}

func (s *GrpcServer) Serve() error {
	socket, err := net.Listen("tcp", s.config.Address)
	if err != nil {
		return err
	}

	return s.Server.Serve(socket)
}
