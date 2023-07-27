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
	notesRemoveCommandName      = "remove"
	notesRemoveShortDescription = "remove notes from store"
	notesRemoveFullDescription  = "Remove notes from secure store,"
)

type notesRemoveCommand struct {
	*baseCommand
	storage storage.LocalSecretsStorage
}

func NewNotesRemoveCommand(
	stream io.CommandStream,
	storage storage.LocalSecretsStorage,
	children ...cli.Command,
) *notesRemoveCommand {
	command := &notesRemoveCommand{
		storage: storage,
	}
	command.baseCommand = newBaseCommand(
		stream,
		notesRemoveCommandName,
		notesRemoveShortDescription,
		notesRemoveFullDescription,
		children,
		[]cli.Argument{
			newArgument("Remove all notes secrets", false, allFullArgName),
			newArgument("Secret identity", true, idFullArgName, idShortArgName),
		},
		command.invoke,
	)
	return command
}

func (c *notesRemoveCommand) invoke(args map[string]string) error {
	var toRemove []*secret.NotesSecret
	_, removeAll := argValue(args, allFullArgName)
	if removeAll {
		notes, err := c.storage.GetAllSecrets(model.Notes)
		if err != nil {
			return logger.WrapError("get all secrets", err)
		}

		for _, note := range notes {
			toRemove = append(toRemove, note.(*secret.NotesSecret))
		}

	} else {
		identity, ok := argValue(args, idShortArgName, idFullArgName)
		if !ok {
			return logger.WrapError(fmt.Sprintf("invoke %s command: secret identity is missed", c.name), cli.ErrRequiredArgNotFound)
		}

		notes, err := c.storage.GetSecretById(model.Notes, identity)
		if err != nil {
			return logger.WrapError("get secret", err)
		}

		toRemove = append(toRemove, notes.(*secret.NotesSecret))
	}

	for _, notes := range toRemove {
		logger.InfoFormat("Remove %s notes", notes.GetIdentity())
		err := c.storage.RemoveSecret(notes)
		if err != nil {
			return logger.WrapError("get secret", err)
		}
	}

	return nil
}
