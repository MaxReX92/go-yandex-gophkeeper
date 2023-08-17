package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

// JwtTokenValidatorConfig contains required configuration for jwt token validator instance.
type JwtTokenValidatorConfig interface {
	// SecretKey returns service crypto secret key.
	SecretKey() []byte
}

type jwtTokenValidator struct {
	secretKey []byte
}

// NewValidator creates a new instance of jwt token validator.
func NewValidator(conf JwtTokenValidatorConfig) *jwtTokenValidator {
	return &jwtTokenValidator{
		secretKey: conf.SecretKey(),
	}
}

// Check provides methods for auth jwt token validation.
func (v *jwtTokenValidator) Check(tokenString string) (bool, error) {
	claims := &jwt.RegisteredClaims{}
	parsedToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return v.secretKey, nil
	})
	if err != nil {
		return false, logger.WrapError("parse token", err)
	}

	now := jwt.NewNumericDate(time.Now()).Unix()
	return parsedToken.Valid &&
		now >= claims.NotBefore.Unix() &&
		now <= claims.ExpiresAt.Unix(), nil
}
