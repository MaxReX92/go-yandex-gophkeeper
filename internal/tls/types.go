package tls

import (
	"crypto/tls"

	"google.golang.org/grpc/credentials"
)

type CredentialsProvider interface {
	GetTransportCredentials() (credentials.TransportCredentials, error)
	GetTlsConfig() (*tls.Config, error)
}
