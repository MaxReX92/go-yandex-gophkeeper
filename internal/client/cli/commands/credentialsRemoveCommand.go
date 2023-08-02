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
	credentialRemoveCommandName      = "remove"
	credentialRemoveShortDescription = "remove credential from store"
	credentialRemoveFullDescription  = "Remove credential from secure store,"
)

type credentialRemoveCommand struct {
	*baseCommand
	storage storage.LocalSecretsStorage
}

func NewCredentialRemoveCommand(
	stream io.CommandStream,
	storage storage.LocalSecretsStorage,
	children ...cli.Command,
) *credentialRemoveCommand {
	command := &credentialRemoveCommand{
		storage: storage,
	}
	command.baseCommand = newBaseCommand(
		stream,
		credentialRemoveCommandName,
		credentialRemoveShortDescription,
		credentialRemoveFullDescription,
		children,
		[]cli.Argument{
			newArgument("Remove all credential secrets", false, allFullArgName),
			newArgument("Secret identity", true, idFullArgName, idShortArgName),
		},
		command.invoke,
	)
	return command
}

func (c *credentialRemoveCommand) invoke(args map[string]string) error {
	var toRemove []*secret.CredentialSecret
	_, removeAll := argValue(args, allFullArgName)
	if removeAll {
		credential, err := c.storage.GetAllSecrets(model.Credential)
		if err != nil {
			return logger.WrapError("get all secrets", err)
		}

		for _, cred := range credential {
			toRemove = append(toRemove, cred.(*secret.CredentialSecret))
		}

	} else {
		identity, ok := argValue(args, idShortArgName, idFullArgName)
		if !ok {
			return logger.WrapError(fmt.Sprintf("invoke %s command: secret identity is missed", c.name), cli.ErrRequiredArgNotFound)
		}

		credential, err := c.storage.GetSecretById(model.Credential, identity)
		if err != nil {
			return logger.WrapError("get secret", err)
		}

		toRemove = append(toRemove, credential.(*secret.CredentialSecret))
	}

	for _, cred := range toRemove {
		logger.InfoFormat("Remove %s %s credential", cred.GetIdentity(), cred.UserName)
		err := c.storage.RemoveSecret(cred)
		if err != nil {
			return logger.WrapError("get secret", err)
		}
	}

	return nil
}
