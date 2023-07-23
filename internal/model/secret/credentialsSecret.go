package secret

import "github.com/MaxReX92/go-yandex-gophkeeper/internal/model"

type credentialsSecret struct {
	*baseSecret

	userName string
	password string
}

func newCredentialsSecret(userName string, password string, identity string, comment string) *credentialsSecret {
	return &credentialsSecret{
		baseSecret: newBaseSecret(identity, model.Credentials, comment),
		userName:   userName,
		password:   password,
	}
}
