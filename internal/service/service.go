package service

import (
	"github.com/AlpacaLabs/auth/internal/configuration"
	"github.com/AlpacaLabs/auth/internal/db"
	"google.golang.org/grpc"
)

type Service struct {
	config         configuration.Config
	dbClient       db.Client
	authConn       *grpc.ClientConn
	iterationCount int
}

func NewService(config configuration.Config, dbClient db.Client) Service {
	return Service{
		config:   config,
		dbClient: dbClient,
		// TODO call CalibrateIterationCount
		iterationCount: 10000,
	}
}
