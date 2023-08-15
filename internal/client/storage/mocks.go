package storage

import (
	"context"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
	"github.com/stretchr/testify/mock"
)

type ClientSecretsStorageMock struct {
	mock.Mock
}

func (c *ClientSecretsStorageMock) AddSecret(ctx context.Context, secret model.Secret) error {
	args := c.Called(ctx, secret)
	return args.Error(0)
}

func (c *ClientSecretsStorageMock) ChangeSecret(ctx context.Context, secret model.Secret) error {
	args := c.Called(ctx, secret)
	return args.Error(0)
}

func (c *ClientSecretsStorageMock) GetSecretByID(ctx context.Context, secretType model.SecretType, identity string) (model.Secret, error) {
	args := c.Called(ctx, secretType, identity)
	return args.Get(0).(model.Secret), args.Error(1)
}

func (c *ClientSecretsStorageMock) GetAllSecrets(ctx context.Context, secretType model.SecretType) ([]model.Secret, error) {
	args := c.Called(ctx, secretType)
	return args.Get(0).([]model.Secret), args.Error(1)
}

func (c *ClientSecretsStorageMock) RemoveSecret(ctx context.Context, secret model.Secret) error {
	args := c.Called(ctx, secret)
	return args.Error(0)
}
