package auth

// Credentials represent user auth credentials.
type Credentials struct {
	Identity      string
	UserName      string
	Token         string
	PersonalToken string
}
