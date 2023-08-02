package model

type SecretType int32

const (
	Binary     SecretType = 0
	Card                  = 1
	Credential            = 2
	Note                  = 3
)

type Secret interface {
	GetIdentity() string
	GetComment() string
	GetType() SecretType
}
