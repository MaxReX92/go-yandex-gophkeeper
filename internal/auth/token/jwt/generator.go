package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

type JwtTokenGeneratorConfig interface {
	SecretKey() []byte
	TokenTTL() time.Duration
}

type jwtTokenGenerator struct {
	secretKey []byte
	tokenTTL  time.Duration
}

func NewGenerator(conf JwtTokenGeneratorConfig) *jwtTokenGenerator {
	return &jwtTokenGenerator{
		secretKey: conf.SecretKey(),
		tokenTTL:  conf.TokenTTL(),
	}
}

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
