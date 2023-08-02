package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli/commands"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/generator"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/generator/rand"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/storage"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/storage/memory"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/runner"
)

func main() {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// config
	//conf, err := createConfig()
	//if err != nil {
	//	panic(logger.WrapError("create config file", err))
	//}

	// interrupt
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// build app
	ioStream := io.NewIOStream(os.Stdin, os.Stdout)
	randomGenerator := rand.NewGenerator()
	memoryStorage := memory.NewStorage()

	initialCommand := buildCommands(ioStream, randomGenerator, memoryStorage)
	handler := cli.NewHandler(ioStream, initialCommand)
	app := runner.NewGracefulRunner(handler)

	// runtime
	app.Start(ctx)

	// shutdown
	var err error
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

func buildCommands(
	ioStream io.CommandStream,
	generator generator.Generator,
	storage storage.LocalSecretsStorage,
) cli.Command {
	// credential
	credentialAddCommand := commands.NewCredentialAddCommand(ioStream, generator, storage, commands.NewHelpCommand())
	credentialListCommand := commands.NewCredentialListCommand(ioStream, storage, commands.NewHelpCommand())
	credentialRemoveCommand := commands.NewCredentialRemoveCommand(ioStream, storage, commands.NewHelpCommand())
	credentialCommand := commands.NewCredentialCommand(
		ioStream,
		commands.NewHelpCommand(),
		credentialAddCommand,
		credentialListCommand,
		credentialRemoveCommand,
	)

	// status
	statusCommand := commands.NewStatusCommand(ioStream, commands.NewHelpCommand())

	return commands.NewInitialCommand(ioStream, commands.NewHelpCommand(), credentialCommand, statusCommand)
}
