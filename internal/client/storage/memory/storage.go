package memory

import (
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
)

type memoryStorage struct {
}

func NewStorage() *memoryStorage {
	return &memoryStorage{}
}

func (m *memoryStorage) AddSecret(secret model.Secret) error {
	//TODO implement me
	panic("implement me")
}

func (m *memoryStorage) ChangeSecret(secret model.Secret) error {
	//TODO implement me
	panic("implement me")
}

func (m *memoryStorage) GetSecretById(secretType model.SecretType, identity string) (model.Secret, error) {
	//TODO implement me
	panic("implement me")
}

func (m *memoryStorage) GetAllSecrets(secretType model.SecretType) ([]model.Secret, error) {
	//TODO implement me
	panic("implement me")
}

func (m *memoryStorage) RemoveSecret(secret model.Secret) error {
	//TODO implement me
	panic("implement me")
}
