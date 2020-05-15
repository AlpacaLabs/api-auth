package grpc

import (
	"fmt"
	"net"

	"github.com/AlpacaLabs/api-auth/internal/configuration"
	"github.com/AlpacaLabs/api-auth/internal/service"

	authV1 "github.com/AlpacaLabs/protorepo-auth-go/alpacalabs/auth/v1"
	health "google.golang.org/grpc/health/grpc_health_v1"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	config  configuration.Config
	service service.Service
}

func NewServer(config configuration.Config, service service.Service) Server {
	return Server{
		config:  config,
		service: service,
	}
}

func (s Server) Run() {
	address := fmt.Sprintf(":%d", s.config.GrpcPort)

	log.Infof("Preparing to serve gRPC on %s", address)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	authV1.RegisterListSessionsServiceServer(grpcServer, s)
	authV1.RegisterLoginServiceServer(grpcServer, s)
	authV1.RegisterRevokeServiceServer(grpcServer, s)
	health.RegisterHealthServer(grpcServer, s)

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	log.Infof("Serving gRPC on %s", address)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
