package grpc

import (
	"context"
	"sync"
	"time"

	rpc "google.golang.org/grpc"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/auth"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/generated"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/tls"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

// GrpcCredentialsProviderConfig contains required configuration for grpc credentials provider instance.
type GrpcCredentialsProviderConfig interface {
	// AuthServerAddress returns auth server grpc address.
	AuthServerAddress() string
	// RenewTokenInterval returns interval for renew auth token.
	RenewTokenInterval() time.Duration
}

type grpcCredentialsProvider struct {
	client             generated.AuthServiceClient
	lock               sync.RWMutex
	renewTokenInterval time.Duration
	credentials        *auth.Credentials
}

// NewProvider creates a new instance of grpc credentials provider.
func NewProvider(
	conf GrpcCredentialsProviderConfig,
	tlsProvider tls.TLSProvider,
) (*grpcCredentialsProvider, error) {
	transportCredentials, err := tlsProvider.GetTransportCredentials()
	if err != nil {
		return nil, logger.WrapError("load transport credentials", err)
	}

	connection, err := rpc.Dial(conf.AuthServerAddress(), rpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		return nil, logger.WrapError("open grpc connection", err)
	}

	return &grpcCredentialsProvider{
		client:             generated.NewAuthServiceClient(connection),
		lock:               sync.RWMutex{},
		renewTokenInterval: conf.RenewTokenInterval(),
	}, nil
}

func (c *grpcCredentialsProvider) Start(ctx context.Context) error {
	ticker := time.NewTicker(c.renewTokenInterval)
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			c.lock.Lock()
			defer c.lock.Unlock()

			if c.credentials == nil {
				continue
			}

			request := &generated.ProlongTokenRequest{Token: c.credentials.Token}
			response, err := c.client.Prolong(ctx, request)
			if err != nil {
				logger.ErrorFormat("failed to prolong token: %v", err)
			}

			c.credentials.Token = response.Token
		}
	}
}

func (c *grpcCredentialsProvider) GetCredentials() (*auth.Credentials, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if c.credentials == nil {
		return nil, logger.WrapError("get credentials", auth.ErrUnauthorized)
	}

	return c.credentials, nil
}

func (c *grpcCredentialsProvider) Register(ctx context.Context, userName string, password string) error {
	request := &generated.RegisterRequest{
		Name:     userName,
		Password: password,
	}

	response, err := c.client.Register(ctx, request)
	if err != nil {
		return logger.WrapError("register new user", err)
	}

	logger.InfoFormat("New user identity: %s", response.Identity)
	return nil
}

func (c *grpcCredentialsProvider) Login(ctx context.Context, userName string, password string) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.credentials != nil {
		return logger.WrapError("login", auth.ErrAlreadyAuthorized)
	}

	request := &generated.LoginRequest{
		Name:     userName,
		Password: password,
	}

	response, err := c.client.Login(ctx, request)
	if err != nil {
		return logger.WrapError("login", err)
	}

	c.credentials = &auth.Credentials{
		Identity:      response.Identity,
		UserName:      userName,
		Token:         response.Token,
		PersonalToken: response.PersonalToken,
	}

	return nil
}
