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
	credentialsListCommandName      = "list"
	credentialsListShortDescription = "list of all credentials"
	credentialsListFullDescription  = "Command list all stored credentials,"
)

type credentialsListCommand struct {
	*baseCommand
	storage storage.LocalSecretsStorage
}

func NewCredentialsListCommand(stream io.CommandStream, storage storage.LocalSecretsStorage, children ...cli.Command) *credentialsListCommand {
	command := &credentialsListCommand{
		storage: storage,
	}
	command.baseCommand = newBaseCommand(
		stream,
		credentialsListCommandName,
		credentialsListShortDescription,
		credentialsListFullDescription,
		children,
		[]cli.Argument{
			newArgument("Reveal secret values", false, revealFullArgName, revealShortArgName),
			newArgument("Show full secret info", false, verboseFullArgName, verboseShortArgName),
		},
		command.invoke,
	)
	return command
}

func (c *credentialsListCommand) invoke(args map[string]string) error {
	_, reveal := argValue(args, revealFullArgName, revealShortArgName)

	credentials, err := c.storage.GetAllSecrets(model.Credentials)
	if err != nil {
		return logger.WrapError("get secrets", err)
	}

	for _, modelCredentials := range credentials {
		cred := modelCredentials.(*secret.CredentialsSecret)
		value := "***"
		if reveal {
			value = cred.Password
		}

		c.stream.Write(fmt.Sprintf("\t%s\t\t%s\t\t%s\t\t%s", cred.GetIdentity(), cred.UserName, value, cred.GetComment()))
	}

	return nil
}
