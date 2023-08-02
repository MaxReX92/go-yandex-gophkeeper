package commands

import (
	"fmt"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/generator"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/storage"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model/secret"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

const (
	credentialsAddCommandName      = "add"
	credentialsAddShortDescription = "add credential to store"
	credentialsAddFullDescription  = "Add new credential to secure store,"
)

type credentialsAddCommand struct {
	*baseCommand
	generator generator.Generator
	storage   storage.LocalSecretsStorage
}

func NewCredentialsAddCommand(
	stream io.CommandStream,
	generator generator.Generator,
	storage storage.LocalSecretsStorage,
	children ...cli.Command,
) *credentialsAddCommand {
	command := &credentialsAddCommand{
		generator: generator,
		storage:   storage,
	}
	command.baseCommand = newBaseCommand(
		stream,
		credentialsAddCommandName,
		credentialsAddShortDescription,
		credentialsAddFullDescription,
		children,
		[]cli.Argument{
			newArgument("User name", true, userFullArgName, userShortArgName),
			newArgument("Password", true, passFullArgName, passShortArgName),
			newArgument("Comment", true, commentFullArgName, commentShortArgName),
		},
		command.invoke,
	)
	return command
}

func (c *credentialsAddCommand) invoke(args map[string]string) error {
	userName, ok := argValue(args, userFullArgName, userShortArgName)
	if !ok {
		return logger.WrapError(fmt.Sprintf("invoke %s command: user name is missed", c.name), cli.ErrRequiredArgNotFound)
	}

	password, ok := argValue(args, passFullArgName, passShortArgName)
	if !ok {
		return logger.WrapError(fmt.Sprintf("invoke %s command: password is missed", c.name), cli.ErrRequiredArgNotFound)
	}

	comment, _ := argValue(args, commentFullArgName, commentShortArgName)
	cred := secret.NewCredentialsSecret(userName, password, c.generator.GenerateNewIdentity(), comment)

	logger.InfoFormat("Add %s %s credential", userName, password)
	err := c.storage.AddSecret(cred)
	if err != nil {
		return logger.WrapError("add secret", err)
	}

	return nil
}
