package auth

import "context"

type CredentialsProvider interface {
	GetCredentials() (*Credentials, error)
	Register(ctx context.Context, userName string, password string) error
	Login(ctx context.Context, userName string, password string) error
}
