package grpc

import (
	"context"
	"crypto/tls"
	"sync"

	tlsCert "github.com/MaxReX92/go-yandex-gophkeeper/internal/tls"
	"golang.org/x/sync/errgroup"
	rpc "google.golang.org/grpc"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/generated"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model/grpc"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/server/storage"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

type GrpcServerConfig interface {
	GrpcAddress() string
}

type grpcServer struct {
	generated.UnimplementedSecretServiceServer

	listenAddress       string
	eventChannels       map[string][]chan *generated.SecretEvent
	storage             storage.SecretsStorage
	converter           *grpc.Converter
	server              *rpc.Server
	credentialsProvider tlsCert.CredentialsProvider
	lock                sync.RWMutex
}

func NewGrpcServer(conf GrpcServerConfig, storage storage.SecretsStorage, converter *grpc.Converter, credentialsProvider tlsCert.CredentialsProvider) *grpcServer {
	return &grpcServer{
		listenAddress:       conf.GrpcAddress(),
		eventChannels:       map[string][]chan *generated.SecretEvent{},
		storage:             storage,
		converter:           converter,
		server:              rpc.NewServer(),
		credentialsProvider: credentialsProvider,
		lock:                sync.RWMutex{},
	}
}

func (g *grpcServer) Start(_ context.Context) error {
	tlsConfig, err := g.credentialsProvider.GetTlsConfig()
	if err != nil {
		return logger.WrapError("create tls config", err)
	}

	listen, err := tls.Listen("tcp", g.listenAddress, tlsConfig)
	if err != nil {
		return logger.WrapError("start listen TCP", err)
	}
	generated.RegisterSecretServiceServer(g.server, g)

	logger.Info("Start gRPC service")
	err = g.server.Serve(listen)
	if err != nil {
		return logger.WrapError("start gRPC service", err)
	}

	return nil
}

func (g *grpcServer) Stop(ctx context.Context) error {
	logger.Info("Stopping gRPC service")
	g.server.GracefulStop()
	return nil
}

func (g *grpcServer) Ping(ctx context.Context, void *generated.Void) (*generated.Void, error) {
	return &generated.Void{}, nil
}

func (g *grpcServer) AddSecret(ctx context.Context, request *generated.SecretRequest) (*generated.Void, error) {
	err := g.storage.AddSecret(ctx, request.User.Identity, request.Secret)
	if err != nil {
		return nil, logger.WrapError("add secret", err)
	}

	g.writeToAllChannels(request.User.Identity, &generated.SecretEvent{
		Type:   generated.EventType_ADD,
		Secret: request.Secret,
	})

	return &generated.Void{}, nil
}

func (g *grpcServer) ChangeSecret(ctx context.Context, request *generated.SecretRequest) (*generated.Void, error) {
	err := g.storage.ChangeSecret(ctx, request.User.Identity, request.Secret)
	if err != nil {
		return nil, logger.WrapError("change secret", err)
	}

	g.writeToAllChannels(request.User.Identity, &generated.SecretEvent{
		Type:   generated.EventType_EDIT,
		Secret: request.Secret,
	})

	return &generated.Void{}, nil
}

func (g *grpcServer) RemoveSecret(ctx context.Context, request *generated.SecretRequest) (*generated.Void, error) {
	err := g.storage.RemoveSecret(ctx, request.User.Identity, request.Secret)
	if err != nil {
		return nil, logger.WrapError("remove secret", err)
	}

	g.writeToAllChannels(request.User.Identity, &generated.SecretEvent{
		Type:   generated.EventType_REMOVE,
		Secret: request.Secret,
	})

	return &generated.Void{}, nil
}

func (g *grpcServer) SecretEvents(user *generated.User, stream generated.SecretService_SecretEventsServer) error {
	userID := user.Identity
	eventChannel := g.ensureEventChannel(userID)
	eg, ctx := errgroup.WithContext(stream.Context())

	eg.Go(func() error {
		currentSecrets, err := g.storage.GetAllSecrets(ctx, userID)
		if err != nil {
			logger.ErrorFormat("failed to load current secrets list: %v", err)
			return logger.WrapError("load current secrets list", err)
		}

		for _, currentSecret := range currentSecrets {
			eventChannel <- &generated.SecretEvent{
				Type:   generated.EventType_INITIAL,
				Secret: currentSecret,
			}
		}

		return nil
	})

	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return nil
			case event := <-eventChannel:
				err := stream.Send(event)
				if err != nil {
					return logger.WrapError("send message", err)
				}
			}
		}
	})

	return eg.Wait()
}

func (g *grpcServer) ensureEventChannel(userID string) chan *generated.SecretEvent {
	g.lock.Lock()
	defer g.lock.Unlock()
	channelList, ok := g.eventChannels[userID]
	if !ok {
		channelList = make([]chan *generated.SecretEvent, 0)
		g.eventChannels[userID] = channelList
	}

	eventChannel := make(chan *generated.SecretEvent, 10)
	g.eventChannels[userID] = append(channelList, eventChannel)
	return eventChannel
}

func (g *grpcServer) writeToAllChannels(userID string, event *generated.SecretEvent) {
	g.lock.RLock()
	defer g.lock.RUnlock()

	channelList, ok := g.eventChannels[userID]
	if !ok {
		return
	}

	for _, channel := range channelList {
		channel <- event
	}
}
