package grpc

import (
	"context"
	"net"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/generated"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model/grpc"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/server/storage"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
	rpc "google.golang.org/grpc"
)

type GrpcServerConfig interface {
	GrpcAddress() string
}

type grpcServer struct {
	generated.UnimplementedSecretServiceServer

	listenAddress string
	events        chan *generated.SecretEvent
	storage       storage.SecretsStorage
	converter     *grpc.Converter
	server        *rpc.Server
}

func NewGrpcServer(conf GrpcServerConfig, storage storage.SecretsStorage, converter *grpc.Converter) *grpcServer {
	return &grpcServer{
		listenAddress: conf.GrpcAddress(),
		events:        make(chan *generated.SecretEvent, 10),
		storage:       storage,
		converter:     converter,
		server:        rpc.NewServer(),
	}
}

func (g *grpcServer) Start(_ context.Context) error {
	listen, err := net.Listen("tcp", g.listenAddress)
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

	g.events <- &generated.SecretEvent{
		Type:   generated.EventType_ADD,
		Secret: request.Secret,
	}

	return &generated.Void{}, nil
}

func (g *grpcServer) ChangeSecret(ctx context.Context, request *generated.SecretRequest) (*generated.Void, error) {
	err := g.storage.ChangeSecret(ctx, request.User.Identity, request.Secret)
	if err != nil {
		return nil, logger.WrapError("change secret", err)
	}

	g.events <- &generated.SecretEvent{
		Type:   generated.EventType_EDIT,
		Secret: request.Secret,
	}

	return &generated.Void{}, nil
}
func (g *grpcServer) RemoveSecret(ctx context.Context, request *generated.SecretRequest) (*generated.Void, error) {
	err := g.storage.RemoveSecret(ctx, request.User.Identity, request.Secret)
	if err != nil {
		return nil, logger.WrapError("remove secret", err)
	}

	g.events <- &generated.SecretEvent{
		Type:   generated.EventType_REMOVE,
		Secret: request.Secret,
	}

	return &generated.Void{}, nil
}

func (g *grpcServer) SecretEvents(user *generated.User, stream generated.SecretService_SecretEventsServer) error {
	go func(userId string) {
		currentSecrets, err := g.storage.GetAllSecrets(stream.Context(), userId)
		if err != nil {
			logger.ErrorFormat("failed to load current secrets list: %v", err)
			return
		}

		for _, currentSecret := range currentSecrets {
			g.events <- &generated.SecretEvent{
				Type:   generated.EventType_INITIAL,
				Secret: currentSecret,
			}
		}
	}(user.Identity)

	for {
		select {
		case <-stream.Context().Done():
			return nil
		case event := <-g.events:
			err := stream.Send(event)
			if err != nil {
				return logger.WrapError("send message", err)
			}
		}
	}
}
