package tls

import (
	"crypto/tls"

	"google.golang.org/grpc/credentials"
)

// TLSProvider provide transport layer security configuration.
type TLSProvider interface {
	// GetTransportCredentials return client transport layer security configuration.
	GetTransportCredentials() (credentials.TransportCredentials, error)
	// GetTLSConfig return server transport layer security configuration.
	GetTLSConfig() (*tls.Config, error)
}
