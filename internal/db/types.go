package db

import (
	"context"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/generated"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
)

// Service provides methods for call db within a transaction.
type Service interface {
	// CallInTransaction call db query within a transaction.
	CallInTransaction(context.Context, func(context.Context, Executor) error) error
	// CallInTransactionResult call db query within a transaction with some result set.
	CallInTransactionResult(ctx context.Context, action func(context.Context, Executor) ([]any, error)) ([]any, error)
}

type Executor interface {
	// AddUser add new user to db storage.
	AddUser(ctx context.Context, id string, username string, password string, personalToken string) error
	// GetUserByUserName returns user by username.
	GetUserByUserName(ctx context.Context, username string) (*model.User, error)

	// AddSecret add new secret to db storage.
	AddSecret(ctx context.Context, userID string, secret *generated.Secret) error
	// ChangeSecret edit an existing secret in db storage.
	ChangeSecret(ctx context.Context, userID string, secret *generated.Secret) error
	// GetAllSecrets return all stored secrets by secret type.
	GetAllSecrets(ctx context.Context, userID string) ([]*generated.Secret, error)
	// RemoveSecret remove an existing secret from db storage.
	RemoveSecret(ctx context.Context, userID string, secret *generated.Secret) error
}
