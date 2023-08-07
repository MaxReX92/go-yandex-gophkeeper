package storage

import (
	"context"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

type storageStrategy struct {
	memoryStorage ClientSecretsStorage
	remoteStorage ClientSecretsStorage
}

func NewStorageStrategy(memoryStorage ClientSecretsStorage, remoteStorage ClientSecretsStorage) *storageStrategy {
	return &storageStrategy{
		memoryStorage: memoryStorage,
		remoteStorage: remoteStorage,
	}
}

func (s *storageStrategy) AddSecret(ctx context.Context, secret model.Secret) error {
	err := s.remoteStorage.AddSecret(ctx, secret)
	if err != nil {
		return logger.WrapError("add remote secret", err)
	}

	return nil
}

func (s *storageStrategy) ChangeSecret(ctx context.Context, secret model.Secret) error {
	err := s.remoteStorage.ChangeSecret(ctx, secret)
	if err != nil {
		return logger.WrapError("change remote secret", err)
	}

	return nil
}

func (s *storageStrategy) GetSecretById(ctx context.Context, secretType model.SecretType, identity string) (model.Secret, error) {
	return s.memoryStorage.GetSecretById(ctx, secretType, identity)
}

func (s *storageStrategy) GetAllSecrets(ctx context.Context, secretType model.SecretType) ([]model.Secret, error) {
	return s.memoryStorage.GetAllSecrets(ctx, secretType)
}

func (s *storageStrategy) RemoveSecret(ctx context.Context, secret model.Secret) error {
	err := s.remoteStorage.RemoveSecret(ctx, secret)
	if err != nil {
		return logger.WrapError("remove remote secret", err)
	}

	return nil
}
