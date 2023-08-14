package secret

import (
	"io"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
)

type BinarySecret struct {
	*BaseSecret

	Name   string
	Reader io.Reader
}

func NewBinarySecret(name string, reader io.Reader, identity string, comment string) *BinarySecret {
	return &BinarySecret{
		BaseSecret: newBaseSecret(identity, model.Binary, comment),
		Name:       name,
		Reader:     reader,
	}
}
