package grpc

import (
	"context"

	authV1 "github.com/AlpacaLabs/protorepo-auth-go/alpacalabs/auth/v1"
)

func (s Server) RevokeSession(ctx context.Context, request *authV1.RevokeSessionRequest) (*authV1.RevokeSessionResponse, error) {
	return s.service.RevokeSession(ctx, request)
}

func (s Server) RevokeSessionsForAccount(ctx context.Context, request *authV1.RevokeSessionsForAccountRequest) (*authV1.RevokeSessionsForAccountResponse, error) {
	return s.service.RevokeSessionsForAccount(ctx, request)
}
