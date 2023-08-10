package grpc

import (
	"context"
	"time"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/auth"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/generated"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model/grpc"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/tls"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
	rpc "google.golang.org/grpc"
)

type GrpcServiceConfig interface {
	SecretServerAddress() string
}

type grpcService struct {
	client              generated.SecretServiceClient
	credentialsProvider auth.CredentialsProvider
	converter           *grpc.Converter
}

func NewService(
	conf GrpcServiceConfig,
	credentialsProvider auth.CredentialsProvider,
	converter *grpc.Converter,
	tlsProvider tls.TLSProvider,
) (*grpcService, error) {

	transportCredentials, err := tlsProvider.GetTransportCredentials()
	if err != nil {
		return nil, logger.WrapError("load transport credentials", err)
	}

	connection, err := rpc.Dial(conf.SecretServerAddress(), rpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		return nil, logger.WrapError("open grpc connection", err)
	}

	return &grpcService{
		client:              generated.NewSecretServiceClient(connection),
		credentialsProvider: credentialsProvider,
		converter:           converter,
	}, nil
}

func (s *grpcService) AddSecret(ctx context.Context, secret model.Secret) error {
	credentials, err := s.credentialsProvider.GetCredentials()
	if err != nil {
		return err
	}

	grpcSecret, err := s.converter.FromModelSecret(secret, credentials.PersonalToken)
	if err != nil {
		return logger.WrapError("convert model secret", err)
	}

	request, err := s.createSecretRequest(grpcSecret, credentials)
	if err != nil {
		return logger.WrapError("create secret request", err)
	}

	_, err = s.client.AddSecret(ctx, request)
	if err != nil {
		return logger.WrapError("send add secret request", err)
	}

	return nil
}

func (s *grpcService) ChangeSecret(ctx context.Context, secret model.Secret) error {
	credentials, err := s.credentialsProvider.GetCredentials()
	if err != nil {
		return err
	}

	grpcSecret, err := s.converter.FromModelSecret(secret, credentials.PersonalToken)
	if err != nil {
		return logger.WrapError("convert model secret", err)
	}

	request, err := s.createSecretRequest(grpcSecret, credentials)
	if err != nil {
		return logger.WrapError("create secret request", err)
	}

	_, err = s.client.ChangeSecret(ctx, request)
	if err != nil {
		return logger.WrapError("send edit secret request", err)
	}

	return nil
}

func (s *grpcService) RemoveSecret(ctx context.Context, secret model.Secret) error {
	credentials, err := s.credentialsProvider.GetCredentials()
	if err != nil {
		return err
	}

	grpcSecret, err := s.converter.FromModelSecret(secret, credentials.PersonalToken)
	if err != nil {
		return logger.WrapError("convert model secret", err)
	}

	request, err := s.createSecretRequest(grpcSecret, credentials)
	if err != nil {
		return logger.WrapError("create secret request", err)
	}

	_, err = s.client.RemoveSecret(ctx, request)
	if err != nil {
		return logger.WrapError("send remove secret request", err)
	}

	return nil
}

func (s *grpcService) SecretEvents(ctx context.Context) <-chan *model.SecretEvent {
	result := make(chan *model.SecretEvent)
	go func() {
		defer close(result)

		for {
			select {
			case <-ctx.Done():
				logger.Info("Done")
				return
			default:
				credentials, err := s.credentialsProvider.GetCredentials()
				if err == nil {
					user := generated.User{Identity: credentials.Identity}
					eventStream, err := s.client.SecretEvents(ctx, &user)
					if err != nil {
						logger.ErrorFormat("failed to get events stream: %v", err)
					} else {
						s.receiveEvents(ctx, eventStream, credentials, result)
					}
				}

				if ctx.Err() == nil {
					time.Sleep(1 * time.Second)
				}
			}
		}
	}()

	return result
}

func (s *grpcService) receiveEvents(
	ctx context.Context,
	eventStream generated.SecretService_SecretEventsClient,
	credentials *auth.Credentials,
	result chan<- *model.SecretEvent,
) {
	for {
		select {
		case <-ctx.Done():
			logger.Info("Done")
			return
		default:
			event, err := eventStream.Recv()
			if err != nil {
				logger.ErrorFormat("failed to receive secret event: %v", err)
				return
			} else {
				modelEvent, err := s.converter.ToModelEvent(event, credentials.PersonalToken)
				if err != nil {
					logger.ErrorFormat("failed to convert secret event: %v", err)
				} else {
					result <- modelEvent
				}
			}
		}
	}
}

func (s *grpcService) createSecretRequest(secret *generated.Secret, credentials *auth.Credentials) (*generated.SecretRequest, error) {
	return &generated.SecretRequest{
		User:   &generated.User{Identity: credentials.Identity},
		Secret: secret,
	}, nil
}
