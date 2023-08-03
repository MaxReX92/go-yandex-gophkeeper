package remote

import "github.com/MaxReX92/go-yandex-gophkeeper/internal/model"

type remoteStorage struct {
}

func NewStorage() *remoteStorage {
	return &remoteStorage{}
}

func (r *remoteStorage) AddSecret(secret model.Secret) error {
	//TODO implement me
	panic("implement me")
}

func (r *remoteStorage) ChangeSecret(secret model.Secret) error {
	//TODO implement me
	panic("implement me")
}

func (r *remoteStorage) GetSecretById(secretType model.SecretType, identity string) (model.Secret, error) {
	//TODO implement me
	panic("implement me")
}

func (r *remoteStorage) GetAllSecrets(secretType model.SecretType) ([]model.Secret, error) {
	//TODO implement me
	panic("implement me")
}

func (r *remoteStorage) RemoveSecret(secret model.Secret) error {
	//TODO implement me
	panic("implement me")
}
