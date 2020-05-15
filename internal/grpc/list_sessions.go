package grpc

import (
	"context"

	authV1 "github.com/AlpacaLabs/protorepo-auth-go/alpacalabs/auth/v1"
)

func (s Server) ListMySessions(ctx context.Context, request *authV1.ListMySessionsRequest) (*authV1.ListMySessionsResponse, error) {
	return s.service.ListMySessions(ctx, request)
}

func (s Server) ListSessionsForAccount(ctx context.Context, request *authV1.ListSessionsForAccountRequest) (*authV1.ListSessionsForAccountResponse, error) {
	return s.service.ListSessionsForAccount(ctx, request)
}
