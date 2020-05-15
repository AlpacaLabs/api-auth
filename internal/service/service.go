package service

import (
	"github.com/AlpacaLabs/api-auth/internal/configuration"
	"github.com/AlpacaLabs/api-auth/internal/db"
	"google.golang.org/grpc"
)

type Service struct {
	config       configuration.Config
	dbClient     db.Client
	accountConn  *grpc.ClientConn
	mfaConn      *grpc.ClientConn
	passwordConn *grpc.ClientConn
}

func NewService(
	config configuration.Config,
	dbClient db.Client,
	accountConn *grpc.ClientConn,
	mfaConn *grpc.ClientConn,
	passwordConn *grpc.ClientConn) Service {
	return Service{
		config:       config,
		dbClient:     dbClient,
		accountConn:  accountConn,
		mfaConn:      mfaConn,
		passwordConn: passwordConn,
	}
}
