package secret

import "github.com/MaxReX92/go-yandex-gophkeeper/internal/model"

type CredentialsSecret struct {
	*BaseSecret

	UserName string
	Password string
}

func NewCredentialsSecret(userName string, password string, identity string, comment string) *CredentialsSecret {
	return &CredentialsSecret{
		BaseSecret: newBaseSecret(identity, model.Credentials, comment),
		UserName:   userName,
		Password:   password,
	}
}
