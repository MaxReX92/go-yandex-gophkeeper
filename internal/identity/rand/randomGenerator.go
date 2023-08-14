package rand

import (
	"math/rand"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type RandomGeneratorConfig interface {
	IdentityLength() int32
}

type randomGenerator struct {
	identityLen int32
	random      *rand.Rand
}

func NewGenerator(conf RandomGeneratorConfig) *randomGenerator {
	return &randomGenerator{
		identityLen: conf.IdentityLength(),
		random:      rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (r *randomGenerator) GenerateNewIdentity() string {
	identity := make([]byte, r.identityLen)
	for i := range identity {
		identity[i] = letters[r.random.Intn(len(letters))]
	}
	return string(identity)
}
