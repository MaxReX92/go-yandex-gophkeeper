package commands

import (
	"fmt"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/generator"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/storage"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model/secret"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

const (
	notesEditCommandName      = "edit"
	notesEditShortDescription = "edit notes from store"
	notesEditFullDescription  = "Edit new notes from secure store,"
)

type notesEditCommand struct {
	*baseCommand
	generator generator.Generator
	storage   storage.LocalSecretsStorage
}

func NewNotesEditCommand(
	stream io.CommandStream,
	generator generator.Generator,
	storage storage.LocalSecretsStorage,
	children ...cli.Command,
) *notesEditCommand {
	command := &notesEditCommand{
		generator: generator,
		storage:   storage,
	}
	command.baseCommand = newBaseCommand(
		stream,
		notesEditCommandName,
		notesEditShortDescription,
		notesEditFullDescription,
		children,
		[]cli.Argument{
			newArgument("Secret identity", true, initialFullDescription, initialShortDescription),
			newArgument("Note text", true, textFullArgName, textShortArgName),
			newArgument("Comment", true, commentFullArgName, commentShortArgName),
		},
		command.invoke,
	)
	return command
}

func (c *notesEditCommand) invoke(args map[string]string) error {
	identity, ok := argValue(args, idFullArgName, idShortArgName)
	if !ok {
		return logger.WrapError(fmt.Sprintf("invoke %s command: secret identity is missed", c.name), cli.ErrRequiredArgNotFound)
	}

	currentNotes, err := c.storage.GetSecretById(model.Notes, identity)
	if err != nil {
		return logger.WrapError("get secret", err)
	}

	if currentNotes == nil {
		return logger.WrapError("edit secret", cli.ErrSecretNotFound)
	}

	notes, ok := currentNotes.(*secret.NotesSecret)
	if !ok {
		return logger.WrapError("edit secret", cli.ErrInvalidSecretType)
	}

	text, ok := argValue(args, textFullArgName, textShortArgName)
	if ok {
		notes.Text = text
	}

	comment, ok := argValue(args, commentFullArgName, commentShortArgName)
	if ok {
		notes.Comment = comment
	}

	logger.Info("Edit note")
	err = c.storage.ChangeSecret(notes)
	if err != nil {
		return logger.WrapError("edit secret", err)
	}

	return nil
}
