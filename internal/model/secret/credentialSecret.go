package secret

import "github.com/MaxReX92/go-yandex-gophkeeper/internal/model"

type CredentialSecret struct {
	*BaseSecret

	UserName string
	Password string
}

func NewCredentialSecret(userName string, password string, identity string, comment string) *CredentialSecret {
	return &CredentialSecret{
		BaseSecret: newBaseSecret(identity, model.Credential, comment),
		UserName:   userName,
		Password:   password,
	}
}
