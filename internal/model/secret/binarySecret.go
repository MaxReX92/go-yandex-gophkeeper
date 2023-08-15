package secret

import (
	"io"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
)

// BinarySecret represent a binary secret.
type BinarySecret struct {
	*BaseSecret

	Name   string
	Reader io.Reader
}

// NewBinarySecret creates a new instance of binary secret.
func NewBinarySecret(name string, reader io.Reader, identity string, comment string) *BinarySecret {
	return &BinarySecret{
		BaseSecret: newBaseSecret(identity, model.Binary, comment),
		Name:       name,
		Reader:     reader,
	}
}
