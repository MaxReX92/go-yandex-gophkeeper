package token

// Generator provides methods for auth token generation.
type Generator interface {
	// GenerateToken generate new auth token.
	GenerateToken() (string, error)
}

// Validator provides methods for auth token validation.
type Validator interface {
	// Check validate received auth token.
	Check(tokenString string) (bool, error)
}
