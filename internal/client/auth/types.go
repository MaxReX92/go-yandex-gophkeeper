package auth

type Credentials interface {
	GetUserName() string
}
