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
	loginCommandName      = "login"
	loginShortDescription = "login user"
	loginFullDescription  = "Login user in secure store system"
)

type loginCommand struct {
	*baseCommand
	credentialsProvider auth.CredentialsProvider
}

func NewLoginCommand(
	stream io.CommandStream,
	credentialsProvider auth.CredentialsProvider,
	children ...cli.Command,
) *loginCommand {
	command := &loginCommand{
		credentialsProvider: credentialsProvider,
	}
	command.baseCommand = newBaseCommand(
		stream,
		loginCommandName,
		loginShortDescription,
		loginFullDescription,
		children,
		[]cli.Argument{
			newArgument("User name", true, userFullArgName, userShortArgName),
			newArgument("Password", true, passFullArgName, passShortArgName),
		},
		command.invoke,
	)
	return command
}

func (c *loginCommand) invoke(ctx context.Context, args map[string]string) error {
	userName, ok := argValue(args, userFullArgName, userShortArgName)
	if !ok {
		return logger.WrapError(fmt.Sprintf("invoke %s command: user name is missed", c.name), cli.ErrRequiredArgNotFound)
	}

	password, ok := argValue(args, passFullArgName, passShortArgName)
	if !ok {
		return logger.WrapError(fmt.Sprintf("invoke %s command: password is missed", c.name), cli.ErrRequiredArgNotFound)
	}

	err := c.credentialsProvider.Login(ctx, userName, password)
	if err != nil {
		return logger.WrapError("login user", err)
	}

	return nil
}
