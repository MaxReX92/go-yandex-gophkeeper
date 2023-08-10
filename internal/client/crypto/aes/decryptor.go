package aes

import (
	"crypto/aes"

	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/chunk"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

type aesDecryptor struct {
}

func NewDecryptor() *aesDecryptor {
	return &aesDecryptor{}
}

func (r *aesDecryptor) Decrypt(bytes []byte, key []byte) ([]byte, error) {
	blockSize := len(key)
	aesblock, err := aes.NewCipher(key)
	if err != nil {
		return nil, logger.WrapError("create cipher", err)
	}

	var result []byte
	for _, messageChunk := range chunk.SliceToChunks(bytes, blockSize) {
		decryptedBlock := make([]byte, blockSize)
		aesblock.Decrypt(decryptedBlock, messageChunk)

		if err != nil {
			return nil, logger.WrapError("decrypt message", err)
		}
		result = append(result, decryptedBlock...)
	}

	return result, nil
}
