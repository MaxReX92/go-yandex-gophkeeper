package model

type SecretType int32

const (
	Binary      SecretType = 0
	Card                   = 1
	Credentials            = 2
	Notes                  = 3
)

type Secret interface {
	GetIdentity() string
	GetComment() string
	GetType() SecretType
}
