package memory

import (
	"context"
	"sync"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/storage"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

type memoryStorage struct {
	secrets map[model.SecretType]map[string]model.Secret
	lock    sync.RWMutex
}

func NewStorage() *memoryStorage {
	return &memoryStorage{
		secrets: make(map[model.SecretType]map[string]model.Secret),
		lock:    sync.RWMutex{},
	}
}

func (m *memoryStorage) AddSecret(ctx context.Context, secret model.Secret) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	secretsByType := m.ensureSecrets(secret.GetType())
	identity := secret.GetIdentity()
	_, ok := secretsByType[identity]
	if ok {
		return logger.WrapError("add new secret", storage.ErrSecretAlreadyExist)
	}

	secretsByType[identity] = secret
	return nil
}

func (m *memoryStorage) ChangeSecret(ctx context.Context, secret model.Secret) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	secretsByType := m.ensureSecrets(secret.GetType())
	identity := secret.GetIdentity()
	_, ok := secretsByType[identity]
	if !ok {
		return logger.WrapError("change secret", storage.ErrSecretNotFound)
	}

	secretsByType[identity] = secret
	return nil
}

func (m *memoryStorage) GetSecretByID(ctx context.Context, secretType model.SecretType, identity string) (model.Secret, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	secretsByType := m.ensureSecrets(secretType)
	secret, ok := secretsByType[identity]
	if !ok {
		return nil, logger.WrapError("get secret", storage.ErrSecretNotFound)
	}

	return secret, nil
}

func (m *memoryStorage) GetAllSecrets(ctx context.Context, secretType model.SecretType) ([]model.Secret, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	secretsByType := m.ensureSecrets(secretType)
	secretsLen := len(secretsByType)
	result := make([]model.Secret, secretsLen)

	i := 0
	for _, secret := range secretsByType {
		result[i] = secret
		i++
	}

	return result, nil
}

func (m *memoryStorage) RemoveSecret(ctx context.Context, secret model.Secret) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	secretsByType := m.ensureSecrets(secret.GetType())
	identity := secret.GetIdentity()
	delete(secretsByType, identity)

	return nil
}

func (m *memoryStorage) ensureSecrets(secretType model.SecretType) map[string]model.Secret {
	secretsByType, ok := m.secrets[secretType]
	if !ok {
		secretsByType = make(map[string]model.Secret)
		m.secrets[secretType] = secretsByType
	}

	return secretsByType
}
