package model

type (
	SecretType int32
	EventType  int32
)

const (
	Binary     SecretType = 0
	Card                  = 1
	Credential            = 2
	Note                  = 3
)

const (
	Initial EventType = 0
	Add               = 1
	Edit              = 2
	Remove            = 3
)

type Secret interface {
	GetIdentity() string
	GetComment() string
	GetType() SecretType
}

type SecretEvent struct {
	Type   EventType
	Secret Secret
}
