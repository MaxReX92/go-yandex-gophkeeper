package jwt

import (
	"time"

	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
	"github.com/golang-jwt/jwt/v5"
)

type TokenValidatorConfig interface {
	SecretKey() []byte
}

type jwtTokenValidator struct {
	secretKey []byte
}

func NewValidator(conf TokenParserConfig) *jwtTokenValidator {
	return &jwtTokenValidator{
		secretKey: conf.SecretKey(),
	}
}

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
