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
	credentialListCommandName      = "list"
	credentialListShortDescription = "list of all credential"
	credentialListFullDescription  = "Command list all stored credential,"
)

type credentialListCommand struct {
	*baseCommand
	storage storage.ClientSecretsStorage
}

// NewCredentialListCommand creates a new instance of list credentials secret command.
func NewCredentialListCommand(stream io.CommandStream, storage storage.ClientSecretsStorage, children ...cli.Command) *credentialListCommand {
	command := &credentialListCommand{
		storage: storage,
	}
	command.baseCommand = newBaseCommand(
		stream,
		credentialListCommandName,
		credentialListShortDescription,
		credentialListFullDescription,
		children,
		[]cli.Argument{
			newArgument("Reveal secret values", false, revealFullArgName, revealShortArgName),
		},
		command.invoke,
	)
	return command
}

func (c *credentialListCommand) invoke(ctx context.Context, args map[string]string) error {
	credentials, err := c.storage.GetAllSecrets(ctx, model.Credential)
	if err != nil {
		return logger.WrapError("get secrets", err)
	}

	_, reveal := argValue(args, revealFullArgName, revealShortArgName)
	for _, modelCredential := range credentials {
		cred := modelCredential.(*secret.CredentialSecret)
		value := hiddenValue
		if reveal {
			value = cred.Password
		}

		c.stream.Write(fmt.Sprintf("\t%s\t\t%s\t\t%s\t\t%s\n", cred.GetIdentity(), cred.UserName, value, cred.GetComment()))
	}

	return nil
}
