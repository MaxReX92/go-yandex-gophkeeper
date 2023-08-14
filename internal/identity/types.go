package identity

type Generator interface {
	GenerateNewIdentity() string
}
