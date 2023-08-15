package identity

// Generator provides functions for some random generations.
type Generator interface {
	// GenerateNewIdentity generate new random identity.
	GenerateNewIdentity() string
}
