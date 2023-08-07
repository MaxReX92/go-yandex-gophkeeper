package main

import (
	"context"
	"flag"
	"os"
	"os/signal"

	modelGrpc "github.com/MaxReX92/go-yandex-gophkeeper/internal/model/grpc"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/serialization/json"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/server/grpc"
	"github.com/caarlos0/env/v7"

	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/runner"
)

type config struct {
	ListenAddress string `env:"LISTEN_ADDRESS" json:"listen_address,omitempty"`
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
	serializer := json.NewSerializer()
	converter := modelGrpc.NewConverter(serializer)
	server := grpc.NewGrpcServer(conf, converter)
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
