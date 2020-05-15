package service

import (
	"context"
	"errors"

	"github.com/AlpacaLabs/api-auth/internal/db"
	authV1 "github.com/AlpacaLabs/protorepo-auth-go/alpacalabs/auth/v1"
)

var (
	ErrCannotRevokeOtherSession = errors.New("cannot revoke someone else's session")
)

func (s Service) RevokeSession(ctx context.Context, request *authV1.RevokeSessionRequest) (*authV1.RevokeSessionResponse, error) {
	requesterID := getRequesterID(ctx)

	sessionID := request.SessionId

	err := s.dbClient.RunInTransaction(ctx, func(ctx context.Context, tx db.Transaction) error {
		s, err := tx.GetSession(ctx, sessionID)
		if err != nil {
			return err
		}

		// TODO check if requester is admin
		if s.AccountId != requesterID {
			return ErrCannotRevokeOtherSession
		}

		return tx.RevokeSession(ctx, sessionID)
	})

	if err != nil {
		return nil, err
	}

	return &authV1.RevokeSessionResponse{}, nil
}

func (s Service) RevokeSessionsForAccount(ctx context.Context, request *authV1.RevokeSessionsForAccountRequest) (*authV1.RevokeSessionsForAccountResponse, error) {
	requesterID := getRequesterID(ctx)

	accountID := request.AccountId

	// TODO check if requester is admin
	if accountID != requesterID {
		return nil, ErrCannotRevokeOtherSession
	}

	err := s.dbClient.RunInTransaction(ctx, func(ctx context.Context, tx db.Transaction) error {
		return tx.RevokeSessionsForAccount(ctx, accountID)
	})

	if err != nil {
		return nil, err
	}

	return &authV1.RevokeSessionsForAccountResponse{}, nil
}
