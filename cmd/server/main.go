package main

import (
	"context"
	"flag"
	"os"
	"os/signal"

	"github.com/caarlos0/env/v7"

	modelGrpc "github.com/MaxReX92/go-yandex-gophkeeper/internal/model/grpc"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/serialization/json"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/server/grpc"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/server/storage/postgres"

	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/runner"
)

type config struct {
	ListenAddress            string `env:"LISTEN_ADDRESS" json:"listenAddress,omitempty"`
	PostgresConnectionString string `env:"DATABASE_DSN" json:"databaseDsn,omitempty"`
}

func main() {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// config
	conf, err := createConfig()
	if err != nil {
		panic(logger.WrapError("create config", err))
	}

	// interrupt
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// build app
	serializer := json.NewSerializer()
	converter := modelGrpc.NewConverter(serializer)
	dbStorage, err := postgres.NewDBStorage(ctx, conf)
	if err != nil {
		panic(logger.WrapError("create db storage", err))
	}
	server := grpc.NewGrpcServer(conf, dbStorage, converter)
	app := runner.NewGracefulRunner(server)

	// runtime
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

	flag.StringVar(&conf.ListenAddress, "l", "127.0.0.1:3200", "Server grpc URL")
	flag.StringVar(&conf.PostgresConnectionString, "d", "host=localhost user=postgres database=secrets password=postgres", "Database connection stirng")
	flag.Parse()

	err := env.Parse(conf)
	if err != nil {
		return nil, logger.WrapError("parse flags", err)
	}

	return conf, nil
}

func (c *config) GrpcAddress() string {
	return c.ListenAddress
}

func (c *config) ConnectionString() string {
	return c.PostgresConnectionString
}
