package rand

import "math/rand"

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type RandomGeneratorConfig interface {
	IdentityLen() int32
}

type randomGenerator struct {
	identityLen int32
}

func NewGenerator(conf RandomGeneratorConfig) *randomGenerator {
	return &randomGenerator{
		identityLen: conf.IdentityLen(),
	}
}

func (r *randomGenerator) GenerateNewIdentity() string {
	identity := make([]byte, r.identityLen)
	for i := range identity {
		identity[i] = letters[rand.Intn(len(letters))]
	}
	return string(identity)
}
