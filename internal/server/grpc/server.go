package grpc

import (
	"context"
	"net"
	"time"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/generated"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model/grpc"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model/secret"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/parser"
	rpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcServerConfig interface {
	GrpcAddress() string
}

type grpcServer struct {
	generated.UnimplementedSecretServiceServer

	listenAddress string
	converter     *grpc.Converter
	server        *rpc.Server
}

func NewGrpcServer(conf GrpcServerConfig, converter *grpc.Converter) *grpcServer {
	return &grpcServer{
		listenAddress: conf.GrpcAddress(),
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

func (g *grpcServer) AddSecret(context.Context, *generated.Secret) (*generated.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSecret not implemented")
}
func (g *grpcServer) ChangeSecret(context.Context, *generated.Secret) (*generated.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeSecret not implemented")
}
func (g *grpcServer) RemoveSecret(context.Context, *generated.Secret) (*generated.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSecret not implemented")
}

func (g *grpcServer) SecretEvents(user *generated.User, stream generated.SecretService_SecretEventsServer) error {
	ticker := time.NewTicker(1 * time.Second)

	for i := 0; ; i++ {
		select {
		case <-stream.Context().Done():
			return nil

		// тикер заменится на case с внутреннего канала, в который будут писать остальные методы после успешной обработки
		case <-ticker.C:

			modelSecret := secret.NewCredentialSecret(
				"userName"+parser.Int32ToString(int32(i)),
				"password"+parser.Int32ToString(int32(i)),
				"identity"+parser.Int32ToString(int32(i)),
				"comment"+parser.Int32ToString(int32(i)),
			)

			modelEvent := &model.SecretEvent{
				Type:   model.Initial,
				Secret: modelSecret,
			}

			event, err := g.converter.FromModelEvent(modelEvent)
			if err != nil {
				return logger.WrapError("convert model metric", err)
			}

			err = stream.Send(event)
			if err != nil {
				return logger.WrapError("send message", err)
			}
		}
	}
}
