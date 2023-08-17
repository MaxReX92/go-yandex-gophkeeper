package model

type (
	// SecretType represents a secret type.
	SecretType int32
	// EventType represents a secret event type.
	EventType int32
)

const (
	// Binary is a type of binary secret.
	Binary SecretType = 0
	// Card is a type of card secret.
	Card SecretType = 1
	// Credential is a type of credential secret.
	Credential SecretType = 2
	// Note is a type of note secret.
	Note SecretType = 3
)

const (
	// Initial is a type of initial secret event.
	Initial EventType = 0
	// Add is a type of add secret event.
	Add EventType = 1
	// Edit is a type of edit secret event.
	Edit EventType = 2
	// Remove is a type of remove secret event.
	Remove EventType = 3
)

// Secret represents secret with base sign of secret.
type Secret interface {
	// GetIdentity returns secret unique identity.
	GetIdentity() string
	// GetComment returns secret comment.
	GetComment() string
	// GetType returns secret type.
	GetType() SecretType
}

// SecretEvent represents secret event.
type SecretEvent struct {
	// Type is a secret event type.
	Type EventType
	// Secret is a link to secret.
	Secret Secret
}

// User represents clients.
type User struct {
	// Identity is a user unique identity.
	Identity string
	// Name is a username.
	Name string
	// Name is a user password.
	Password string
	// PersonalToken is a user personal crypto token.
	PersonalToken string
}
