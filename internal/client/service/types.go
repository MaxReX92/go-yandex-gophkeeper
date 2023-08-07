package service

import (
	"context"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
)

type SecretService interface {
	AddSecret(ctx context.Context, secret model.Secret) error
	ChangeSecret(ctx context.Context, secret model.Secret) error
	RemoveSecret(ctx context.Context, secret model.Secret) error

	SecretEvents(ctx context.Context) <-chan *model.SecretEvent
}
