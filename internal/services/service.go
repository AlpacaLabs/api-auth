package services

import (
	"github.com/AlpacaLabs/auth/internal/config"
	"github.com/AlpacaLabs/auth/internal/db"
	"google.golang.org/grpc"
)

type Service struct {
	config         config.Config
	dbClient       db.Client
	authConn       *grpc.ClientConn
	iterationCount int
}

func NewService(config config.Config, dbClient db.Client) Service {
	return Service{
		config:   config,
		dbClient: dbClient,
		// TODO call CalibrateIterationCount
		iterationCount: 10000,
	}
}
