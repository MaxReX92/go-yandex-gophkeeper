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
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/storage/remote"
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
	remoteStorage := remote.NewStorage()
	clientStorage := storage.NewStorageStrategy(memoryStorage, remoteStorage)
	initialCommand := buildCommands(ioStream, randomGenerator, clientStorage)
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
