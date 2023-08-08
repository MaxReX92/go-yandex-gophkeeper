package storage

import (
	"context"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/generated"
)

type SecretsStorage interface {
	AddSecret(ctx context.Context, userId string, secret *generated.Secret) error
	ChangeSecret(ctx context.Context, userId string, secret *generated.Secret) error
	GetAllSecrets(ctx context.Context, userId string) ([]*generated.Secret, error)
	RemoveSecret(ctx context.Context, userId string, secret *generated.Secret) error
}
