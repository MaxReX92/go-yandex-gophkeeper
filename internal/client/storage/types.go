package storage

import (
	"context"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
)

type ClientSecretsStorage interface {
	AddSecret(ctx context.Context, secret model.Secret) error
	ChangeSecret(ctx context.Context, secret model.Secret) error
	GetSecretById(ctx context.Context, secretType model.SecretType, identity string) (model.Secret, error)
	GetAllSecrets(ctx context.Context, secretType model.SecretType) ([]model.Secret, error)
	RemoveSecret(ctx context.Context, secret model.Secret) error
}
