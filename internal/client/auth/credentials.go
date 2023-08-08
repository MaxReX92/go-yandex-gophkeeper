package auth

type credentials struct {
	userName string
}

func NewCredentials(userName string) *credentials {
	return &credentials{
		userName: userName,
	}
}

func (c *credentials) GetUserName() string {
	return c.userName
}
