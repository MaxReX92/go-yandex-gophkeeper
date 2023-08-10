package commands

import (
	"context"
	"fmt"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/auth"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

const (
	registerCommandName      = "register"
	registerShortDescription = "register user"
	registerFullDescription  = "Register new user in secure store system"
)

type registerCommand struct {
	*baseCommand
	credentialsProvider auth.CredentialsProvider
}

func NewRegisterCommand(
	stream io.CommandStream,
	credentialsProvider auth.CredentialsProvider,
	children ...cli.Command,
) *registerCommand {
	command := &registerCommand{
		credentialsProvider: credentialsProvider,
	}
	command.baseCommand = newBaseCommand(
		stream,
		registerCommandName,
		registerShortDescription,
		registerFullDescription,
		children,
		[]cli.Argument{
			newArgument("User name", true, userFullArgName, userShortArgName),
			newArgument("Password", true, passFullArgName, passShortArgName),
		},
		command.invoke,
	)
	return command
}

func (c *registerCommand) invoke(ctx context.Context, args map[string]string) error {
	userName, ok := argValue(args, userFullArgName, userShortArgName)
	if !ok {
		return logger.WrapError(fmt.Sprintf("invoke %s command: user name is missed", c.name), cli.ErrRequiredArgNotFound)
	}

	password, ok := argValue(args, passFullArgName, passShortArgName)
	if !ok {
		return logger.WrapError(fmt.Sprintf("invoke %s command: password is missed", c.name), cli.ErrRequiredArgNotFound)
	}

	err := c.credentialsProvider.Register(ctx, userName, password)
	if err != nil {
		return logger.WrapError("register new user", err)
	}

	return nil
}
