// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetTransaction(ctx context.Context, id int64) (Transaction, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListTransactionsForUser(ctx context.Context, username string) ([]Transaction, error)
}

var _ Querier = (*Queries)(nil)
