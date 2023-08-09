package cert

import (
	"crypto/tls"

	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
	"google.golang.org/grpc/credentials"
)

type CertCredentialsProviderConfig interface {
	GetPublicCertPath() string
	GetPrivateKeyPath() string
}

type certCredentialsProvider struct {
	publicCert string
	privateKey string
}

func NewCredentialsProvider(conf CertCredentialsProviderConfig) *certCredentialsProvider {
	return &certCredentialsProvider{
		publicCert: conf.GetPublicCertPath(),
		privateKey: conf.GetPrivateKeyPath(),
	}
}

func (c *certCredentialsProvider) GetTransportCredentials() (credentials.TransportCredentials, error) {
	cert, err := c.loadTlsCert()
	if err != nil {
		return nil, err
	}

	return credentials.NewTLS(&tls.Config{
		Certificates:       []tls.Certificate{*cert},
		InsecureSkipVerify: true,
	}), nil
}

func (c *certCredentialsProvider) GetTlsConfig() (*tls.Config, error) {
	cert, err := c.loadTlsCert()
	if err != nil {
		return nil, err
	}

	return &tls.Config{Certificates: []tls.Certificate{*cert}}, nil
}

func (c *certCredentialsProvider) loadTlsCert() (*tls.Certificate, error) {
	cert, err := tls.LoadX509KeyPair(c.publicCert, c.privateKey)
	if err != nil {
		return nil, logger.WrapError("load certificate", err)
	}

	return &cert, nil
}
