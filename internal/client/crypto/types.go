package crypto

type Encryptor interface {
	Encrypt(bytes []byte, key []byte) ([]byte, error)
}

type Decryptor interface {
	Decrypt(bytes []byte, key []byte) ([]byte, error)
}
