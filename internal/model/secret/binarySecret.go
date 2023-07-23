package secret

import "github.com/MaxReX92/go-yandex-gophkeeper/internal/model"

type binarySecret struct {
	*baseSecret

	filePath string
}

func newBinarySecret(filePath string, identity string, comment string) *binarySecret {
	return &binarySecret{
		baseSecret: newBaseSecret(identity, model.Credentials, comment),
		filePath:   filePath,
	}
}
