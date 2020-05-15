package service

import (
	"context"
	"errors"

	"github.com/AlpacaLabs/api-auth/internal/db"
	authV1 "github.com/AlpacaLabs/protorepo-auth-go/alpacalabs/auth/v1"
)

var (
	ErrMustBeAdminToListOtherSessions = errors.New("must be an admin to list someone else's sessions")
)

// TODO we don't need 2 endpoints. if you wanna fetch your own sessions you would just call ListSessionsForAccount.
func (s Service) ListMySessions(ctx context.Context, request *authV1.ListMySessionsRequest) (*authV1.ListMySessionsResponse, error) {
	var sessions []*authV1.Session
	requesterID := getRequesterID(ctx)
	err := s.dbClient.RunInTransaction(ctx, func(ctx context.Context, tx db.Transaction) error {
		s, err := tx.GetSessionsForAccount(ctx, requesterID)
		if err != nil {
			return err
		}
		sessions = s
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &authV1.ListMySessionsResponse{
		Sessions: sessions,
	}, nil
}

func (s Service) ListSessionsForAccount(ctx context.Context, request *authV1.ListSessionsForAccountRequest) (*authV1.ListSessionsForAccountResponse, error) {
	var sessions []*authV1.Session
	requesterID := getRequesterID(ctx)

	// TODO check if requester is admin
	if requesterID != request.AccountId {
		return nil, ErrMustBeAdminToListOtherSessions
	}

	err := s.dbClient.RunInTransaction(ctx, func(ctx context.Context, tx db.Transaction) error {
		s, err := tx.GetSessionsForAccount(ctx, request.AccountId)
		if err != nil {
			return err
		}
		sessions = s
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &authV1.ListSessionsForAccountResponse{
		Sessions: sessions,
	}, nil
}
