package crypto

// Encryptor provides logic for encrypt byte arrays.
type Encryptor interface {
	// Encrypt encodes some byte array with provided secret key.
	Encrypt(bytes []byte, key []byte) ([]byte, error)
}

// Decryptor provides logic for decrypt byte arrays.
type Decryptor interface {
	// Decrypt decodes some byte array with provided secret key.
	Decrypt(bytes []byte, key []byte) ([]byte, error)
}
