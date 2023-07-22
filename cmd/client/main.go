package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli/commands"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
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
	initialCommand := buildCommands(ioStream)
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

func buildCommands(ioStream io.CommandStream) cli.Command {
	// credentials
	credentialsListCommand := commands.NewCredentialsListCommand(ioStream, commands.NewHelpCommand())
	credentialsCommand := commands.NewCredentialsCommand(ioStream, commands.NewHelpCommand(), credentialsListCommand)

	// status
	statucCommand := commands.NewStatusCommand(ioStream, commands.NewHelpCommand())

	return commands.NewInitialCommand(ioStream, commands.NewHelpCommand(), credentialsCommand, statucCommand)
}
