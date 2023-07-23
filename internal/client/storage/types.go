package storage

import (
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
)

type LocalSecretsStorage interface {
	AddSecret(secret model.Secret) error
	ChangeSecret(secret model.Secret) error
	GetSecretById(secretType model.SecretType, identity string) (model.Secret, error)
	GetAllSecrets(secretType model.SecretType) ([]model.Secret, error)
	RemoveSecret(secret model.Secret) error
}
