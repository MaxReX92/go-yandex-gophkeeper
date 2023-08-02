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
	noteRemoveCommandName      = "remove"
	noteRemoveShortDescription = "remove note from store"
	noteRemoveFullDescription  = "Remove note from secure store,"
)

type noteRemoveCommand struct {
	*baseCommand
	storage storage.LocalSecretsStorage
}

func NewNoteRemoveCommand(
	stream io.CommandStream,
	storage storage.LocalSecretsStorage,
	children ...cli.Command,
) *noteRemoveCommand {
	command := &noteRemoveCommand{
		storage: storage,
	}
	command.baseCommand = newBaseCommand(
		stream,
		noteRemoveCommandName,
		noteRemoveShortDescription,
		noteRemoveFullDescription,
		children,
		[]cli.Argument{
			newArgument("Remove all note secrets", false, allFullArgName),
			newArgument("Secret identity", true, idFullArgName, idShortArgName),
		},
		command.invoke,
	)
	return command
}

func (c *noteRemoveCommand) invoke(args map[string]string) error {
	var toRemove []*secret.NoteSecret
	_, removeAll := argValue(args, allFullArgName)
	if removeAll {
		note, err := c.storage.GetAllSecrets(model.Note)
		if err != nil {
			return logger.WrapError("get all secrets", err)
		}

		for _, note := range note {
			toRemove = append(toRemove, note.(*secret.NoteSecret))
		}

	} else {
		identity, ok := argValue(args, idShortArgName, idFullArgName)
		if !ok {
			return logger.WrapError(fmt.Sprintf("invoke %s command: secret identity is missed", c.name), cli.ErrRequiredArgNotFound)
		}

		note, err := c.storage.GetSecretById(model.Note, identity)
		if err != nil {
			return logger.WrapError("get secret", err)
		}

		toRemove = append(toRemove, note.(*secret.NoteSecret))
	}

	for _, note := range toRemove {
		logger.InfoFormat("Remove %s note", note.GetIdentity())
		err := c.storage.RemoveSecret(note)
		if err != nil {
			return logger.WrapError("get secret", err)
		}
	}

	return nil
}
