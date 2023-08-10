package jwt

import (
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
	"github.com/golang-jwt/jwt/v5"
)

type TokenParserConfig interface {
	SecretKey() []byte
}

type jwtTokenParser struct {
	secretKey []byte
}

func NewParser(conf TokenParserConfig) *jwtTokenParser {
	return &jwtTokenParser{
		secretKey: conf.SecretKey(),
	}
}

func (p *jwtTokenParser) Parse(tokenString string) (*jwt.Token, *jwt.RegisteredClaims, error) {
	claims := &jwt.RegisteredClaims{}
	parsedToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return p.secretKey, nil
	})
	if err != nil {
		return nil, nil, logger.WrapError("parse token", err)
	}

	return parsedToken, claims, nil
}
