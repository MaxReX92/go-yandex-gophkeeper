package main

import (
	"context"
	"encoding/json"
	"os"
	"os/signal"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/generator/rand"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/tls/cert"
	"github.com/caarlos0/env/v7"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/auth"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli/commands"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/service/grpc"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/storage"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/storage/memory"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/storage/remote"
	modelGrpc "github.com/MaxReX92/go-yandex-gophkeeper/internal/model/grpc"
	internalJson "github.com/MaxReX92/go-yandex-gophkeeper/internal/serialization/json"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/runner"
)

type config struct {
	ConfigPath     string `env:"CONFIG"`
	LogsPath       string `env:"LOGS_PATH" envDefault:"./log.txt" json:"logsPath,omitempty"`
	GrpcAddress    string `env:"GRPC_ADDRESS" envDefault:"127.0.0.1:3200" json:"grpcAddress,omitempty"`
	IdentityLen    int32  `env:"IDENTITY_LEN" envDefault:"8" json:"identityLen,omitempty"`
	PublicCertPath string `env:"CERT_PATH" envDefault:"../../credentials/public.crt" json:"certPath,omitempty"`
	PrivateKeyPath string `env:"KEY_PATH" envDefault:"../../credentials/private.key" json:"keyPath,omitempty"`
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

	// init logger
	logFile, err := os.OpenFile(conf.LogsPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	if err != nil {
		panic(logger.WrapError("error opening log file", err))
	}
	defer logFile.Close()
	logger.SetOutput(logFile)

	// interrupt
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// build app
	ioStream := io.NewIOStream(os.Stdin, os.Stdout)
	randomGenerator := rand.NewGenerator(conf)
	serializer := internalJson.NewSerializer()
	converter := modelGrpc.NewConverter(serializer)
	credentials := auth.NewCredentials("test_user")
	tlsProvider := cert.NewTLSProvider(conf)
	service, err := grpc.NewService(conf, credentials, converter, tlsProvider)
	if err != nil {
		panic(logger.WrapError("create grpc service", err))
	}
	memoryStorage := memory.NewStorage()
	remoteStorage := remote.NewStorage(service)
	supervisor := storage.NewStorageSupervisor(service, memoryStorage)
	clientStorage := storage.NewStorageStrategy(memoryStorage, remoteStorage)
	initialCommand := buildCommands(ioStream, randomGenerator, clientStorage)
	handler := cli.NewHandler(ioStream, initialCommand)
	multiRunner := runner.NewMultiWorker(
		supervisor,
		handler,
	)
	app := runner.NewGracefulRunner(multiRunner)

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
	err := env.Parse(conf)
	if err != nil {
		return nil, logger.WrapError("parse flags", err)
	}
	if conf.ConfigPath != "" {
		content, err := os.ReadFile(conf.ConfigPath)
		if err != nil {
			return nil, logger.WrapError("read json config file", err)
		}

		err = json.Unmarshal(content, conf)
		if err != nil {
			return nil, logger.WrapError("unmarshal json config file", err)
		}
	}

	return conf, nil
}

func buildCommands(
	ioStream io.CommandStream,
	generator identity.Generator,
	storage storage.ClientSecretsStorage,
) cli.Command {
	// binary
	binaryAddCommand := commands.NewBinaryAddCommand(ioStream, generator, storage, commands.NewHelpCommand())
	binaryEditCommand := commands.NewBinaryEditCommand(ioStream, storage, commands.NewHelpCommand())
	binaryGetCommand := commands.NewBinaryGetCommand(ioStream, storage, commands.NewHelpCommand())
	binaryListCommand := commands.NewBinaryListCommand(ioStream, storage, commands.NewHelpCommand())
	binaryRemoveCommand := commands.NewBinaryRemoveCommand(ioStream, storage, commands.NewHelpCommand())
	binaryCommand := commands.NewBinaryCommand(
		ioStream,
		commands.NewHelpCommand(),
		binaryAddCommand,
		binaryEditCommand,
		binaryGetCommand,
		binaryListCommand,
		binaryRemoveCommand,
	)

	// card
	cardAddCommand := commands.NewCardAddCommand(ioStream, generator, storage, commands.NewHelpCommand())
	cardEditCommand := commands.NewCardEditCommand(ioStream, storage, commands.NewHelpCommand())
	cardGetCommand := commands.NewCardGetCommand(ioStream, storage, commands.NewHelpCommand())
	cardListCommand := commands.NewCardListCommand(ioStream, storage, commands.NewHelpCommand())
	cardRemoveCommand := commands.NewCardRemoveCommand(ioStream, storage, commands.NewHelpCommand())
	cardCommand := commands.NewCardCommand(
		ioStream,
		commands.NewHelpCommand(),
		cardAddCommand,
		cardEditCommand,
		cardGetCommand,
		cardListCommand,
		cardRemoveCommand,
	)

	// credential
	credentialAddCommand := commands.NewCredentialAddCommand(ioStream, generator, storage, commands.NewHelpCommand())
	credentialEditCommand := commands.NewCredentialEditCommand(ioStream, storage, commands.NewHelpCommand())
	credentialGetCommand := commands.NewCredentialGetCommand(ioStream, storage, commands.NewHelpCommand())
	credentialListCommand := commands.NewCredentialListCommand(ioStream, storage, commands.NewHelpCommand())
	credentialRemoveCommand := commands.NewCredentialRemoveCommand(ioStream, storage, commands.NewHelpCommand())
	credentialCommand := commands.NewCredentialCommand(
		ioStream,
		commands.NewHelpCommand(),
		credentialAddCommand,
		credentialEditCommand,
		credentialGetCommand,
		credentialListCommand,
		credentialRemoveCommand,
	)

	// note
	noteAddCommand := commands.NewNoteAddCommand(ioStream, generator, storage, commands.NewHelpCommand())
	noteEditCommand := commands.NewNoteEditCommand(ioStream, storage, commands.NewHelpCommand())
	noteGetCommand := commands.NewNoteGetCommand(ioStream, storage, commands.NewHelpCommand())
	noteListCommand := commands.NewNoteListCommand(ioStream, storage, commands.NewHelpCommand())
	noteRemoveCommand := commands.NewNoteRemoveCommand(ioStream, storage, commands.NewHelpCommand())
	noteCommand := commands.NewNoteCommand(
		ioStream,
		commands.NewHelpCommand(),
		noteAddCommand,
		noteEditCommand,
		noteGetCommand,
		noteListCommand,
		noteRemoveCommand,
	)

	// status
	statusCommand := commands.NewStatusCommand(ioStream, commands.NewHelpCommand())

	return commands.NewInitialCommand(
		ioStream,
		commands.NewHelpCommand(),
		binaryCommand,
		cardCommand,
		credentialCommand,
		noteCommand,
		statusCommand,
	)
}

func (c *config) GrpcServerAddress() string {
	return c.GrpcAddress
}

func (c *config) IdentityLength() int32 {
	return c.IdentityLen
}

func (c *config) GetPublicCertPath() string {
	return c.PublicCertPath
}

func (c *config) GetPrivateKeyPath() string {
	return c.PrivateKeyPath
}
