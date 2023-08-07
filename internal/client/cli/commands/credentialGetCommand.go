package commands

import (
	"context"
	"fmt"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/storage"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model/secret"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

const (
	credentialGetCommandName      = "get"
	credentialGetShortDescription = "get credential secret"
	credentialGetFullDescription  = "Command get all stored credential,"
)

type credentialGetCommand struct {
	*baseCommand
	storage storage.ClientSecretsStorage
}

func NewCredentialGetCommand(stream io.CommandStream, storage storage.ClientSecretsStorage, children ...cli.Command) *credentialGetCommand {
	command := &credentialGetCommand{
		storage: storage,
	}
	command.baseCommand = newBaseCommand(
		stream,
		credentialGetCommandName,
		credentialGetShortDescription,
		credentialGetFullDescription,
		children,
		[]cli.Argument{
			newArgument("Identity", true, idFullArgName, idShortArgName),
			newArgument("Reveal secret values", false, revealFullArgName, revealShortArgName),
		},
		command.invoke,
	)
	return command
}

func (c *credentialGetCommand) invoke(ctx context.Context, args map[string]string) error {
	identity, ok := argValue(args, idFullArgName, idShortArgName)
	if !ok {
		return logger.WrapError(fmt.Sprintf("invoke %s command: secret identity is missed", c.name), cli.ErrRequiredArgNotFound)
	}

	credentialSecret, err := c.storage.GetSecretById(ctx, model.Credential, identity)
	if err != nil {
		return logger.WrapError("get credential secret", err)
	}

	_, reveal := argValue(args, revealFullArgName, revealShortArgName)
	credential := credentialSecret.(*secret.CredentialSecret)
	value := "***"
	if reveal {
		value = credential.Password
	}

	c.stream.Write(fmt.Sprintf("\t%s\t\t%s\t\t%s\t\t%s", credential.GetIdentity(), credential.UserName, value, credential.GetComment()))

	return nil
}
