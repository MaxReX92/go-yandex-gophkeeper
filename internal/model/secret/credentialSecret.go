package secret

import "github.com/MaxReX92/go-yandex-gophkeeper/internal/model"

// CredentialSecret represent a credential secret.
type CredentialSecret struct {
	*BaseSecret

	UserName string
	Password string
}

// NewCredentialSecret creates a new instance of credential secret.
func NewCredentialSecret(userName string, password string, identity string, comment string) *CredentialSecret {
	return &CredentialSecret{
		BaseSecret: newBaseSecret(identity, model.Credential, comment),
		UserName:   userName,
		Password:   password,
	}
}
