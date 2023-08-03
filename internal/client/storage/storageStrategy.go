package storage

import (
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

func (s *storageStrategy) AddSecret(secret model.Secret) error {
	err := s.memoryStorage.AddSecret(secret)
	if err != nil {
		return logger.WrapError("add memory secret", err)
	}

	err = s.remoteStorage.AddSecret(secret)
	if err != nil {
		return logger.WrapError("add remote secret", err)
	}

	return nil
}

func (s *storageStrategy) ChangeSecret(secret model.Secret) error {
	err := s.memoryStorage.ChangeSecret(secret)
	if err != nil {
		return logger.WrapError("change memory secret", err)
	}

	err = s.remoteStorage.ChangeSecret(secret)
	if err != nil {
		return logger.WrapError("change remote secret", err)
	}

	return nil
}

func (s *storageStrategy) GetSecretById(secretType model.SecretType, identity string) (model.Secret, error) {
	return s.memoryStorage.GetSecretById(secretType, identity)
}

func (s *storageStrategy) GetAllSecrets(secretType model.SecretType) ([]model.Secret, error) {
	return s.memoryStorage.GetAllSecrets(secretType)
}

func (s *storageStrategy) RemoveSecret(secret model.Secret) error {
	err := s.memoryStorage.RemoveSecret(secret)
	if err != nil {
		return logger.WrapError("remove memory secret", err)
	}

	err = s.remoteStorage.RemoveSecret(secret)
	if err != nil {
		return logger.WrapError("remove remote secret", err)
	}

	return nil
}
