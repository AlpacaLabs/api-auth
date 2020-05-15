package grpc

import (
	"context"

	authV1 "github.com/AlpacaLabs/protorepo-auth-go/alpacalabs/auth/v1"
)

func (s Server) Login(ctx context.Context, request *authV1.LoginRequest) (*authV1.LoginResponse, error) {
	return s.service.Login(ctx, request)
}
