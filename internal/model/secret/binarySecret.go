package secret

import "github.com/MaxReX92/go-yandex-gophkeeper/internal/model"

type BinarySecret struct {
	*BaseSecret

	FilePath string
}

func NewBinarySecret(filePath string, identity string, comment string) *BinarySecret {
	return &BinarySecret{
		BaseSecret: newBaseSecret(identity, model.Credential, comment),
		FilePath:   filePath,
	}
}
