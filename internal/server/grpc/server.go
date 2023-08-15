package grpc

import (
	"context"
	"crypto/tls"
	"sync"

	"golang.org/x/sync/errgroup"
	rpc "google.golang.org/grpc"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/db"
	tlsCert "github.com/MaxReX92/go-yandex-gophkeeper/internal/tls"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/generated"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model/grpc"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

// GrpcServerConfig contains required configuration for grpc server instance.
type GrpcServerConfig interface {
	GrpcAddress() string
}

type grpcServer struct {
	generated.UnimplementedSecretServiceServer

	listenAddress string
	eventChannels map[string][]chan *generated.SecretEvent
	converter     *grpc.Converter
	server        *rpc.Server
	dbService     db.Service
	tlsProvider   tlsCert.TLSProvider
	lock          sync.RWMutex
}

func NewServer(conf GrpcServerConfig, dbService db.Service, converter *grpc.Converter, tlsProvider tlsCert.TLSProvider) *grpcServer {
	return &grpcServer{
		listenAddress: conf.GrpcAddress(),
		eventChannels: map[string][]chan *generated.SecretEvent{},
		dbService:     dbService,
		converter:     converter,
		server:        rpc.NewServer(),
		tlsProvider:   tlsProvider,
		lock:          sync.RWMutex{},
	}
}

func (g *grpcServer) Start(_ context.Context) error {
	tlsConfig, err := g.tlsProvider.GetTLSConfig()
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
	err := g.dbService.CallInTransaction(ctx, func(ctx context.Context, executor db.Executor) error {
		err := executor.AddSecret(ctx, request.User.Identity, request.Secret)
		if err != nil {
			return logger.WrapError("add secret", err)
		}

		g.writeToAllChannels(request.User.Identity, &generated.SecretEvent{
			Type:   generated.EventType_ADD,
			Secret: request.Secret,
		})
		return nil
	})
	if err != nil {
		return nil, logger.WrapError("call add secret query", err)
	}

	return &generated.Void{}, nil
}

func (g *grpcServer) ChangeSecret(ctx context.Context, request *generated.SecretRequest) (*generated.Void, error) {
	err := g.dbService.CallInTransaction(ctx, func(ctx context.Context, executor db.Executor) error {
		err := executor.ChangeSecret(ctx, request.User.Identity, request.Secret)
		if err != nil {
			return logger.WrapError("change secret", err)
		}

		g.writeToAllChannels(request.User.Identity, &generated.SecretEvent{
			Type:   generated.EventType_EDIT,
			Secret: request.Secret,
		})

		return nil
	})
	if err != nil {
		return nil, logger.WrapError("call change secret query", err)
	}

	return &generated.Void{}, nil
}

func (g *grpcServer) RemoveSecret(ctx context.Context, request *generated.SecretRequest) (*generated.Void, error) {
	err := g.dbService.CallInTransaction(ctx, func(ctx context.Context, executor db.Executor) error {
		err := executor.RemoveSecret(ctx, request.User.Identity, request.Secret)
		if err != nil {
			return logger.WrapError("remove secret", err)
		}

		g.writeToAllChannels(request.User.Identity, &generated.SecretEvent{
			Type:   generated.EventType_REMOVE,
			Secret: request.Secret,
		})

		return nil
	})
	if err != nil {
		return nil, logger.WrapError("call remove secret query", err)
	}

	return &generated.Void{}, nil
}

func (g *grpcServer) SecretEvents(user *generated.User, stream generated.SecretService_SecretEventsServer) error {
	userID := user.Identity
	eventChannel := g.ensureEventChannel(userID)
	eg, ctx := errgroup.WithContext(stream.Context())

	eg.Go(func() error {
		secretEvents, err := g.dbService.CallInTransactionResult(ctx, func(ctx context.Context, executor db.Executor) ([]any, error) {
			dbSecrets, err := executor.GetAllSecrets(ctx, userID)
			if err != nil {
				return nil, logger.WrapError("load current secrets list", err)
			}

			secretsLen := len(dbSecrets)
			events := make([]any, secretsLen)
			for i := 0; i < secretsLen; i++ {
				events[i] = &generated.SecretEvent{
					Type:   generated.EventType_INITIAL,
					Secret: dbSecrets[i],
				}
			}

			return events, nil
		})
		if err != nil {
			return logger.WrapError("call get all secrets query", err)
		}

		for _, secretEvent := range secretEvents {
			eventChannel <- secretEvent.(*generated.SecretEvent)
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
