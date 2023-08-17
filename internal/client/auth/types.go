package auth

import "context"

// CredentialsProvider implement functions for user authentication.
type CredentialsProvider interface {
	// GetCredentials current user credentials.
	GetCredentials() (*Credentials, error)
	// Register register new user in system.
	Register(ctx context.Context, userName string, password string) error
	// Login auth user by passed credentials.
	Login(ctx context.Context, userName string, password string) error
}
