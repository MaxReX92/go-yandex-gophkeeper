package secret

import "github.com/MaxReX92/go-yandex-gophkeeper/internal/model"

// BaseSecret represent secret with base sign of secret.
type BaseSecret struct {
	Identity   string
	SecretType model.SecretType
	Comment    string
}

func newBaseSecret(identity string, secretType model.SecretType, comment string) *BaseSecret {
	return &BaseSecret{
		Comment:    comment,
		Identity:   identity,
		SecretType: secretType,
	}
}

func (s *BaseSecret) GetIdentity() string {
	return s.Identity
}

func (s *BaseSecret) GetType() model.SecretType {
	return s.SecretType
}

func (s *BaseSecret) GetComment() string {
	return s.Comment
}
