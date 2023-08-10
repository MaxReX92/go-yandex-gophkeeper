package grpc

import (
	"context"
	"crypto/tls"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/auth/server"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/auth/token"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/db"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/generated"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/identity"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
	tlsCert "github.com/MaxReX92/go-yandex-gophkeeper/internal/tls"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
	rpc "google.golang.org/grpc"
)

type GrpcServerConfig interface {
	GrpcAddress() string
}

type grpcServer struct {
	generated.UnimplementedAuthServiceServer

	listenAddress     string
	server            *rpc.Server
	dbService         db.Service
	tlsProvider       tlsCert.TLSProvider
	tokenGenerator    token.Generator
	tokenValidator    token.Validator
	identityGenerator identity.Generator
}

func NewGrpcServer(
	conf GrpcServerConfig,
	dbService db.Service,
	tlsProvider tlsCert.TLSProvider,
	tokenGenerator token.Generator,
	tokenValidator token.Validator,
	identityGenerator identity.Generator,
) *grpcServer {
	return &grpcServer{
		listenAddress:     conf.GrpcAddress(),
		dbService:         dbService,
		server:            rpc.NewServer(),
		tlsProvider:       tlsProvider,
		tokenGenerator:    tokenGenerator,
		tokenValidator:    tokenValidator,
		identityGenerator: identityGenerator,
	}
}

func (g *grpcServer) Start(_ context.Context) error {
	tlsConfig, err := g.tlsProvider.GetTlsConfig()
	if err != nil {
		return logger.WrapError("create tls config", err)
	}

	listen, err := tls.Listen("tcp", g.listenAddress, tlsConfig)
	if err != nil {
		return logger.WrapError("start listen TCP", err)
	}
	generated.RegisterAuthServiceServer(g.server, g)

	logger.Info("Start gRPC service")
	err = g.server.Serve(listen)
	if err != nil {
		return logger.WrapError("start gRPC service", err)
	}

	return nil
}

func (g *grpcServer) Stop(_ context.Context) error {
	logger.Info("Stopping gRPC service")
	g.server.GracefulStop()
	return nil
}

func (g *grpcServer) Register(ctx context.Context, request *generated.RegisterRequest) (*generated.RegisterResponse, error) {
	if request.Name == "" {
		return nil, logger.WrapError("register user", server.ErrInvalidRequest)
	}

	if request.Password == "" {
		return nil, logger.WrapError("register user", server.ErrInvalidRequest)
	}

	userId := g.identityGenerator.GenerateNewIdentity()
	personalToken := g.identityGenerator.GenerateNewIdentity()
	err := g.dbService.CallInTransaction(ctx, func(ctx context.Context, executor db.Executor) error {
		err := executor.AddUser(ctx, userId, request.Name, request.Password, personalToken)
		if err != nil {
			return logger.WrapError("call new user query", err)
		}

		return nil
	})

	if err != nil {
		return nil, logger.WrapError("create new user", err)
	}

	return &generated.RegisterResponse{Identity: userId}, nil
}

func (g *grpcServer) Login(ctx context.Context, request *generated.LoginRequest) (*generated.LoginResponse, error) {
	if request.Name == "" {
		return nil, logger.WrapError("login user", server.ErrInvalidRequest)
	}

	if request.Password == "" {
		return nil, logger.WrapError("login user", server.ErrInvalidRequest)
	}

	user, err := g.getUser(ctx, request.Name)
	if err != nil {
		return nil, logger.WrapError("validate credentials", err)
	}

	if user.Password != request.Password {
		return nil, logger.WrapError("login user", server.ErrInvalidCredentials)
	}

	newToken, err := g.tokenGenerator.GenerateToken()
	if err != nil {
		return nil, logger.WrapError("generate new token", err)
	}

	return &generated.LoginResponse{
		Identity:      user.Identity,
		Token:         newToken,
		PersonalToken: user.PersonalToken,
	}, nil
}

func (g *grpcServer) Prolong(ctx context.Context, request *generated.ProlongTokenRequest) (*generated.ProlongTokenResponse, error) {
	if request.Token == "" {
		return nil, logger.WrapError("login user", server.ErrInvalidRequest)
	}

	ok, err := g.tokenValidator.Check(request.Token)
	if err != nil {
		return nil, logger.WrapError("check token", err)
	}

	if !ok {
		return nil, logger.WrapError("validate token", server.ErrInvalidCredentials)
	}

	newToken, err := g.tokenGenerator.GenerateToken()
	if err != nil {
		return nil, logger.WrapError("generate new token", err)
	}

	return &generated.ProlongTokenResponse{Token: newToken}, err
}

func (g *grpcServer) getUser(ctx context.Context, loginName string) (*model.User, error) {
	result, err := g.dbService.CallInTransactionResult(ctx, func(ctx context.Context, executor db.Executor) ([]any, error) {
		loginUser, err := executor.GetUserByUserName(ctx, loginName)
		if err != nil {
			return nil, logger.WrapError("call users db", err)
		}

		if loginUser == nil {
			return nil, logger.WrapError("get user", server.ErrLoginNotFound)
		}

		return []any{loginUser}, nil
	})
	if err != nil {
		return nil, logger.WrapError("get user info", err)
	}

	return result[0].(*model.User), nil
}
