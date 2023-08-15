package aes

import (
	"crypto/aes"
	"crypto/cipher"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/crypto"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

type aesDecryptor struct{}

// NewDecryptor creates a new instance of aes dectyptor.
func NewDecryptor() *aesDecryptor {
	return &aesDecryptor{}
}

func (r *aesDecryptor) Decrypt(bytes []byte, key []byte) ([]byte, error) {
	aesBlock, err := aes.NewCipher(key)
	blockSize := aesBlock.BlockSize()
	if err != nil {
		return nil, logger.WrapError("create cipher", err)
	}

	if len(bytes) < blockSize {
		return nil, logger.WrapError("create cipher", crypto.ErrTooShortBlock)
	}

	iv := bytes[:blockSize]
	encrypted := bytes[blockSize:]
	dectypted := make([]byte, len(encrypted))

	stream := cipher.NewCFBDecrypter(aesBlock, iv)
	stream.XORKeyStream(dectypted, encrypted)

	return dectypted, nil
}
