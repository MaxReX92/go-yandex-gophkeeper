package cert

import (
	"crypto/tls"

	"google.golang.org/grpc/credentials"

	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

// CertTLSProviderConfig contains required cert tls provider instance.
type CertTLSProviderConfig interface {
	GetPublicCertPath() string
	GetPrivateKeyPath() string
}

type certTLSProvider struct {
	publicCert string
	privateKey string
}

// NewTLSProvider creates a new instance of cert tls provider.
func NewTLSProvider(conf CertTLSProviderConfig) *certTLSProvider {
	return &certTLSProvider{
		publicCert: conf.GetPublicCertPath(),
		privateKey: conf.GetPrivateKeyPath(),
	}
}

func (c *certTLSProvider) GetTransportCredentials() (credentials.TransportCredentials, error) {
	cert, err := c.loadTLSCert()
	if err != nil {
		return nil, err
	}

	return credentials.NewTLS(&tls.Config{
		Certificates:       []tls.Certificate{*cert},
		InsecureSkipVerify: true,
	}), nil
}

func (c *certTLSProvider) GetTLSConfig() (*tls.Config, error) {
	cert, err := c.loadTLSCert()
	if err != nil {
		return nil, err
	}

	return &tls.Config{Certificates: []tls.Certificate{*cert}}, nil
}

func (c *certTLSProvider) loadTLSCert() (*tls.Certificate, error) {
	cert, err := tls.LoadX509KeyPair(c.publicCert, c.privateKey)
	if err != nil {
		return nil, logger.WrapError("load certificate", err)
	}

	return &cert, nil
}
