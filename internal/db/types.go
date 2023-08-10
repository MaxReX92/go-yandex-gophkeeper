package db

import (
	"context"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/generated"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
)

type Service interface {
	CallInTransaction(context.Context, func(context.Context, Executor) error) error
	CallInTransactionResult(ctx context.Context, action func(context.Context, Executor) ([]any, error)) ([]any, error)
}

type Executor interface {
	AddUser(ctx context.Context, id string, username string, password string, personalToken string) error
	GetUserByUserName(ctx context.Context, username string) (*model.User, error)

	AddSecret(ctx context.Context, userID string, secret *generated.Secret) error
	ChangeSecret(ctx context.Context, userID string, secret *generated.Secret) error
	GetAllSecrets(ctx context.Context, userID string) ([]*generated.Secret, error)
	RemoveSecret(ctx context.Context, userID string, secret *generated.Secret) error
}
