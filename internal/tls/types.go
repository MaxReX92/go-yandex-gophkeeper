package tls

import (
	"crypto/tls"

	"google.golang.org/grpc/credentials"
)

type TLSProvider interface {
	GetTransportCredentials() (credentials.TransportCredentials, error)
	GetTLSConfig() (*tls.Config, error)
}
