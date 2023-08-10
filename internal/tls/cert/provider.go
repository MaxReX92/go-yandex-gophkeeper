package cert

import (
	"crypto/tls"

	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
	"google.golang.org/grpc/credentials"
)

type CertTLSProviderConfig interface {
	GetPublicCertPath() string
	GetPrivateKeyPath() string
}

type certTLSProvider struct {
	publicCert string
	privateKey string
}

func NewTLSProvider(conf CertTLSProviderConfig) *certTLSProvider {
	return &certTLSProvider{
		publicCert: conf.GetPublicCertPath(),
		privateKey: conf.GetPrivateKeyPath(),
	}
}

func (c *certTLSProvider) GetTransportCredentials() (credentials.TransportCredentials, error) {
	cert, err := c.loadTlsCert()
	if err != nil {
		return nil, err
	}

	return credentials.NewTLS(&tls.Config{
		Certificates:       []tls.Certificate{*cert},
		InsecureSkipVerify: true,
	}), nil
}

func (c *certTLSProvider) GetTlsConfig() (*tls.Config, error) {
	cert, err := c.loadTlsCert()
	if err != nil {
		return nil, err
	}

	return &tls.Config{Certificates: []tls.Certificate{*cert}}, nil
}

func (c *certTLSProvider) loadTlsCert() (*tls.Certificate, error) {
	cert, err := tls.LoadX509KeyPair(c.publicCert, c.privateKey)
	if err != nil {
		return nil, logger.WrapError("load certificate", err)
	}

	return &cert, nil
}