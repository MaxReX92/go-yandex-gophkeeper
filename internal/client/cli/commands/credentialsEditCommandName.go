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
	credentialsEditCommandName      = "edit"
	credentialsEditShortDescription = "edit credential from store"
	credentialsEditFullDescription  = "Edit credential from secure store,"
)

type credentialsEditCommand struct {
	*baseCommand
	storage storage.LocalSecretsStorage
}

func NewCredentialsEditCommand(
	stream io.CommandStream,
	storage storage.LocalSecretsStorage,
	children ...cli.Command,
) *credentialsEditCommand {
	command := &credentialsEditCommand{
		storage: storage,
	}
	command.baseCommand = newBaseCommand(
		stream,
		credentialsEditCommandName,
		credentialsEditShortDescription,
		credentialsEditFullDescription,
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

func (c *credentialsEditCommand) invoke(args map[string]string) error {
	identity, ok := argValue(args, idFullArgName, idShortArgName)
	if !ok {
		return logger.WrapError(fmt.Sprintf("invoke %s command: secret identity is missed", c.name), cli.ErrRequiredArgNotFound)
	}

	currentCred, err := c.storage.GetSecretById(model.Credentials, identity)
	if err != nil {
		return logger.WrapError("get secret", err)
	}

	if currentCred == nil {
		return logger.WrapError("edit secret", cli.ErrSecretNotFound)
	}

	cred, ok := currentCred.(*secret.CredentialsSecret)
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
	err = c.storage.ChangeSecret(currentCred)
	if err != nil {
		return logger.WrapError("edit secret", err)
	}

	return nil
}
