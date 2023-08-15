package remote

import (
	"context"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/secret"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

type remoteStorage struct {
	service secret.Service
}

// NewStorage creates a new instance of remote secrets storage.
func NewStorage(service secret.Service) *remoteStorage {
	return &remoteStorage{
		service: service,
	}
}

func (r *remoteStorage) AddSecret(ctx context.Context, secret model.Secret) error {
	err := r.service.AddSecret(ctx, secret)
	if err != nil {
		return logger.WrapError("send secret service add request", err)
	}

	return nil
}

func (r *remoteStorage) ChangeSecret(ctx context.Context, secret model.Secret) error {
	err := r.service.ChangeSecret(ctx, secret)
	if err != nil {
		return logger.WrapError("send secret service change request", err)
	}

	return nil
}

func (r *remoteStorage) GetSecretByID(ctx context.Context, secretType model.SecretType, identity string) (model.Secret, error) {
	// TODO implement me
	panic("implement me")
}

func (r *remoteStorage) GetAllSecrets(ctx context.Context, secretType model.SecretType) ([]model.Secret, error) {
	// TODO implement me
	panic("implement me")
}

func (r *remoteStorage) RemoveSecret(ctx context.Context, secret model.Secret) error {
	err := r.service.RemoveSecret(ctx, secret)
	if err != nil {
		return logger.WrapError("send secret service change request", err)
	}

	return nil
}
