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
	noteGetCommandName      = "get"
	noteGetShortDescription = "get note secret"
	noteGetFullDescription  = "Command get all stored note."
)

type noteGetCommand struct {
	*baseCommand
	storage storage.ClientSecretsStorage
}

func NewNoteGetCommand(stream io.CommandStream, storage storage.ClientSecretsStorage, children ...cli.Command) *noteGetCommand {
	command := &noteGetCommand{
		storage: storage,
	}
	command.baseCommand = newBaseCommand(
		stream,
		noteGetCommandName,
		noteGetShortDescription,
		noteGetFullDescription,
		children,
		[]cli.Argument{
			newArgument("Identity", true, idFullArgName, idShortArgName),
			newArgument("Reveal secret values", false, revealFullArgName, revealShortArgName),
		},
		command.invoke,
	)
	return command
}

func (c *noteGetCommand) invoke(ctx context.Context, args map[string]string) error {
	identity, ok := argValue(args, idFullArgName, idShortArgName)
	if !ok {
		return logger.WrapError(fmt.Sprintf("invoke %s command: secret identity is missed", c.name), cli.ErrRequiredArgNotFound)
	}

	noteSecret, err := c.storage.GetSecretByID(ctx, model.Note, identity)
	if err != nil {
		return logger.WrapError("get secrets", err)
	}

	_, reveal := argValue(args, revealFullArgName, revealShortArgName)
	note := noteSecret.(*secret.NoteSecret)
	value := hiddenValue
	if reveal {
		value = note.Text
	}

	c.stream.Write(fmt.Sprintf("\t%s\t\t%s\t\t%s", note.GetIdentity(), value, note.GetComment()))

	return nil
}
