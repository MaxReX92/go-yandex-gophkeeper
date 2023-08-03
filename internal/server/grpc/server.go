package grpc

import (
	"context"
	"net"
	"time"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/generated"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
	rpc "google.golang.org/grpc"
)

type GrpcServerConfig interface {
	ListenTCP() string
}

type grpcServer struct {
	generated.UnimplementedSecretServiceServer

	listenTCP string
	server    *rpc.Server
}

func NewGrpcServer(conf GrpcServerConfig) *grpcServer {
	return &grpcServer{
		listenTCP: conf.ListenTCP(),
		server:    rpc.NewServer(),
	}
}

func (g *grpcServer) Start(_ context.Context) error {
	listen, err := net.Listen("tcp", g.listenTCP)
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

func (g *grpcServer) SecretEvents(user *generated.User, stream generated.SecretService_SecretEventsServer) error {
	ticker := time.NewTicker(1 * time.Second)
	for i := 0; ; i++ {
		select {
		case <-stream.Context().Done():
			break
		case <-ticker.C:
			err := stream.Send(&generated.SecretEvent{
				Type: generated.EventType_INITIAL,
				Secret: &generated.Secret{
					Identity: "12345",
					Type:     generated.SecretType_CARD,
					Content:  []byte("cha cha cha"),
				},
			})

			if err != nil {
				return logger.WrapError("send message", err)
			}
		}
	}
}
