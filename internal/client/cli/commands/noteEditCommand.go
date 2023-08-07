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
	noteEditCommandName      = "edit"
	noteEditShortDescription = "edit note from store"
	noteEditFullDescription  = "Edit new note from secure store,"
)

type noteEditCommand struct {
	*baseCommand
	storage storage.ClientSecretsStorage
}

func NewNoteEditCommand(
	stream io.CommandStream,
	storage storage.ClientSecretsStorage,
	children ...cli.Command,
) *noteEditCommand {
	command := &noteEditCommand{
		storage: storage,
	}
	command.baseCommand = newBaseCommand(
		stream,
		noteEditCommandName,
		noteEditShortDescription,
		noteEditFullDescription,
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

func (c *noteEditCommand) invoke(args map[string]string) error {
	identity, ok := argValue(args, idFullArgName, idShortArgName)
	if !ok {
		return logger.WrapError(fmt.Sprintf("invoke %s command: secret identity is missed", c.name), cli.ErrRequiredArgNotFound)
	}

	currentNote, err := c.storage.GetSecretById(nil, model.Note, identity)
	if err != nil {
		return logger.WrapError("get secret", err)
	}

	if currentNote == nil {
		return logger.WrapError("edit secret", cli.ErrSecretNotFound)
	}

	note, ok := currentNote.(*secret.NoteSecret)
	if !ok {
		return logger.WrapError("edit secret", cli.ErrInvalidSecretType)
	}

	text, ok := argValue(args, textFullArgName, textShortArgName)
	if ok {
		note.Text = text
	}

	comment, ok := argValue(args, commentFullArgName, commentShortArgName)
	if ok {
		note.Comment = comment
	}

	logger.Info("Edit note")
	err = c.storage.ChangeSecret(nil, note)
	if err != nil {
		return logger.WrapError("edit secret", err)
	}

	return nil
}
