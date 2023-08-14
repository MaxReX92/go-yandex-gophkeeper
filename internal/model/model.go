package model

type (
	SecretType int32
	EventType  int32
)

const (
	Binary     SecretType = 0
	Card       SecretType = 1
	Credential SecretType = 2
	Note       SecretType = 3
)

const (
	Initial EventType = 0
	Add     EventType = 1
	Edit    EventType = 2
	Remove  EventType = 3
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

type User struct {
	Identity      string
	Name          string
	Password      string
	PersonalToken string
}
