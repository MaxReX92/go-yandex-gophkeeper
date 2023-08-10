package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"time"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/auth/server/grpc"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/auth/token/jwt"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/db/postgres"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/identity/rand"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/tls/cert"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/runner"
	"github.com/caarlos0/env/v7"
)

type config struct {
	ListenAddress            string        `env:"LISTEN_ADDRESS"`
	PostgresConnectionString string        `env:"DATABASE_DSN"`
	PublicCertPath           string        `env:"CERT_PATH"`
	PrivateKeyPath           string        `env:"KEY_PATH"`
	TokenSecretKey           string        `env:"SECRET_KEY"`
	TokensTTL                time.Duration `env:"TOKEN_TTL"`
	IdentityLen              int           `env:"IDENTITY_LEN"`
}

func main() {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// config
	conf, err := createConfig()
	if err != nil {
		panic(logger.WrapError("create config file", err))
	}

	// interrupt
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// build app
	dbService, err := postgres.NewDBService(ctx, conf)
	if err != nil {
		panic(logger.WrapError("create db storage", err))
	}
	identityGenerator := rand.NewGenerator(conf)
	tlsProvider := cert.NewTLSProvider(conf)
	tokenGenerator := jwt.NewGenerator(conf)
	tokenValidator := jwt.NewValidator(conf)

	server := grpc.NewServer(conf, dbService, tlsProvider, tokenGenerator, tokenValidator, identityGenerator)
	app := runner.NewGracefulRunner(server)

	// app runtime
	app.Start(ctx)

	// shutdown
	select {
	case err = <-app.Error():
		err = logger.WrapError("start application", err)
	case <-interrupt:
		err = app.Stop(ctx)
	}

	if err != nil {
		logger.ErrorObj(err)
	}
}

func createConfig() (*config, error) {
	conf := &config{}

	flag.StringVar(&conf.ListenAddress, "l", "127.0.0.1:3201", "Server grpc URL")
	flag.StringVar(&conf.PostgresConnectionString, "d", "host=localhost user=postgres database=secrets password=postgres", "Database connection stirng")
	flag.StringVar(&conf.PublicCertPath, "c", "../../credentials/public.crt", "Path to public cert")
	flag.StringVar(&conf.PrivateKeyPath, "k", "../../credentials/private.key", "Path to private key")
	flag.StringVar(&conf.TokenSecretKey, "s", "totalSecret", "Token secret key")
	flag.IntVar(&conf.IdentityLen, "i", 16, "Identity generator length")
	flag.DurationVar(&conf.TokensTTL, "t", time.Second*60, "Store backup interval")
	flag.Parse()

	err := env.Parse(conf)
	return conf, err
}

func (c *config) GrpcAddress() string {
	return c.ListenAddress
}

func (c *config) ConnectionString() string {
	return c.PostgresConnectionString
}

func (c *config) SecretKey() []byte {
	return []byte(c.TokenSecretKey)
}

func (c *config) TokenTTL() time.Duration {
	return c.TokensTTL
}

func (c *config) IdentityLength() int32 {
	return int32(c.IdentityLen)
}
func (c *config) GetPublicCertPath() string {
	return c.PublicCertPath
}

func (c *config) GetPrivateKeyPath() string {
	return c.PrivateKeyPath
}
