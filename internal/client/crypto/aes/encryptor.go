package aes

import (
	"crypto/aes"
	"crypto/rsa"

	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/chunk"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

type aesEncryptor struct {
	publicKey *rsa.PublicKey
}

func NewEncryptor() *aesEncryptor {
	return &aesEncryptor{}
}

func (r *aesEncryptor) Encrypt(bytes []byte, key []byte) ([]byte, error) {
	blockSize := len(key)
	aesblock, err := aes.NewCipher(key)
	if err != nil {
		return nil, logger.WrapError("create cipher", err)
	}

	var result []byte
	for _, messageChunk := range chunk.SliceToChunks(bytes, blockSize) {
		encryptedBlock := make([]byte, blockSize)
		aesblock.Encrypt(encryptedBlock, messageChunk)

		if err != nil {
			return nil, logger.WrapError("encrypt message", err)
		}
		result = append(result, encryptedBlock...)
	}

	return result, nil
}
