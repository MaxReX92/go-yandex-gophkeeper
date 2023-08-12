package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"io"

	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

type aesEncryptor struct {
	publicKey *rsa.PublicKey
}

func NewEncryptor() *aesEncryptor {
	return &aesEncryptor{}
}

func (r *aesEncryptor) Encrypt(bytes []byte, key []byte) ([]byte, error) {
	aesBlock, err := aes.NewCipher(key)
	blockSize := aesBlock.BlockSize()
	if err != nil {
		return nil, logger.WrapError("create cipher", err)
	}

	encrypted := make([]byte, blockSize+len(bytes))
	iv := encrypted[:blockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, logger.WrapError("create cipher", err)
	}

	stream := cipher.NewCFBEncrypter(aesBlock, iv)
	stream.XORKeyStream(encrypted[blockSize:], bytes)
	return encrypted, nil
}
