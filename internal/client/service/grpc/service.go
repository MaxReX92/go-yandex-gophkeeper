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
	GrpcServerAddress() string
}

type grpcService struct {
	client      generated.SecretServiceClient
	credentials auth.Credentials
	converter   *grpc.Converter
}

func NewService(
	conf GrpcServiceConfig,
	credentials auth.Credentials,
	converter *grpc.Converter,
	credentialsProvider tls.CredentialsProvider,
) (*grpcService, error) {

	transportCredentials, err := credentialsProvider.GetTransportCredentials()
	if err != nil {
		return nil, logger.WrapError("load transport credentials", err)
	}

	connection, err := rpc.Dial(conf.GrpcServerAddress(), rpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		return nil, logger.WrapError("open grpc connection", err)
	}

	return &grpcService{
		client:      generated.NewSecretServiceClient(connection),
		credentials: credentials,
		converter:   converter,
	}, nil
}

func (s *grpcService) AddSecret(ctx context.Context, secret model.Secret) error {
	grpcSecret, err := s.converter.FromModelSecret(secret)
	if err != nil {
		return logger.WrapError("convert model secret", err)
	}

	_, err = s.client.AddSecret(ctx, s.createSecretRequest(grpcSecret))
	if err != nil {
		return logger.WrapError("send add secret request", err)
	}

	return nil
}

func (s *grpcService) ChangeSecret(ctx context.Context, secret model.Secret) error {
	grpcSecret, err := s.converter.FromModelSecret(secret)
	if err != nil {
		return logger.WrapError("convert model secret", err)
	}

	_, err = s.client.ChangeSecret(ctx, s.createSecretRequest(grpcSecret))
	if err != nil {
		return logger.WrapError("send edit secret request", err)
	}

	return nil
}

func (s *grpcService) RemoveSecret(ctx context.Context, secret model.Secret) error {
	grpcSecret, err := s.converter.FromModelSecret(secret)
	if err != nil {
		return logger.WrapError("convert model secret", err)
	}

	_, err = s.client.RemoveSecret(ctx, s.createSecretRequest(grpcSecret))
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
				user := generated.User{Identity: s.credentials.GetUserName()}
				eventStream, err := s.client.SecretEvents(ctx, &user)
				if err != nil {
					logger.ErrorFormat("failed to get events stream: %v", err)
				} else {
					s.receiveEvents(ctx, eventStream, result)
				}

				if ctx.Err() == nil {
					time.Sleep(1 * time.Second)
				}
			}
		}
	}()

	return result
}

func (s *grpcService) receiveEvents(ctx context.Context, eventStream generated.SecretService_SecretEventsClient, result chan<- *model.SecretEvent) {
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
				modelEvent, err := s.converter.ToModelEvent(event)
				if err != nil {
					logger.ErrorFormat("failed to convert secret event: %v", err)
				} else {
					result <- modelEvent
				}
			}
		}
	}
}

func (s *grpcService) createSecretRequest(secret *generated.Secret) *generated.SecretRequest {
	return &generated.SecretRequest{
		User:   &generated.User{Identity: s.credentials.GetUserName()},
		Secret: secret,
	}
}
