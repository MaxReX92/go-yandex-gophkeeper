package secret

import "github.com/MaxReX92/go-yandex-gophkeeper/internal/model"

type BaseSecret struct {
	identity   string
	secretType model.SecretType
	comment    string
}

func newBaseSecret(identity string, secretType model.SecretType, comment string) *BaseSecret {
	return &BaseSecret{
		comment:    comment,
		identity:   identity,
		secretType: secretType,
	}
}

func (s *BaseSecret) GetIdentity() string {
	return s.identity
}

func (s *BaseSecret) GetType() model.SecretType {
	return s.secretType
}

func (s *BaseSecret) GetComment() string {
	return s.comment
}
