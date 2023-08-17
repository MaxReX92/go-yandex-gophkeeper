package secret

import (
	"context"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
)

// Service implements logic for working with secrets.
type Service interface {
	// AddSecret add new secret to secret storage.
	AddSecret(ctx context.Context, secret model.Secret) error
	// ChangeSecret edit an existing secret in secret storage.
	ChangeSecret(ctx context.Context, secret model.Secret) error
	// RemoveSecret remove an existing secret from secret storage.
	RemoveSecret(ctx context.Context, secret model.Secret) error
	// SecretEvents provides secret events stream channel.
	SecretEvents(ctx context.Context) <-chan *model.SecretEvent
}
