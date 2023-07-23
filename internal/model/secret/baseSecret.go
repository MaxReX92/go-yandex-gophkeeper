package secret

import "github.com/MaxReX92/go-yandex-gophkeeper/internal/model"

type baseSecret struct {
	identity   string
	secretType model.SecretType
	comment    string
}

func newBaseSecret(identity string, secretType model.SecretType, comment string) *baseSecret {
	return &baseSecret{
		comment:    comment,
		identity:   identity,
		secretType: secretType,
	}
}

func (s *baseSecret) GetIdentity() string {
	return s.identity
}

func (s *baseSecret) GetType() model.SecretType {
	return s.secretType
}

func (s *baseSecret) GetComment() string {
	return s.comment
}
