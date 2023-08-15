package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

// JwtTokenGeneratorConfig contains required configuration for jwt token generator instance.
type JwtTokenGeneratorConfig interface {
	// SecretKey returns service crypto secret key.
	SecretKey() []byte
	// TokenTTL return generated token time-to-live.
	TokenTTL() time.Duration
}

type jwtTokenGenerator struct {
	secretKey []byte
	tokenTTL  time.Duration
}

// NewGenerator creates jwt token generator instance.
func NewGenerator(conf JwtTokenGeneratorConfig) *jwtTokenGenerator {
	return &jwtTokenGenerator{
		secretKey: conf.SecretKey(),
		tokenTTL:  conf.TokenTTL(),
	}
}

// GenerateToken generate new jwt auth token.
func (m *jwtTokenGenerator) GenerateToken() (string, error) {
	now := time.Now()
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(now.Add(m.tokenTTL)),
		NotBefore: jwt.NewNumericDate(now),
	}

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(m.secretKey)
	if err != nil {
		return "", logger.WrapError("sign token", err)
	}

	return tokenString, nil
}
