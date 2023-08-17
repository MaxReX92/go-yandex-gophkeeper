package storage

import (
	"context"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
)

// ClientSecretsStorage is a main storage.
type ClientSecretsStorage interface {
	// AddSecret add new secret to secret storage.
	AddSecret(ctx context.Context, secret model.Secret) error
	// ChangeSecret edit an existing secret in secret storage.
	ChangeSecret(ctx context.Context, secret model.Secret) error
	// GetSecretByID returns secret by identity.
	GetSecretByID(ctx context.Context, secretType model.SecretType, identity string) (model.Secret, error)
	// GetAllSecrets return all stored secrets by secret type.
	GetAllSecrets(ctx context.Context, secretType model.SecretType) ([]model.Secret, error)
	// RemoveSecret remove an existing secret from secret storage.
	RemoveSecret(ctx context.Context, secret model.Secret) error
}
