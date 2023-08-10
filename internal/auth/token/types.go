package token

type Generator interface {
	GenerateToken() (string, error)
}

type Validator interface {
	Check(tokenString string) (bool, error)
}
