package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/AlpacaLabs/api-auth/internal/db/entities"
	authV1 "github.com/AlpacaLabs/protorepo-auth-go/alpacalabs/auth/v1"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	paginationV1 "github.com/AlpacaLabs/protorepo-pagination-go/alpacalabs/pagination/v1"
)

var (
	ErrNotFound = status.Error(codes.NotFound, "entity not found")
)

type Transaction interface {
	GetSession(ctx context.Context, sessionID string) (*authV1.Session, error)
	GetSessionsForAccount(ctx context.Context, accountID string) ([]*authV1.Session, error)

	CreateSession(ctx context.Context, in authV1.Session) error

	RevokeSession(ctx context.Context, sessionID string) error
	RevokeSessionsForAccount(ctx context.Context, accountID string) error
	RevokeSessionsForAccountExcept(ctx context.Context, accountID, sessionID string) error
}

type txImpl struct {
	tx pgx.Tx
}

func newTransaction(tx pgx.Tx) Transaction {
	return &txImpl{
		tx: tx,
	}
}

func (tx *txImpl) GetSession(ctx context.Context, sessionID string) (*authV1.Session, error) {
	var e entities.Session

	query := `
SELECT 
    id, secret, account_id 
  FROM session
  WHERE id=$1 
`

	row := tx.tx.QueryRow(ctx, query, sessionID)
	err := row.Scan(&e.ID, &e.Secret, &e.AccountID)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return e.ToProtobuf(), nil
}

func (tx *txImpl) GetSessionsForAccount(ctx context.Context, accountID string) ([]*authV1.Session, error) {
	query := `
SELECT 
    id, secret, account_id 
  FROM session
  WHERE account_id=$1 
`

	rows, err := tx.tx.Query(ctx, query, accountID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	sessions := []*authV1.Session{}

	for rows.Next() {
		var s entities.Session
		if err := rows.Scan(&s.ID, &s.Secret, &s.AccountID); err != nil {
			return nil, err
		}
		sessions = append(sessions, s.ToProtobuf())
	}

	return sessions, nil
}

func (tx *txImpl) CreateSession(ctx context.Context, in authV1.Session) error {
	query := `
INSERT INTO session(id, secret, account_id)
  VALUES($1, $2, $3)
`
	_, err := tx.tx.Exec(ctx, query, in.Id, in.Secret, in.AccountId)
	return err
}

func (tx *txImpl) RevokeSession(ctx context.Context, sessionID string) error {
	query := `
DELETE FROM session 
  WHERE id = $1
`

	_, err := tx.tx.Exec(ctx, query, sessionID)
	return err
}

func (tx *txImpl) RevokeSessionsForAccount(ctx context.Context, accountID string) error {
	query := `
DELETE FROM session 
  WHERE account_id = $1
`

	_, err := tx.tx.Exec(ctx, query, accountID)
	return err
}

func (tx *txImpl) RevokeSessionsForAccountExcept(ctx context.Context, accountID, sessionID string) error {
	query := `
DELETE FROM session 
  WHERE account_id = $1
  AND id != $2
`

	_, err := tx.tx.Exec(ctx, query, accountID, sessionID)
	return err
}

func buildOrderByClause(request paginationV1.CursorRequest) string {
	var arr []string
	for _, sortClause := range request.SortClauses {
		sortString := sortKeyword(sortClause.Sort)
		arr = append(arr, fmt.Sprintf("%s %s", sortClause.FieldName, sortString))
	}
	return strings.Join(arr, ", ")
}

func sortKeyword(sort paginationV1.Sort) string {
	if sort == paginationV1.Sort_SORT_DESC {
		return "DESC"
	}
	return "ASC"
}
