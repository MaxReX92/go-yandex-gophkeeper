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
	credentialEditCommandName      = "edit"
	credentialEditShortDescription = "edit credential from store"
	credentialEditFullDescription  = "Edit credential from secure store,"
)

type credentialEditCommand struct {
	*baseCommand
	storage storage.ClientSecretsStorage
}

func NewCredentialEditCommand(
	stream io.CommandStream,
	storage storage.ClientSecretsStorage,
	children ...cli.Command,
) *credentialEditCommand {
	command := &credentialEditCommand{
		storage: storage,
	}
	command.baseCommand = newBaseCommand(
		stream,
		credentialEditCommandName,
		credentialEditShortDescription,
		credentialEditFullDescription,
		children,
		[]cli.Argument{
			newArgument("Secret identity", true, initialFullDescription, initialShortDescription),
			newArgument("User name", true, userFullArgName, userShortArgName),
			newArgument("Password", true, passFullArgName, passShortArgName),
			newArgument("Comment", true, commentFullArgName, commentShortArgName),
		},
		command.invoke,
	)
	return command
}

func (c *credentialEditCommand) invoke(args map[string]string) error {
	identity, ok := argValue(args, idFullArgName, idShortArgName)
	if !ok {
		return logger.WrapError(fmt.Sprintf("invoke %s command: secret identity is missed", c.name), cli.ErrRequiredArgNotFound)
	}

	currentCred, err := c.storage.GetSecretById(nil, model.Credential, identity)
	if err != nil {
		return logger.WrapError("get secret", err)
	}

	if currentCred == nil {
		return logger.WrapError("edit secret", cli.ErrSecretNotFound)
	}

	cred, ok := currentCred.(*secret.CredentialSecret)
	if !ok {
		return logger.WrapError("edit secret", cli.ErrInvalidSecretType)
	}

	userName, ok := argValue(args, userFullArgName, userShortArgName)
	if ok {
		cred.UserName = userName
	}

	password, ok := argValue(args, passFullArgName, passShortArgName)
	if ok {
		cred.Password = password
	}

	comment, ok := argValue(args, commentFullArgName, commentShortArgName)
	if ok {
		cred.Comment = comment
	}

	logger.InfoFormat("Edit %s %s credential", userName, password)
	err = c.storage.ChangeSecret(nil, currentCred)
	if err != nil {
		return logger.WrapError("edit secret", err)
	}

	return nil
}
