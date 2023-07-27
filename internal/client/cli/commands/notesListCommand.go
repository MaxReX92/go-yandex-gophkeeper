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
	notesListCommandName      = "list"
	notesListShortDescription = "list of all notes"
	notesListFullDescription  = "Command list all stored notes."
)

type notesListCommand struct {
	*baseCommand
	storage storage.LocalSecretsStorage
}

func NewNotesListCommand(stream io.CommandStream, storage storage.LocalSecretsStorage, children ...cli.Command) *notesListCommand {
	command := &notesListCommand{
		storage: storage,
	}
	command.baseCommand = newBaseCommand(
		stream,
		notesListCommandName,
		notesListShortDescription,
		notesListFullDescription,
		children,
		[]cli.Argument{
			newArgument("Reveal secret values", false, revealFullArgName, revealShortArgName),
		},
		command.invoke,
	)
	return command
}

func (c *notesListCommand) invoke(args map[string]string) error {
	_, reveal := argValue(args, revealFullArgName, revealShortArgName)

	notess, err := c.storage.GetAllSecrets(model.Notes)
	if err != nil {
		return logger.WrapError("get secrets", err)
	}

	for _, modelNotes := range notess {
		notes := modelNotes.(*secret.NotesSecret)
		value := "***"
		if reveal {
			value = notes.Text
		}

		c.stream.Write(fmt.Sprintf("\t%s\t\t%s\t\t%s", notes.GetIdentity(), value, notes.GetComment()))
	}

	return nil
}
