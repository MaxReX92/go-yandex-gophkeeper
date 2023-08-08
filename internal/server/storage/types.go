package storage

import (
	"context"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/generated"
)

type SecretsStorage interface {
	AddSecret(ctx context.Context, userID string, secret *generated.Secret) error
	ChangeSecret(ctx context.Context, userID string, secret *generated.Secret) error
	GetAllSecrets(ctx context.Context, userID string) ([]*generated.Secret, error)
	RemoveSecret(ctx context.Context, userID string, secret *generated.Secret) error
}
