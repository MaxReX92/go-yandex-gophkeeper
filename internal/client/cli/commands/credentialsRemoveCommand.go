package commands

import (
	"fmt"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/storage"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model/secret"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

const (
	credentialsRemoveCommandName      = "remove"
	credentialsRemoveShortDescription = "remove credential from store"
	credentialsRemoveFullDescription  = "Remove credential from secure store,"
)

type credentialsRemoveCommand struct {
	*baseCommand
	storage storage.LocalSecretsStorage
}

func NewCredentialsRemoveCommand(
	stream io.CommandStream,
	storage storage.LocalSecretsStorage,
	children ...cli.Command,
) *credentialsRemoveCommand {
	command := &credentialsRemoveCommand{
		storage: storage,
	}
	command.baseCommand = newBaseCommand(
		stream,
		credentialsRemoveCommandName,
		credentialsRemoveShortDescription,
		credentialsRemoveFullDescription,
		children,
		[]cli.Argument{
			newArgument("Remove all credential secrets", false, allFullArgName),
			newArgument("Secret identity", true, idFullArgName, idShortArgName),
		},
		command.invoke,
	)
	return command
}

func (c *credentialsRemoveCommand) invoke(args map[string]string) error {
	var toRemove []*secret.CredentialsSecret
	_, removeAll := argValue(args, allFullArgName)
	if removeAll {
		credentials, err := c.storage.GetAllSecrets(model.Credentials)
		if err != nil {
			return logger.WrapError("get all secrets", err)
		}

		for _, cred := range credentials {
			toRemove = append(toRemove, cred.(*secret.CredentialsSecret))
		}

	} else {
		identity, ok := argValue(args, idShortArgName, idFullArgName)
		if !ok {
			return logger.WrapError(fmt.Sprintf("invoke %s command: secret identity is missed", c.name), cli.ErrRequiredArgNotFound)
		}

		credentials, err := c.storage.GetSecretById(model.Credentials, identity)
		if err != nil {
			return logger.WrapError("get secret", err)
		}

		toRemove = append(toRemove, credentials.(*secret.CredentialsSecret))
	}

	for _, cred := range toRemove {
		logger.InfoFormat("Remove %s %s credential", cred.UserName, cred.Password)
		err := c.storage.RemoveSecret(cred)
		if err != nil {
			return logger.WrapError("get secret", err)
		}
	}

	return nil
}
